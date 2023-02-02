package mysql

import (
	"fmt"
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
