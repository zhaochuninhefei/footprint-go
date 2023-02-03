package utils

import (
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/db/mysql"
	"gitee.com/zhaochuninhefei/footprint-go/test/resources"
	"testing"
)

func TestReadSqls(t *testing.T) {
	mysqlClient, err := mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		t.Fatal(err)
	}

	scriptSQL, err := resources.DBFilesTest.ReadFile("db/test01/smtp_V1.0.0_init.sql")
	if err != nil {
		t.Fatal(err)
	}

	sqls, err := ReadSqls(string(scriptSQL))
	if err != nil {
		t.Fatal(err)
	}
	for _, sql := range sqls {
		fmt.Printf("读取到sql: %s\nend\n", sql)
	}

	for _, sql := range sqls {
		db := mysqlClient.Exec(sql)
		if err = db.Error; err != nil {
			t.Fatal(err)
		}
	}
}

//goland:noinspection SqlResolve
func TestRunSqlScript(t *testing.T) {
	mysqlClient, err := mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		t.Fatal(err)
	}

	scriptSQL, err := resources.DBFilesTest.ReadFile("db/test01/smtp_V1.0.0_init.sql")
	if err != nil {
		t.Fatal(err)
	}

	err = RunSqlScript(mysqlClient, string(scriptSQL))
	if err != nil {
		t.Fatal(err)
	}

	data := make([]string, 0)
	mysqlClient.Raw("SELECT smtp_host from rv_smtps").Scan(&data)
	fmt.Printf("查询到数据: %s\n", data)

	scriptSQL2, err := resources.DBFilesTest.ReadFile("db/test02/smtp_V1.0.999_add_smtp01.sql")
	if err != nil {
		t.Fatal(err)
	}

	err = RunSqlScript(mysqlClient, string(scriptSQL2))
	if err != nil {
		t.Fatal(err)
	}

	mysqlClient.Raw("SELECT smtp_host from rv_smtps").Scan(&data)
	fmt.Printf("查询到数据: %s\n", data)
}
