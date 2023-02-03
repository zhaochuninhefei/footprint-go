package versionctl

import (
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/db/mysql"
	"gorm.io/gorm"
	"strings"
)

// ctlProps 数据库版本控制相关配置
var ctlProps *DbVersionCtlProps

// db 目标数据库客户端
var dbClient *gorm.DB

func DoDBVersionControl(existDB *gorm.DB, props *DbVersionCtlProps) error {
	ctlProps = FillDefaultFields(props)
	var err error
	dbClient, err = mysql.ConnectMysqlByDefault(existDB, props.Host, props.Port, props.Username, props.Password, props.Database)
	if err != nil {
		return err
	}

	// 判断本次数据库版本控制的操作模式
	operationMode := chargeOperationMode()

	switch operationMode {
	case DEPLOY_INIT:
	case BASELINE_INIT:
	case BASELINE_RESET:
	case DEPLOY_INCREASE:
	default:
		return fmt.Errorf("不支持的数据库版本控制操作模式: %d", operationMode)
	}

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
	tables := queryExistTblNames()
	// 判断当前database是否非空
	if len(tables) == 0 {
		// 当前database为空，首次启动服务，导入全部数据库脚本，并创建数据库版本控制表，并生成数据库版本记录。
		return DEPLOY_INIT
	}
	// 如果当前database非空，判断是否已经创建了数据库版本控制表"brood_db_version_ctl"
	var ctlTblExists = false
	dbVersionTableName := ctlProps.DbVersionTableName
	if dbVersionTableName == "" {
		dbVersionTableName = "brood_db_version_ctl"
	}
	for _, table := range tables {
		if table == dbVersionTableName {
			ctlTblExists = true
			break
		}
	}
	if ctlTblExists {
		// 判断是否需要重置数据库版本控制表
		if strings.EqualFold("y", ctlProps.BaselineReset) &&
			checkBaselineResetConditionSql() {
			// 查询数据库版本控制表的最新记录。
			// 只有属性[baselineResetConditionSql]配置的sql查询到有记录，才会执行基线重置操作。
			// baselineResetConditionSql在配置时建议将install_time字段作为条件去查询，这样以后不会再有满足该条件的记录。
			return BASELINE_RESET
		}
		// 已经存在数据库版本控制表，根据当前资源目录下的sql脚本与版本控制表中各个业务空间的最新版本做增量的sql脚本执行。
		return DEPLOY_INCREASE
	}
	// database非空，但还没有数据库版本控制表，根据配置参数[baselineBusinessSpaceAndVersions]决定各个业务空间的基线版本，
	// 创建数据库版本控制表，生成baseline记录；然后做增量的sql脚本执行。
	return BASELINE_INIT
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

func checkBaselineResetConditionSql() bool {
	baselineResetConditionSql := ctlProps.BaselineResetConditionSql
	if baselineResetConditionSql == "" {
		return false
	}
	results := make([]interface{}, 0)
	dbClient.Raw(baselineResetConditionSql).Scan(&results)
	return len(results) > 0
}
