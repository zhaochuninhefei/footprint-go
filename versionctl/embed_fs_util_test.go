package versionctl

import (
	"fmt"
	"gitee.com/zhaochuninhefei/footprint-go/test/resources"
	"testing"
)

func TestPrintEmbedFs(t *testing.T) {
	fmt.Println("db/beforeclass 下的文件")
	files, err := ReadEmbedFsByDirName(&resources.DBFilesTest, "db/beforeclass")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test01 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test01")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test02 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test02")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test03 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test03")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test04 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test04")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test05 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test05")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db/test06 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db/test06")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}

	fmt.Println()
	fmt.Println("db 下的文件")
	files, err = ReadEmbedFsByDirName(&resources.DBFilesTest, "db")
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		fmt.Printf("文件名: %s, 访问路径: %s\n", fileInfo.name, fileInfo.path)
		//fmt.Printf("文件名: %s, 访问路径: %s, \n文件内容:\n----------\n%s\n----------\n", fileInfo.name, fileInfo.path, fileInfo.content)
	}
}
