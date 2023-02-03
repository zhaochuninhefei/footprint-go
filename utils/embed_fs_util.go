package utils

import (
	"embed"
	"path/filepath"
)

type EmbedSqlFileInfo struct {
	// sql脚本文件名
	//  格式为"[业务空间]_V[major].[minor].[patch].[extend]_[自定义名称].sql"
	//  其中，`.[extend]`可以省略，如: "smtp_V1.0.0_init.sql"
	name string
	// sql脚本文件访问路径，从embed.FS的根目录开始的完整访问路径，如:"db/test01/smtp_V1.0.0_init.sql"
	path string
	// sql脚本文件内容
	content string
	// 业务空间
	//  用于同一database下数据表的集合划分，通常根据业务功能划分；
	//  同一个业务空间中的表的版本管理采用统一的版本号递增顺序，不同业务空间的版本号的递增顺序是不同的;
	//  业务空间命名只支持大小写字母与数字。
	businessSpace string
	// 主版本号，一个业务空间对应的主版本号，对应"x.y.z.t"中的x，只支持非负整数
	majorVersion int64
	// 次版本号，一个业务空间对应的次版本号，对应"x.y.z.t"中的y，只支持非负整数
	minorVersion int64
	// 补丁版本号，一个业务空间对应的补丁版本号，对应"x.y.z.t"中的z，只支持非负整数
	patchVersion int64
	// 扩展版本号，一个业务空间对应的扩展版本号，对应"x.y.z.t"中的4，只支持非负整数
	extendVersion int64
	// 一个业务空间的完整版本号，格式为"[businessSpace]_V[majorVersion].[minorVersion].[patchVersion]"
	version string
	// 该sql脚本的自定义名称，支持大小写字母，数字与下划线
	customName string
}

func ReadEmbedFsByDirName(embedFs *embed.FS, dirPath string) ([]EmbedSqlFileInfo, error) {
	files := make([]EmbedSqlFileInfo, 0)

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
			files = append(files, EmbedSqlFileInfo{
				name:    name,
				path:    path,
				content: string(fileBytes),
			})
		}
	}
	return files, nil
}
