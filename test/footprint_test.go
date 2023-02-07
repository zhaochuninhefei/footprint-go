package test

import (
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/db/mysql"
	"gitee.com/zhaochuninhefei/footprint-go/test/resources"
	"gitee.com/zhaochuninhefei/footprint-go/utils"
	"gitee.com/zhaochuninhefei/footprint-go/versionctl"
	"gorm.io/gorm"
	"testing"
)

const (
	dbHost = "localhost"
	dbPort = "3307"
	dbUser = "zhaochun1"
	dbPwd  = "zhaochun@GITHUB"
	dbName = "db_footprint_test"
)

func Test01_deploy_init(t *testing.T) {
	err := clearDB()
	if err != nil {
		t.Fatal(err)
	}
	tbls, err := showTables()
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) > 0 {
		t.Fatal("未能成功清理测试数据库")
	}

	myDb, err := mysql.ConnectMysqlByDefault(nil, dbHost, dbPort, dbUser, dbPwd, dbName)
	if err != nil {
		t.Fatal(err)
	}

	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01",
		BaselineBusinessSpaceAndVersions: "",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = versionctl.DoDBVersionControl(myDb, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

func Test02_deploy_increase(t *testing.T) {
	tbls, err := showTables()
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) < 2 {
		t.Fatal("测试数据不正确，即存表未导入，请确认是否先执行了Test01_deploy_init")
	}
	hasCtlTbl := false
	for _, tbl := range tbls {
		if tbl == versionctl.DefaultDbVersionTableName {
			hasCtlTbl = true
			break
		}
	}
	if !hasCtlTbl {
		t.Fatal("测试数据不正确，没有导入版本控制表，请确认是否先执行了Test01_deploy_init")
	}

	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02",
		BaselineBusinessSpaceAndVersions: "",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = versionctl.DoDBVersionControl(nil, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

func Test03_baseline_init(t *testing.T) {
	// 删除数据库版本表
	err := dropCtlTbl(versionctl.DefaultDbVersionTableName)
	if err != nil {
		t.Fatal(err)
	}
	tbls, err := showTables()
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) < 1 {
		t.Fatal("测试数据不正确，需要导入即存表，请确实是否先执行了Test01_deploy_init与Test02_deploy_increase")
	}
	hasCtlTbl := false
	for _, tbl := range tbls {
		if tbl == versionctl.DefaultDbVersionTableName {
			hasCtlTbl = true
			break
		}
	}
	if hasCtlTbl {
		t.Fatal("测试数据不正确，版本控制表未删除")
	}

	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02,embedfs:db/test03",
		BaselineBusinessSpaceAndVersions: "template_V2.11.0,smtp_V2.0.0",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = versionctl.DoDBVersionControl(nil, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

func Test04_deploy_increase(t *testing.T) {
	tbls, err := showTables()
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) < 2 {
		t.Fatal("测试数据不正确，即存表未导入，请确认是否先执行了Test01_xxx到Test03_xxx")
	}
	hasCtlTbl := false
	for _, tbl := range tbls {
		if tbl == versionctl.DefaultDbVersionTableName {
			hasCtlTbl = true
			break
		}
	}
	if !hasCtlTbl {
		t.Fatal("测试数据不正确，没有导入版本控制表，请确认是否先执行了Test01_xxx到Test03_xxx")
	}

	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02,embedfs:db/test03,embedfs:db/test04",
		BaselineBusinessSpaceAndVersions: "template_V2.11.0,smtp_V2.0.0",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = versionctl.DoDBVersionControl(nil, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
func Test05_baseline_reset(t *testing.T) {
	tbls, err := showTables()
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) < 2 {
		t.Fatal("测试数据不正确，即存表未导入，请确认是否先执行了Test01_xxx到Test04_xxx")
	}
	hasCtlTbl := false
	for _, tbl := range tbls {
		if tbl == versionctl.DefaultDbVersionTableName {
			hasCtlTbl = true
			break
		}
	}
	if !hasCtlTbl {
		t.Fatal("测试数据不正确，没有导入版本控制表，请确认是否先执行了Test01_xxx到Test04_xxx")
	}

	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02,embedfs:db/test03,embedfs:db/test04,embedfs:db/test05",
		BaselineBusinessSpaceAndVersions: "template_V3.11.999,smtp_V3.0.999",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "y",
		BaselineResetConditionSql:        "SELECT 1 FROM brood_db_version_ctl WHERE version = 'template_V3.10.11'",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = versionctl.DoDBVersionControl(nil, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
func Test06_deploy_increase(t *testing.T) {
	tbls, err := showTables()
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) < 2 {
		t.Fatal("测试数据不正确，即存表未导入，请确认是否先执行了Test01_xxx到Test05_xxx")
	}
	hasCtlTbl := false
	for _, tbl := range tbls {
		if tbl == versionctl.DefaultDbVersionTableName {
			hasCtlTbl = true
			break
		}
	}
	if !hasCtlTbl {
		t.Fatal("测试数据不正确，没有导入版本控制表，请确认是否先执行了Test01_xxx到Test05_xxx")
	}

	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.FILESYSTEM,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02,embedfs:db/test03,embedfs:db/test04,embedfs:db/test05,filesystem:resources/db/test06",
		BaselineBusinessSpaceAndVersions: "template_V3.11.999,smtp_V3.0.999",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "y",
		BaselineResetConditionSql:        "SELECT 1 FROM brood_db_version_ctl WHERE version = 'template_V3.10.11'",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = versionctl.DoDBVersionControl(nil, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

func clearDB() error {
	clearSqlBytes, err := resources.DBFilesTest.ReadFile("db/beforeclass/clear_footprinttest.sql")
	if err != nil {
		return err
	}
	myDb, err := mysql.ConnectMysqlByDefault(nil, dbHost, dbPort, dbUser, dbPwd, dbName)
	if err != nil {
		return err
	}
	err = myDb.Transaction(func(tx *gorm.DB) error {
		err = utils.RunSqlScript(tx, string(clearSqlBytes))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func showTables() ([]string, error) {
	myDb, err := mysql.ConnectMysqlByDefault(nil, dbHost, dbPort, dbUser, dbPwd, dbName)
	if err != nil {
		return nil, err
	}
	tables := make([]string, 0)
	result := myDb.Raw(versionctl.DefaultExistTblQuerySql).Scan(&tables)
	if err = result.Error; err != nil {
		return nil, err
	}
	fmt.Printf("show tables: %s\n", tables)
	return tables, nil
}

//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
func dropCtlTbl(tblName string) error {
	myDb, err := mysql.ConnectMysqlByDefault(nil, dbHost, dbPort, dbUser, dbPwd, dbName)
	if err != nil {
		return err
	}
	result := myDb.Exec("DROP TABLE IF EXISTS " + tblName)
	if err = result.Error; err != nil {
		return err
	}
	return nil
}
