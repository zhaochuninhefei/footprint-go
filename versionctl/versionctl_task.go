package versionctl

import (
	"embed"
	"errors"
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/db/model"
	"gitee.com/zhaochuninhefei/footprint-go/utils"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"gorm.io/gorm"
	"io/ioutil"
	"sort"
	"strings"
)

//goland:noinspection SqlResolve,GoUnusedConst
const (
	sqlGetLastVersionByBSForMySQL8 = "SELECT id, business_space, major_version ,minor_version ,patch_version ,extend_version " +
		"FROM (" +
		"SELECT ROW_NUMBER() OVER(" +
		"PARTITION BY business_space " +
		"ORDER BY major_version desc,minor_version desc,patch_version desc,extend_version desc" +
		") row_no, id, business_space, major_version ,minor_version ,patch_version ,extend_version FROM brood_db_version_ctl" +
		") a where a.row_no=1"
	sqlGetVersionsOrderByVersions = "SELECT id, business_space, major_version ,minor_version ,patch_version ,extend_version " +
		"FROM brood_db_version_ctl " +
		"ORDER BY business_space DESC, major_version DESC,minor_version DESC,patch_version DESC,extend_version DESC"
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
	dbVersionTableName := dvtt.props.DbVersionTableName
	err := dvtt.dbClient.Migrator().DropTable(dbVersionTableName)
	if err != nil {
		return err
	}
	return nil
}

type IncreaseVersionTask struct {
	DbVersionCtlContext
}

func (ivt *IncreaseVersionTask) RunTask() error {

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

	// 生成数据库版本插入SQL语句

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

	// 遍历增量SQL脚本，读取SQL脚本，插入数据库版本控制数据，执行SQL脚本，更新版本记录
	for _, subInfos := range group {

		for _, scriptInfo := range subInfos {
			// 开启事务执行
			err := ivt.dbClient.Transaction(func(tx *gorm.DB) error {
				result := tx.Exec("")
				if result.Error != nil {
					return result.Error
				}
				sql := scriptInfo.Content
				result = tx.Exec(sql)
				if result.Error != nil {
					return result.Error
				}
				return nil
			})
			if err != nil {
				return err
			}

		}
	}

	return nil
}
