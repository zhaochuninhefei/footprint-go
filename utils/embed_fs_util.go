package utils

import (
	"embed"
	"path/filepath"
)

type EmbedFileInfo struct {
	name    string // 文件名
	path    string // 访问路径
	content string // 文件内容
}

func ReadEmbedFsByDirName(embedFs *embed.FS, dirPath string) ([]EmbedFileInfo, error) {
	files := make([]EmbedFileInfo, 0)

	entries, err := embedFs.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(dirPath, name)
		isDir := entry.IsDir()
		if isDir {
			subFiles, err := ReadEmbedFsByDirName(embedFs, path)
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			fileBytes, err := embedFs.ReadFile(path)
			if err != nil {
				return nil, err
			}
			files = append(files, EmbedFileInfo{
				name:    name,
				path:    path,
				content: string(fileBytes),
			})
		}
	}
	return files, nil
}
