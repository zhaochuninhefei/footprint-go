package versionctl

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/test/resources"
	"strings"
	"testing"
)

func TestPrintEmbedFs(t *testing.T) {
	fmt.Println("db/beforeclass 下的文件")
	files, err := ReadEmbedFsByDirName(&resources.DBFilesTest, "db/beforeclass")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test01 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test01")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test02 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test02")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test03 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test03")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test04 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test04")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test05 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test05")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test06 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test06")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		jsonFileInfo, _ := json.Marshal(fileInfo)
		fmt.Println(string(jsonFileInfo))
		//fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.Name, fileInfo.Path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.Path, fileInfo.content)
	}
}

func TestAnalyzeDetailsFromSqlFileName(t *testing.T) {
	fmt.Println("---- bsTest_V1.2.3_init_test.sql ----")
	fileInfo1 := &EmbedSqlFileInfo{
		Name: "bsTest_V1.2.3_init_test.sql",
	}
	err := AnalyzeDetailsFromSqlFileName(fileInfo1)
	if err != nil {
		t.Fatal(err)
	}
	jsonFileInfo1, _ := json.Marshal(fileInfo1)
	fmt.Println(string(jsonFileInfo1))

	fmt.Println("---- bsTest_V1.2.3.4_init_test.sql ----")
	fileInfo2 := &EmbedSqlFileInfo{
		Name: "bsTest_V1.2.3.4_init_test.sql",
	}
	err = AnalyzeDetailsFromSqlFileName(fileInfo2)
	if err != nil {
		t.Fatal(err)
	}
	jsonFileInfo2, _ := json.Marshal(fileInfo2)
	fmt.Println(string(jsonFileInfo2))

	fmt.Println("---- bsTest_V1.2.3.4.5_init_test.sql ----")
	fileInfo3 := &EmbedSqlFileInfo{
		Name: "bsTest_V1.2.3.4.5_init_test.sql",
	}
	err = AnalyzeDetailsFromSqlFileName(fileInfo3)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sqlFileName未能正确匹配正则表达式:") {
			fmt.Println("返回了正确的错误信息")
			jsonFileInfo3, _ := json.Marshal(fileInfo3)
			fmt.Println(string(jsonFileInfo3))
			return
		} else {
			t.Fatal(err)
		}
	} else {
		t.Fatal("匹配了不正确的文件名")
	}
}
