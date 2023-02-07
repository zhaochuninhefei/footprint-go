package mysql

import (
	"gitee.com/zhaochuninhefei/footprint-go/resources"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"testing"
)

func TestConnectMysqlByDefault(t *testing.T) {
	zclog.Info("===== TestConnectMysqlByDefault 开始 =====")
	RunDBTest()
	zclog.Info("===== TestConnectMysqlByDefault 结束 =====")
}

//goland:noinspection SqlNoDataSourceInspection
func RunDBTest() {
	mysqlClient, err := ConnectMysqlByDefault(nil, "localhost", "3307", "zhaochun1", "zhaochun@GITHUB", "db_footprint_test")
	if err != nil {
		zclog.Errorln(err)
		return
	}

	db := mysqlClient.Exec("drop table if exists `brood_db_version_ctl`")
	if err := db.Error; err != nil {
		zclog.Errorln(err)
		return
	}
	tables := make([]string, 0)
	mysqlClient.Raw("show tables").Scan(&tables)
	zclog.Info(tables)

	createSql, err := resources.DBFiles.ReadFile("db/versionctl/create_brood_db_version_ctl.sql")
	if err != nil {
		zclog.Errorln(err)
		return
	}

	zclog.Info(string(createSql))

	mysqlClient.Exec(string(createSql))
	if err := db.Error; err != nil {
		zclog.Errorln(err)
		return
	}
	tables = make([]string, 0)
	mysqlClient.Raw("show tables").Scan(&tables)
	zclog.Info(tables)
}
