package versionctl

import (
	"embed"
	"errors"
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/db/model"
	"gitee.com/zhaochuninhefei/footprint-go/utils"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"gitee.com/zhaochuninhefei/zcutils-go/zctime"
	"gorm.io/gorm"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

//goland:noinspection SqlResolve,GoUnusedConst,GoSnakeCaseUsage,SqlInsertValues
const (
	DB_VERSION_CTL_COLS_WITHOUT_ID = "business_space, major_version, minor_version, patch_version, extend_version, version, custom_name, version_type, script_file_name, script_digest_hex, success, execution_time, install_time, install_user"

	sqlGetLastVersionByBSForMySQL8 = "SELECT " +
		DB_VERSION_CTL_COLS_WITHOUT_ID +
		" FROM (SELECT ROW_NUMBER() OVER(" +
		"PARTITION BY business_space " +
		"ORDER BY major_version desc,minor_version desc,patch_version desc,extend_version desc" +
		") row_no, " + DB_VERSION_CTL_COLS_WITHOUT_ID + " FROM brood_db_version_ctl" +
		") a where a.row_no=1"
	sqlGetVersionsOrderByVersions = "SELECT " +
		DB_VERSION_CTL_COLS_WITHOUT_ID +
		" FROM brood_db_version_ctl" +
		" ORDER BY business_space DESC, major_version DESC,minor_version DESC,patch_version DESC,extend_version DESC"
	sqlInsertVersionCtl = "INSERT INTO brood_db_version_ctl(business_space, major_version, minor_version, patch_version, extend_version, version, custom_name, version_type, script_file_name, script_digest_hex, success, execution_time, install_time, install_user) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	sqlUpdateVersionCtl = "UPDATE brood_db_version_ctl set success = 1, execution_time = ? WHERE business_space = ? AND major_version = ? AND minor_version = ? AND patch_version = ? AND extend_version = ?"
)

// DbVersionCtlTask 数据库版本控制任务接口
type DbVersionCtlTask interface {
	// RunTask 执行任务
	//  @return error
	RunTask() error
}

// DbVersionCtlContext 数据库版本控制上下文
type DbVersionCtlContext struct {
	dbClient *gorm.DB           // 数据库客户端
	props    *DbVersionCtlProps // 版本控制配置
	dbFS     *embed.FS          // sql脚本目录嵌入FS
	lastTask bool               // 是否最后一个任务
}

// CreateVersionTblTask 数据库版本控制表创建任务
type CreateVersionTblTask struct {
	// 嵌入上下文
	DbVersionCtlContext
}

// RunTask 执行数据库版本控制表创建任务
//  注意，如果使用了自定义的数据库版本管理表建表文路径(DbVersionTableCreateSqlPath)，
//  即使DbVersionTableName依然错误地定义为默认表名`brood_db_version_ctl`，
//  也不会替换为默认表名，而是以自定义的数据库版本管理表建表文路径的建表文为准。
//  @receiver cvtt 数据库版本控制表创建任务
//  @return error
func (cvtt *CreateVersionTblTask) RunTask() error {
	zclog.Info("CreateVersionTblTask begin...")
	dbVersionTableCreateSqlPath := cvtt.props.DbVersionTableCreateSqlPath
	zclog.Debugf("版本控制表建表文路径: %s", dbVersionTableCreateSqlPath)
	pathTmpArr := strings.Split(dbVersionTableCreateSqlPath, ":")
	if len(pathTmpArr) != 2 {
		return fmt.Errorf("数据库版本管理表建表文路径(DbVersionTableCreateSqlPath)不正确: %s", dbVersionTableCreateSqlPath)
	}
	var sqlBytes []byte
	var err error
	switch ScriptResourceMode(pathTmpArr[0]) {
	case EMBEDFS:
		sqlBytes, err = cvtt.dbFS.ReadFile(dbVersionTableCreateSqlPath)
		if err != nil {
			return err
		}
	case FILESYSTEM:
		sqlBytes, err = ioutil.ReadFile(dbVersionTableCreateSqlPath)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("不支持的SQL脚本文件类型: %s", pathTmpArr[0])
	}
	// 判断是否需要替换数据库版本控制表的表名
	sqlCreate := string(sqlBytes)
	if strings.EqualFold(dbVersionTableCreateSqlPath, defaultDbVersionTableCreateSqlPath) &&
		!strings.EqualFold(defaultDbVersionTableName, cvtt.props.DbVersionTableName) {
		sqlCreate = strings.ReplaceAll(sqlCreate, defaultDbVersionTableName, cvtt.props.DbVersionTableName)
	}
	// 执行建表SQL
	err = utils.RunSqlScript(cvtt.dbClient, sqlCreate)
	if err != nil {
		return err
	}

	zclog.Info("CreateVersionTblTask end...")
	return nil
}

// DropVersionTblTask 数据库版本控制表删除任务
type DropVersionTblTask struct {
	// 嵌入上下文
	DbVersionCtlContext
}

// RunTask 执行数据库版本控制表删除任务
//  注意，如果数据库版本控制表的建表文是自定义的，那么这里一定要确保配置的DbVersionTableName与自定义建表文中的表名一致。
//  @receiver dvtt 数据库版本控制表删除任务
//  @return error
func (dvtt *DropVersionTblTask) RunTask() error {
	zclog.Info("DropVersionTblTask begin...")
	dbVersionTableName := dvtt.props.DbVersionTableName
	err := dvtt.dbClient.Migrator().DropTable(dbVersionTableName)
	if err != nil {
		return err
	}
	zclog.Info("DropVersionTblTask end...")
	return nil
}

// IncreaseVersionTask 增量SQL执行任务
type IncreaseVersionTask struct {
	// 嵌入上下文
	DbVersionCtlContext
}

// RunTask 执行增量SQL任务
//  @receiver ivt 增量SQL执行任务
//  @return error
func (ivt *IncreaseVersionTask) RunTask() error {
	zclog.Info("IncreaseVersionTask begin...")

	// 从数据库版本控制表读取各个业务空间的最新版本
	versionCtls := make([]model.BroodDbVersionCtl, 0)
	filters := make(map[string]SqlScriptFilter)

	if strings.EqualFold("mysql8", ivt.props.DriverClassName) {
		// mysql8支持`ROW_NUMBER() OVER()`函数，使用SQL直接获取每种业务空间的最新版本
		ivt.dbClient.
			Raw(strings.ReplaceAll(sqlGetLastVersionByBSForMySQL8,
				defaultDbVersionTableName, ivt.props.DbVersionTableName)).
			Scan(&versionCtls)
		for _, versionCtl := range versionCtls {
			filters[versionCtl.BusinessSpace] = SqlScriptFilter{
				BusinessSpace: versionCtl.BusinessSpace,
				MajorVersion:  versionCtl.MajorVersion,
				MinorVersion:  versionCtl.MinorVersion,
				PatchVersion:  versionCtl.PatchVersion,
				ExtendVersion: versionCtl.ExtendVersion,
			}
		}
	} else {
		// 不支持`ROW_NUMBER() OVER()`函数，获取排序后的版本数据，每种业务空间取第一条数据
		ivt.dbClient.
			Raw(strings.ReplaceAll(sqlGetVersionsOrderByVersions,
				defaultDbVersionTableName, ivt.props.DbVersionTableName)).
			Scan(&versionCtls)
		for _, versionCtl := range versionCtls {
			bs := versionCtl.BusinessSpace
			_, ok := filters[bs]
			if !ok {
				filters[bs] = SqlScriptFilter{
					BusinessSpace: bs,
					MajorVersion:  versionCtl.MajorVersion,
					MinorVersion:  versionCtl.MinorVersion,
					PatchVersion:  versionCtl.PatchVersion,
					ExtendVersion: versionCtl.ExtendVersion,
				}
			}
		}
	}

	// 获取数据库版本升级SQL脚本目录集合
	sqlScriptDirPaths := ivt.props.ScriptDirs
	sqlDirPaths := strings.Split(sqlScriptDirPaths, ",")
	if len(sqlDirPaths) == 0 {
		return errors.New("数据库版本控制的属性sql脚本文件目录(script_dirs)未配置")
	}

	allFileInfos := make([]*SqlScriptInfo, 0)
	// 读取各个脚本目录下的SQL脚本，根据业务空间过滤出增量SQL脚本，获得增量脚本集合
	for _, sqlDirPath := range sqlDirPaths {
		sqlDirPath = strings.TrimSpace(sqlDirPath)
		zclog.Debugf("读取SQL脚本目录: %s", sqlDirPath)
		tmpArr := strings.Split(sqlDirPath, ":")
		if len(tmpArr) != 2 {
			return fmt.Errorf("数据库版本控制的属性sql脚本文件目录(script_dirs)配置的sql脚本目录格式不正确: %s", sqlDirPath)
		}
		var reader SqlScriptReader
		switch tmpArr[0] {
		case string(EMBEDFS):
			reader = ivt.dbFS
		case string(FILESYSTEM):
			reader = &FSSqlReader{}
		default:
			return fmt.Errorf("不支持的SQL脚本资源模式: %s", tmpArr[0])
		}

		subScriptInfos, err := ReadSql(reader, tmpArr[1], filters)
		if err != nil {
			return err
		}
		allFileInfos = append(allFileInfos, subScriptInfos...)
	}

	// 对增量脚本集合按业务空间做分组，并排序
	group := make(map[string][]*SqlScriptInfo)
	for _, fileInfo := range allFileInfos {
		group[fileInfo.BusinessSpace] = append(group[fileInfo.BusinessSpace], fileInfo)
	}
	for _, subInfos := range group {
		sort.SliceStable(subInfos, func(i, j int) bool {
			infoI := subInfos[i]
			infoJ := subInfos[j]
			if infoI.MajorVersion == infoJ.MajorVersion {
				if infoI.MinorVersion == infoJ.MinorVersion {
					if infoI.PatchVersion == infoJ.PatchVersion {
						return infoI.ExtendVersion < infoJ.ExtendVersion
					} else {
						return infoI.PatchVersion < infoJ.PatchVersion
					}
				} else {
					return infoI.MinorVersion < infoJ.MinorVersion
				}
			} else {
				return infoI.MajorVersion < infoJ.MajorVersion
			}
		})
	}

	// 生成数据库版本插入与更新SQL语句
	insertSql := strings.ReplaceAll(sqlInsertVersionCtl, defaultDbVersionTableName, ivt.props.DbVersionTableName)
	updateSql := strings.ReplaceAll(sqlUpdateVersionCtl, defaultDbVersionTableName, ivt.props.DbVersionTableName)

	// 遍历增量SQL脚本，读取SQL脚本，插入数据库版本控制数据，执行SQL脚本，更新版本记录
	for bs, subInfos := range group {
		if len(subInfos) == 0 {
			zclog.Infof("业务空间 %s 没有增量sql脚本需要执行.", bs)
		}
		for _, scriptInfo := range subInfos {
			zclog.Infof("增量执行脚本: %s", scriptInfo.Name)
			// 开启事务执行
			err := ivt.dbClient.Transaction(func(tx *gorm.DB) error {
				// 开始时间
				startTime := time.Now()
				// 插入版本记录
				result := tx.Exec(insertSql, scriptInfo.BusinessSpace, scriptInfo.MajorVersion, scriptInfo.MinorVersion,
					scriptInfo.PatchVersion, scriptInfo.ExtendVersion, scriptInfo.Version, scriptInfo.CustomName,
					"SQL", scriptInfo.Name, "none", 0, -1, startTime.Format(zctime.TIME_FORMAT_DASH), ivt.props.Username)
				if result.Error != nil {
					return result.Error
				}
				// 执行脚本
				err := utils.RunSqlScript(tx, scriptInfo.Content)
				if err != nil {
					return err
				}
				// 结束时间
				stopTime := time.Now()
				mills := stopTime.Sub(startTime).Milliseconds()
				zclog.Infof("sql脚本 %s 执行耗时 : %d ms.", scriptInfo.Name, mills)
				// 更新版本记录
				result = tx.Exec(updateSql, mills, scriptInfo.BusinessSpace, scriptInfo.MajorVersion,
					scriptInfo.MinorVersion, scriptInfo.PatchVersion, scriptInfo.ExtendVersion)
				if result.Error != nil {
					return result.Error
				}
				return nil
			})
			if err != nil {
				return err
			}
			zclog.Infof("数据库版本记录更新, business_space: %s , major_version: %d , minor_version: %d , patch_version: %d , extend_version: %d .",
				scriptInfo.BusinessSpace, scriptInfo.MajorVersion, scriptInfo.MinorVersion, scriptInfo.PatchVersion, scriptInfo.ExtendVersion)
		}
	}
	zclog.Info("IncreaseVersionTask end...")
	return nil
}

// InsertBaselineTask 数据库基线版本记录插入任务
type InsertBaselineTask struct {
	// 嵌入上下文
	DbVersionCtlContext
}

// PTN_VERSION_DEFAULT 基线版本正则表达式(默认)
//goland:noinspection GoSnakeCaseUsage
var PTN_VERSION_DEFAULT *regexp.Regexp

// PTN_VERSION_EXTEND 基线版本正则表达式(带扩展版本号)
//goland:noinspection GoSnakeCaseUsage
var PTN_VERSION_EXTEND *regexp.Regexp

// init 初始化基线版本正则表达式
func init() {
	PTN_VERSION_DEFAULT = regexp.MustCompile("^([A-Za-z0-9]+)_V(\\d+)\\.(\\d+)\\.(\\d+)$")
	if PTN_VERSION_DEFAULT == nil {
		panic("正则表达式不正确: ^([A-Za-z0-9]+)_V(\\d+)\\.(\\d+)\\.(\\d+)$")
	}
	PTN_VERSION_EXTEND = regexp.MustCompile("^([A-Za-z0-9]+)_V(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+)$")
	if PTN_VERSION_EXTEND == nil {
		panic("正则表达式不正确: ^([A-Za-z0-9]+)_V(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+)$")
	}
}

func (ibt *InsertBaselineTask) RunTask() error {
	zclog.Info("InsertBaselineTask begin...")

	// 生成数据库版本插入SQL语句
	insertSql := strings.ReplaceAll(sqlInsertVersionCtl, defaultDbVersionTableName, ibt.props.DbVersionTableName)

	// 获取基线版本信息
	baselineArr := strings.Split(strings.TrimSpace(ibt.props.BaselineBusinessSpaceAndVersions), ",")
	insertTimes := 0
	for _, version := range baselineArr {
		version = strings.TrimSpace(version)
		if version == "" {
			continue
		}
		var bs string
		var major, minor, patch, extend int64
		var err error
		matched := false
		matcherDefault := PTN_VERSION_DEFAULT.FindAllStringSubmatch(version, -1)
		if len(matcherDefault) > 0 {
			for _, strMatched := range matcherDefault {
				if len(strMatched) == 5 {
					bs = strMatched[1]
					major, err = strconv.ParseInt(strMatched[2], 10, 64)
					if err != nil {
						return err
					}
					minor, err = strconv.ParseInt(strMatched[3], 10, 64)
					if err != nil {
						return err
					}
					patch, err = strconv.ParseInt(strMatched[4], 10, 64)
					if err != nil {
						return err
					}
					extend = 0
					matched = true
					break
				}
			}
		} else {
			// 使用带扩展版本号的正则表达式解析SQL文件名
			matcherExtend := PTN_VERSION_EXTEND.FindAllStringSubmatch(version, -1)
			for _, strMatched := range matcherExtend {
				if len(strMatched) == 6 {
					bs = strMatched[1]
					major, err = strconv.ParseInt(strMatched[2], 10, 64)
					if err != nil {
						return err
					}
					minor, err = strconv.ParseInt(strMatched[3], 10, 64)
					if err != nil {
						return err
					}
					patch, err = strconv.ParseInt(strMatched[4], 10, 64)
					if err != nil {
						return err
					}
					extend, err = strconv.ParseInt(strMatched[5], 10, 64)
					if err != nil {
						return err
					}
					matched = true
					break
				}
			}
		}
		if !matched {
			return fmt.Errorf("数据库基线版本(BaselineBusinessSpaceAndVersions)配置格式错误: %s",
				ibt.props.DbVersionTableName)
		}
		// 插入数据库版本控制的基线版本记录
		result := ibt.dbClient.Exec(insertSql, bs, major, minor,
			patch, extend, version, "none", "BaseLine", "none", "none", 1, 0,
			time.Now().Format(zctime.TIME_FORMAT_DASH), ibt.props.Username)
		if result.Error != nil {
			return result.Error
		}
		insertTimes++
		zclog.Infof("数据库基线版本添加, business_space: %s , major_version: %d , minor_version: %d , patch_version: %d, extend_version: %d .",
			bs, major, minor, patch, extend)
	}
	if insertTimes == 0 {
		return fmt.Errorf("未能成功执行任何基线版本记录插入，请检查数据库基线版本(BaselineBusinessSpaceAndVersions)配置: %s",
			ibt.props.DbVersionTableName)
	}

	zclog.Info("InsertBaselineTask end...")
	return nil
}
