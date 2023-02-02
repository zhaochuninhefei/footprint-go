package utils

import (
	"fmt"
	"github.com/zhaochuninhefei/footprint-go/test/resources"
	"testing"
)

func TestPrintEmbedFs(t *testing.T) {
	files, err := PrintEmbedFs(&resources.DBFilesTest, "db")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}
}
