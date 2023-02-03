package mysql

import (
	"gitee.com/zhaochuninhefei/footprint-go/test/resources"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"testing"
)

func TestConnectMysqlByDefault(t *testing.T) {
	zclog.Info("===== TestConnectMysqlByDefault 开始 =====")
	RunDBTest()
	zclog.Info("===== TestConnectMysqlByDefault 结束 =====")
}

func TestExecScriptFile(t *testing.T) {
	scriptSQL, err := resources.DBFilesTest.ReadFile("db/beforeclass/clear_footprinttest.sql")
	if err != nil {
		t.Fatal(err)
	}

	mysqlClient, err := ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		t.Fatal(err)
	}

	db := mysqlClient.Exec(string(scriptSQL))
	if err = db.Error; err != nil {
		t.Fatal(err)
	}
}
