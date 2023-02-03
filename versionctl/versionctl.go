package versionctl

import (
	"gitee.com/zhaochuninhefei/footprint-go/db/mysql"
	"gorm.io/gorm"
)

// ctlProps 数据库版本控制相关配置
var ctlProps *DbVersionCtlProps

// db 目标数据库客户端
var dbClient *gorm.DB

func DoDBVersionControl(existDB *gorm.DB, props *DbVersionCtlProps) error {
	ctlProps = props
	var err error
	dbClient, err = mysql.ConnectMysqlByDefault(existDB, props.Host, props.Port, props.Username, props.Password, props.Database)
	if err != nil {
		return err
	}

	// 判断本次数据库版本控制的操作模式

	return nil
}

// OperationMode 数据库版本控制操作模式
type OperationMode int32

//goland:noinspection GoSnakeCaseUsage
const (
	// DEPLOY_INIT 项目首次部署，数据库没有任何表。
	//  该操作会生成数据库版本控制表，执行数据库初始化脚本，更新数据库版本控制表数据。
	DEPLOY_INIT OperationMode = 1

	// DEPLOY_INCREASE 项目增量部署，之前已经导入业务表与数据库版本控制表。
	//  该操作根据已有的数据库版本控制表中的记录判断哪些脚本需要执行，然后执行脚本并插入新的数据库版本记录。
	DEPLOY_INCREASE OperationMode = 2

	// BASELINE_INIT 一个已经上线的项目初次使用数据库版本控制，之前已经导入业务表，但没有数据库版本控制表。
	//  该操作会创建数据库版本控制表，并写入一条版本基线记录，然后基于属性配置的基线版本确定哪些脚本需要执行。
	//  执行脚本后向数据库版本控制表插入新的版本记录。
	BASELINE_INIT OperationMode = 3

	// BASELINE_RESET 对一个已经使用数据库版本控制的项目，重置其数据库版本的基线。
	//  该操作会删除既有的数据库版本控制表，然后重新做一次`BASELINE_INIT`操作。
	//  注意该操作需要特殊的属性控制，要慎用。
	BASELINE_RESET OperationMode = 4
)

func chargeOperationMode() OperationMode {

	return DEPLOY_INCREASE
}

func queryExistTblNames() []string {
	showTblSql := ctlProps.ExistTblQuerySql
	if showTblSql == "" {
		showTblSql = "show tables"
	}
	tables := make([]string, 0)
	dbClient.Raw(showTblSql).Scan(&tables)
	return tables
}
