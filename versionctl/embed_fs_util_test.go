package versionctl

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/test/resources"
	"sort"
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
	err := FilledDetailsFromSqlFileName(fileInfo1)
	if err != nil {
		t.Fatal(err)
	}
	jsonFileInfo1, _ := json.Marshal(fileInfo1)
	fmt.Println(string(jsonFileInfo1))

	fmt.Println("---- bsTest_V1.2.3.4_init_test.sql ----")
	fileInfo2 := &EmbedSqlFileInfo{
		Name: "bsTest_V1.2.3.4_init_test.sql",
	}
	err = FilledDetailsFromSqlFileName(fileInfo2)
	if err != nil {
		t.Fatal(err)
	}
	jsonFileInfo2, _ := json.Marshal(fileInfo2)
	fmt.Println(string(jsonFileInfo2))

	fmt.Println("---- bsTest_V1.2.3.4.5_init_test.sql ----")
	fileInfo3 := &EmbedSqlFileInfo{
		Name: "bsTest_V1.2.3.4.5_init_test.sql",
	}
	err = FilledDetailsFromSqlFileName(fileInfo3)
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

func TestGroupAndSort(t *testing.T) {
	allFileInfos, err := ReadEmbedFsByDirName(&resources.DBFilesTest, "db")
	if err != nil {
		t.Fatal(err)
	}

	group := make(map[string][]*EmbedSqlFileInfo)
	for _, fileInfo := range allFileInfos {
		group[fileInfo.BusinessSpace] = append(group[fileInfo.BusinessSpace], fileInfo)
	}

	for bs, subInfos := range group {
		fmt.Printf("业务空间: %s\n", bs)
		fmt.Println("排序前:")
		for _, info := range subInfos {
			fmt.Printf("  sql脚本名: %s\n", info.Name)
		}
		sort.SliceStable(subInfos, func(i, j int) bool {
			infoI := subInfos[i]
			infoJ := subInfos[j]
			if infoI.MajorVersion == infoJ.MajorVersion {
				if infoI.MinorVersion == infoJ.MinorVersion {
					if infoI.PatchVersion == infoJ.PatchVersion {
						return infoI.ExtendVersion < infoJ.ExtendVersion
					} else {
						return infoI.PatchVersion < infoJ.PatchVersion
					}
				} else {
					return infoI.MinorVersion < infoJ.MinorVersion
				}
			} else {
				return infoI.MajorVersion < infoJ.MajorVersion
			}
		})
		fmt.Println("排序后:")
		for _, info := range subInfos {
			fmt.Printf("  sql脚本名: %s\n", info.Name)
		}
	}
}
