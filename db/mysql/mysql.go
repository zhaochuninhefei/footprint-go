package mysql

import (
	"fmt"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"github.com/zhaochuninhefei/footprint-go/resources"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectMysqlByDefault MySQL默认连接
func ConnectMysqlByDefault(dbExist *gorm.DB, host, port, user, pass, dbname string) (*gorm.DB, error) {
	if dbExist != nil {
		return dbExist, nil
	}
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

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
