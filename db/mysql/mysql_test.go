package mysql

import (
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"testing"
)

func TestConnectMysqlByDefault(t *testing.T) {
	zclog.Info("===== TestConnectMysqlByDefault 开始 =====")
	RunDBTest()
	zclog.Info("===== TestConnectMysqlByDefault 结束 =====")
}
