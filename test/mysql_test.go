package test

import (
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"github.com/zhaochuninhefei/footprint-go/db/mysql"
	"io/ioutil"
	"testing"
)

func TestConnectMysqlByDefault(t *testing.T) {
	zclog.Info("===== TestConnectMysqlByDefault 开始 =====")
	mysqlClient, err := mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		t.Fatal(err)
	}

	mysqlClient.Exec("drop table if exists `brood_db_version_ctl`")

	createSql, err := ioutil.ReadFile("../resources/db/versionctl/create_brood_db_version_ctl.sql")
	if err != nil {
		t.Fatal(err)
	}
	mysqlClient.Exec(string(createSql))
	zclog.Info("===== TestConnectMysqlByDefault 结束 =====")
}
