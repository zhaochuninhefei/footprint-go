package test

import (
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"github.com/zhaochuninhefei/footprint-go/db/mysql"
	"github.com/zhaochuninhefei/footprint-go/resources/db/versionctl"
	"testing"
)

func TestConnectMysqlByDefault(t *testing.T) {
	zclog.Info("===== TestConnectMysqlByDefault 开始 =====")
	runDBTest()
	zclog.Info("===== TestConnectMysqlByDefault 结束 =====")
}

func runDBTest() {
	mysqlClient, err := mysql.ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		zclog.Errorln(err)
		return
	}

	mysqlClient.Exec("drop table if exists `brood_db_version_ctl`")

	zclog.Info(versionctl.CreateSql)

	mysqlClient.Exec(versionctl.CreateSql)
}
