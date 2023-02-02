package resources

import (
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"testing"
)

func TestScanDBFilesTest(t *testing.T) {
	zclog.Info("===== TestScanDBFilesTest 开始 =====")

	PrintDBFilesTest("db")

	zclog.Info("===== TestScanDBFilesTest 结束 =====")
}
