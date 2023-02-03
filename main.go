package main

import (
	"gitee.com/zhaochuninhefei/footprint-go/db/mysql"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
)

func main() {
	zclog.Info("footprint-go main start...")

	mysql.RunDBTest()

	zclog.Info("footprint-go main stop...")
}
