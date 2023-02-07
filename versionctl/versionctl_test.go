package versionctl

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/db/mysql"
	"gitee.com/zhaochuninhefei/footprint-go/test/resources"
	"gitee.com/zhaochuninhefei/footprint-go/utils"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func TestDoDBVersionControl_deploy_init(t *testing.T) {
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

	myDb, err := mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		t.Fatal(err)
	}

	myProps := &DbVersionCtlProps{
		ScriptResourceMode:               EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01",
		BaselineBusinessSpaceAndVersions: "template_V2.11.0,smtp_V2.0.0",
		DbVersionTableName:               defaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      defaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql8",
		Host:                             "localhost",
		Port:                             "3307",
		Database:                         "db_footprint_test",
		Username:                         "zhaochun1",
		Password:                         "zhaochun@GITHUB",
		ExistTblQuerySql:                 defaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = DoDBVersionControl(myDb, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoDBVersionControl_deploy_increase(t *testing.T) {
	tbls, err := showTables()
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) < 2 {
		t.Fatal("测试数据不正确，即存表未导入，请确认是否先执行了TestDoDBVersionControl_deploy_init")
	}
	hasCtlTbl := false
	for _, tbl := range tbls {
		if tbl == defaultDbVersionTableName {
			hasCtlTbl = true
			break
		}
	}
	if !hasCtlTbl {
		t.Fatal("测试数据不正确，没有导入版本控制表，请确认是否先执行了TestDoDBVersionControl_deploy_init")
	}

	myProps := &DbVersionCtlProps{
		ScriptResourceMode:               EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02",
		BaselineBusinessSpaceAndVersions: "template_V2.11.0,smtp_V2.0.0",
		DbVersionTableName:               defaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      defaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql8",
		Host:                             "localhost",
		Port:                             "3307",
		Database:                         "db_footprint_test",
		Username:                         "zhaochun1",
		Password:                         "zhaochun@GITHUB",
		ExistTblQuerySql:                 defaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = DoDBVersionControl(nil, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoDBVersionControl_baseline_init(t *testing.T) {
	// 删除数据库版本表
	err := dropCtlTbl(defaultDbVersionTableName)
	if err != nil {
		t.Fatal(err)
	}
	tbls, err := showTables()
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) < 1 {
		t.Fatal("测试数据不正确，需要导入即存表，请确实是否先执行了TestDoDBVersionControl_deploy_init与TestDoDBVersionControl_deploy_increase")
	}
	hasCtlTbl := false
	for _, tbl := range tbls {
		if tbl == defaultDbVersionTableName {
			hasCtlTbl = true
			break
		}
	}
	if hasCtlTbl {
		t.Fatal("测试数据不正确，版本控制表未删除")
	}

	myProps := &DbVersionCtlProps{
		ScriptResourceMode:               EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02,embedfs:db/test03",
		BaselineBusinessSpaceAndVersions: "template_V2.11.0,smtp_V2.0.0",
		DbVersionTableName:               defaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      defaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql8",
		Host:                             "localhost",
		Port:                             "3307",
		Database:                         "db_footprint_test",
		Username:                         "zhaochun1",
		Password:                         "zhaochun@GITHUB",
		ExistTblQuerySql:                 defaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}

	err = DoDBVersionControl(nil, myProps, &resources.DBFilesTest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadSql(t *testing.T) {
	fmt.Println()
	fmt.Println("db/test01 下的文件")
	filter := make(map[string]SqlScriptFilter)
	filter["template"] = SqlScriptFilter{
		BusinessSpace: "template",
		MajorVersion:  1,
		MinorVersion:  0,
		PatchVersion:  12,
		ExtendVersion: 0,
	}
	files, err := ReadSql(&resources.DBFilesTest, "db/test01", filter)
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test02 下的文件")
	files, err = ReadSql(&FSSqlReader{}, "../test/resources/db/test02", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test03 下的文件")
	files, err = ReadSql(&resources.DBFilesTest, "db/test03", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test04 下的文件")
	files, err = ReadSql(&resources.DBFilesTest, "db/test04", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test05 下的文件")
	files, err = ReadSql(&resources.DBFilesTest, "db/test05", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test06 下的文件")
	files, err = ReadSql(&resources.DBFilesTest, "db/test06", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}
}

func TestAnalyzeDetailsFromSqlFileName(t *testing.T) {
	fmt.Println("---- bsTest_V1.2.3_init_test.sql ----")
	fileInfo1, err := createFileInfo("bsTest_V1.2.3_init_test.sql", "")
	if err != nil {
		t.Fatal(err)
	}
	jsonFileInfo1, _ := json.Marshal(fileInfo1)
	fmt.Println(string(jsonFileInfo1))

	fmt.Println("---- bsTest_V1.2.3.4_init_test.sql ----")
	fileInfo2, err := createFileInfo("bsTest_V1.2.3.4_init_test.sql", "")
	if err != nil {
		t.Fatal(err)
	}
	jsonFileInfo2, _ := json.Marshal(fileInfo2)
	fmt.Println(string(jsonFileInfo2))

	fmt.Println("---- bsTest_V1.2.3.4.5_init_test.sql ----")
	fileInfo3, err := createFileInfo("bsTest_V1.2.3.4.5_init_test.sql", "")
	if err != nil {
		if strings.HasPrefix(err.Error(), "sqlFileName未能正确匹配正则表达式:") {
			fmt.Println("返回了正确的错误信息")
			jsonFileInfo3, _ := json.Marshal(fileInfo3)
			fmt.Println(string(jsonFileInfo3))
			return
		} else {
			t.Fatal(err)
		}
	} else {
		t.Fatal("匹配了不正确的文件名")
	}
}

//func TestGroupAndSort(t *testing.T) {
//	allFileInfos, err := ReadSql(&resources.DBFilesTest, "db", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	group := make(map[string][]*SqlScriptInfo)
//	for _, fileInfo := range allFileInfos {
//		group[fileInfo.BusinessSpace] = append(group[fileInfo.BusinessSpace], fileInfo)
//	}
//
//	for bs, subInfos := range group {
//		fmt.Printf("业务空间: %s\n", bs)
//		fmt.Println("排序前:")
//		for _, info := range subInfos {
//			fmt.Printf("  sql脚本名: %s\n", info.Name)
//		}
//		sort.SliceStable(subInfos, func(i, j int) bool {
//			infoI := subInfos[i]
//			infoJ := subInfos[j]
//			if infoI.MajorVersion == infoJ.MajorVersion {
//				if infoI.MinorVersion == infoJ.MinorVersion {
//					if infoI.PatchVersion == infoJ.PatchVersion {
//						return infoI.ExtendVersion < infoJ.ExtendVersion
//					} else {
//						return infoI.PatchVersion < infoJ.PatchVersion
//					}
//				} else {
//					return infoI.MinorVersion < infoJ.MinorVersion
//				}
//			} else {
//				return infoI.MajorVersion < infoJ.MajorVersion
//			}
//		})
//		fmt.Println("排序后:")
//		for _, info := range subInfos {
//			fmt.Printf("  sql脚本名: %s\n", info.Name)
//		}
//	}
//}

func TestQueryExistTblNames(t *testing.T) {
	ctlProps = FillDefaultFields(prepareCtlProps())
	var err error
	dbClient, err = mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		zclog.Errorln(err)
		return
	}

	tables := queryExistTblNames()
	fmt.Println(tables)
}

func TestCheckBaselineResetConditionSql(t *testing.T) {
	ctlProps = FillDefaultFields(prepareCtlProps())
	//goland:noinspection SqlResolve
	ctlProps.BaselineResetConditionSql = "select 1 from brood_db_version_ctl"
	var err error
	dbClient, err = mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		t.Fatal(err)
	}

	result := checkBaselineResetConditionSql()
	fmt.Println(result)
}

func prepareCtlProps() *DbVersionCtlProps {
	props := &DbVersionCtlProps{
		ScriptResourceMode:               EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01",
		BaselineBusinessSpaceAndVersions: "template_V2.11.0,smtp_V2.0.0",
		DbVersionTableName:               "brood_db_version_ctl1",
		DbVersionTableCreateSqlPath:      "embedfs:db/versionctl/create_brood_db_version_ctl.sql",
		DriverClassName:                  "mysql",
		Host:                             "127.0.0.1",
		Port:                             "3307",
		Database:                         "db_footprint_test",
		Username:                         "zhaochun1",
		Password:                         "zhaochun@GITHUB",
		ExistTblQuerySql:                 "show tables",
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}
	return props
}

func clearDB() error {
	clearSqlBytes, err := resources.DBFilesTest.ReadFile("db/beforeclass/clear_footprinttest.sql")
	if err != nil {
		return err
	}
	myDb, err := mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
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
	myDb, err := mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		return nil, err
	}
	tables := make([]string, 0)
	result := myDb.Raw(defaultExistTblQuerySql).Scan(&tables)
	if err = result.Error; err != nil {
		return nil, err
	}
	fmt.Printf("show tables: %s\n", tables)
	return tables, nil
}

func dropCtlTbl(tblName string) error {
	myDb, err := mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		return err
	}
	result := myDb.Exec("DROP TABLE IF EXISTS " + tblName)
	if err = result.Error; err != nil {
		return err
	}
	return nil
}
