package versionctl

import (
	"embed"
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/utils"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

type DbVersionCtlTask interface {
	RunTask() error
}

type DbVersionCtlContext struct {
	dbClient *gorm.DB           // 数据库客户端
	props    *DbVersionCtlProps // 版本控制配置
	dbFS     *embed.FS          // sql脚本目录嵌入FS
	lastTask bool               // 是否最后一个任务
}

type CreateVersionTblTask struct {
	DbVersionCtlContext
}

func (cvtt CreateVersionTblTask) RunTask() error {
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

	err = utils.RunSqlScript(cvtt.dbClient, string(sqlBytes))
	if err != nil {
		return err
	}

	zclog.Info("CreateVersionTblTask end...")
	return nil
}
