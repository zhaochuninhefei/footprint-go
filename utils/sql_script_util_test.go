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

	scriptSQL, err := resources.DBFilesTest.ReadFile("db/beforeclass/clear_footprinttest.sql")
	if err != nil {
		t.Fatal(err)
	}

	sqls := ReadSqls(string(scriptSQL))
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
