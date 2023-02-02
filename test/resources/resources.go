package resources

import (
	"embed"
	"fmt"
	"path/filepath"
)

//go:embed db
var DBFilesTest embed.FS

func PrintDBFilesTest(dirName string) {
	entries, err := DBFilesTest.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	fmt.Println("dir: ", dirName)
	for _, entry := range entries {
		name := entry.Name()
		isDir := entry.IsDir()
		info, _ := entry.Info()
		fmt.Printf("name: %s, 是否目录: %v, 大小: %d bytes\n", name, isDir, info.Size())
		if isDir {
			PrintDBFilesTest(filepath.Join(dirName, name))
		}
	}
}
