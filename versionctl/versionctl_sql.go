package versionctl

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// SqlScriptInfo SQL脚本信息
type SqlScriptInfo struct {
	// sql脚本文件名
	//  格式为"[业务空间]_V[major].[minor].[patch].[extend]_[自定义名称].sql"
	//  其中，`.[extend]`可以省略，如: "smtp_V1.0.0_init.sql"
	Name string `json:"name"`
	// sql脚本文件访问路径，从embed.FS的根目录开始的完整访问路径，如:"db/test01/smtp_V1.0.0_init.sql"
	Path string `json:"path"`
	// sql脚本文件内容
	Content string `json:"content"`
	// 业务空间
	//  用于同一database下数据表的集合划分，通常根据业务功能划分；
	//  同一个业务空间中的表的版本管理采用统一的版本号递增顺序，不同业务空间的版本号的递增顺序是不同的;
	//  业务空间命名只支持大小写字母与数字。
	BusinessSpace string `json:"business_space"`
	// 主版本号，一个业务空间对应的主版本号，对应"x.y.z.t"中的x，只支持非负整数
	MajorVersion int64 `json:"major_version"`
	// 次版本号，一个业务空间对应的次版本号，对应"x.y.z.t"中的y，只支持非负整数
	MinorVersion int64 `json:"minor_version"`
	// 补丁版本号，一个业务空间对应的补丁版本号，对应"x.y.z.t"中的z，只支持非负整数
	PatchVersion int64 `json:"patch_version"`
	// 扩展版本号，一个业务空间对应的扩展版本号，对应"x.y.z.t"中的4，只支持非负整数
	ExtendVersion int64 `json:"extend_version"`
	// 一个业务空间的完整版本号，格式为"[businessSpace]_V[majorVersion].[minorVersion].[patchVersion]"
	Version string `json:"version"`
	// 该sql脚本的自定义名称，支持大小写字母，数字与下划线
	CustomName string `json:"custom_name"`
}

// SqlScriptFilter SQL脚本过滤条件
type SqlScriptFilter struct {
	// 业务空间
	BusinessSpace string `json:"business_space"`
	// 主版本号
	MajorVersion int64 `json:"major_version"`
	// 次版本号
	MinorVersion int64 `json:"minor_version"`
	// 补丁版本号
	PatchVersion int64 `json:"patch_version"`
	// 扩展版本号
	ExtendVersion int64 `json:"extend_version"`
}

// PTN_SCRIPT_NAME_DEFAULT sql脚本文件名正则表达式(默认)
//goland:noinspection GoSnakeCaseUsage
var PTN_SCRIPT_NAME_DEFAULT *regexp.Regexp

// PTN_SCRIPT_NAME_EXTEND sql脚本文件名正则表达式(带扩展版本号)
//goland:noinspection GoSnakeCaseUsage
var PTN_SCRIPT_NAME_EXTEND *regexp.Regexp

// init 初始化sql脚本文件名正则表达式
func init() {
	PTN_SCRIPT_NAME_DEFAULT = regexp.MustCompile("^([A-Za-z0-9]+)_V(\\d+)\\.(\\d+)\\.(\\d+)_(\\w+)\\.sql$")
	if PTN_SCRIPT_NAME_DEFAULT == nil {
		panic("正则表达式不正确: ^([A-Za-z0-9]+)_V(\\d+)\\.(\\d+)\\.(\\d+)_(\\w+)\\.sql$")
	}
	PTN_SCRIPT_NAME_EXTEND = regexp.MustCompile("^([A-Za-z0-9]+)_V(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+)_(\\w+)\\.sql$")
	if PTN_SCRIPT_NAME_EXTEND == nil {
		panic("正则表达式不正确: ^([A-Za-z0-9]+)_V(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+)_(\\w+)\\.sql$")
	}
}

// ReadEmbedSql 读取嵌入文件目录下的SQL文件(包括子目录)
//  @param embedFs 嵌入FS资源, 如根目录为`db`的嵌入FS
//  @param dirPath 访问目录路径, 如:`db/footprint`
//  @param filter SQL脚本过滤条件
//  @return []*SqlScriptInfo 嵌入SQL文件信息结构体数组(切片)
//  @return error
//func ReadEmbedSql(embedFs *embed.FS, dirPath string, filters map[string]SqlScriptFilter) ([]*SqlScriptInfo, error) {
//	files := make([]*SqlScriptInfo, 0)
//
//	entries, err := embedFs.ReadDir(dirPath)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, entry := range entries {
//		name := entry.Name()
//		path := filepath.Join(dirPath, name)
//		isDir := entry.IsDir()
//		if isDir {
//			subFiles, err := ReadEmbedSql(embedFs, path, filters)
//			if err != nil {
//				return nil, err
//			}
//			files = append(files, subFiles...)
//		} else {
//			fileInfo, err := createFileInfo(name, path)
//			if err != nil {
//				return nil, err
//			}
//			if filters != nil {
//				sqlFilter := filters[fileInfo.BusinessSpace]
//				// 比较当前脚本的版本是否是增量
//				if filterIncreaseFileInfoByVersions(fileInfo, sqlFilter) {
//					fileBytes, err := embedFs.ReadFile(path)
//					if err != nil {
//						return nil, err
//					}
//					fileInfo.Content = string(fileBytes)
//					files = append(files, &fileInfo)
//				}
//			} else {
//				fileBytes, err := embedFs.ReadFile(path)
//				if err != nil {
//					return nil, err
//				}
//				fileInfo.Content = string(fileBytes)
//				files = append(files, &fileInfo)
//			}
//		}
//	}
//	return files, nil
//}

func ReadSql(reader SqlScriptReader, dirPath string, filters map[string]SqlScriptFilter) ([]*SqlScriptInfo, error) {
	files := make([]*SqlScriptInfo, 0)

	entries, err := reader.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(dirPath, name)
		isDir := entry.IsDir()
		if isDir {
			subFiles, err := ReadSql(reader, path, filters)
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			fileInfo, err := createFileInfo(name, path)
			if err != nil {
				return nil, err
			}
			if filters != nil {
				sqlFilter := filters[fileInfo.BusinessSpace]
				// 比较当前脚本的版本是否是增量
				if filterIncreaseFileInfoByVersions(fileInfo, sqlFilter) {
					fileBytes, err := reader.ReadFile(path)
					if err != nil {
						return nil, err
					}
					fileInfo.Content = string(fileBytes)
					files = append(files, &fileInfo)
				}
			} else {
				fileBytes, err := reader.ReadFile(path)
				if err != nil {
					return nil, err
				}
				fileInfo.Content = string(fileBytes)
				files = append(files, &fileInfo)
			}
		}
	}

	return files, nil
}

// createFileInfo 根据SQL文件名填充细节
//  @param fileInfo 嵌入SQL文件信息结构体
//  @return error
func createFileInfo(name string, path string) (SqlScriptInfo, error) {
	fileInfo := SqlScriptInfo{
		Name: name,
		Path: path,
	}
	// 优先使用默认的正则表达式解析SQL文件名
	matcherDefault := PTN_SCRIPT_NAME_DEFAULT.FindAllStringSubmatch(name, -1)
	if len(matcherDefault) > 0 {
		for _, strMatched := range matcherDefault {
			if len(strMatched) == 6 {
				fileInfo.BusinessSpace = strMatched[1]
				major, err := strconv.ParseInt(strMatched[2], 10, 64)
				if err != nil {
					return fileInfo, err
				}
				fileInfo.MajorVersion = major
				minor, err := strconv.ParseInt(strMatched[3], 10, 64)
				if err != nil {
					return fileInfo, err
				}
				fileInfo.MinorVersion = minor
				patch, err := strconv.ParseInt(strMatched[4], 10, 64)
				if err != nil {
					return fileInfo, err
				}
				fileInfo.PatchVersion = patch
				fileInfo.CustomName = strMatched[5]
				fileInfo.Version = fmt.Sprintf("%s_V%d.%d.%d", fileInfo.BusinessSpace, major, minor, patch)
				return fileInfo, nil
			}
		}
	} else {
		// 使用带扩展版本号的正则表达式解析SQL文件名
		matcherExtend := PTN_SCRIPT_NAME_EXTEND.FindAllStringSubmatch(fileInfo.Name, -1)
		for _, strMatched := range matcherExtend {
			if len(strMatched) == 7 {
				fileInfo.BusinessSpace = strMatched[1]
				major, err := strconv.ParseInt(strMatched[2], 10, 64)
				if err != nil {
					return fileInfo, err
				}
				fileInfo.MajorVersion = major
				minor, err := strconv.ParseInt(strMatched[3], 10, 64)
				if err != nil {
					return fileInfo, err
				}
				fileInfo.MinorVersion = minor
				patch, err := strconv.ParseInt(strMatched[4], 10, 64)
				if err != nil {
					return fileInfo, err
				}
				fileInfo.PatchVersion = patch
				extend, err := strconv.ParseInt(strMatched[5], 10, 64)
				if err != nil {
					return fileInfo, err
				}
				fileInfo.ExtendVersion = extend
				fileInfo.CustomName = strMatched[6]
				fileInfo.Version = fmt.Sprintf("%s_V%d.%d.%d.%d", fileInfo.BusinessSpace, major, minor, patch, extend)
				return fileInfo, nil
			}
		}
	}
	// 解析失败
	return fileInfo, fmt.Errorf("sqlFileName未能正确匹配正则表达式: %s", fileInfo.Name)
}

func filterIncreaseFileInfoByVersions(fileInfo SqlScriptInfo, filter SqlScriptFilter) bool {
	if fileInfo.MajorVersion == filter.MajorVersion {
		if fileInfo.MinorVersion == filter.MinorVersion {
			if fileInfo.PatchVersion == filter.PatchVersion {
				return fileInfo.ExtendVersion > filter.ExtendVersion
			}
			return fileInfo.PatchVersion > filter.PatchVersion
		}
		return fileInfo.MinorVersion > filter.MinorVersion
	}
	return fileInfo.MajorVersion > filter.MajorVersion
}

type SqlScriptReader interface {
	ReadDir(dirPath string) ([]fs.DirEntry, error)
	ReadFile(filePath string) ([]byte, error)
}

type FSSqlReader struct {
}

func (f *FSSqlReader) ReadDir(dirPath string) ([]fs.DirEntry, error) {
	return os.ReadDir(dirPath)
}

func (f *FSSqlReader) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
