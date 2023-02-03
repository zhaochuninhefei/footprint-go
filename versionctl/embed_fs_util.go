package versionctl

import (
	"embed"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
)

//goland:noinspection GoSnakeCaseUsage
var PTN_SCRIPT_NAME_DEFAULT *regexp.Regexp

//goland:noinspection GoSnakeCaseUsage
var PTN_SCRIPT_NAME_EXTEND *regexp.Regexp

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

func ReadEmbedFsByDirName(embedFs *embed.FS, dirPath string) ([]*EmbedSqlFileInfo, error) {
	files := make([]*EmbedSqlFileInfo, 0)

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
			fileInfo := EmbedSqlFileInfo{
				Name:    name,
				Path:    path,
				Content: string(fileBytes),
			}
			err = AnalyzeDetailsFromSqlFileName(&fileInfo)
			files = append(files, &fileInfo)
		}
	}
	return files, nil
}

func AnalyzeDetailsFromSqlFileName(fileInfo *EmbedSqlFileInfo) error {
	matcherDefault := PTN_SCRIPT_NAME_DEFAULT.FindAllStringSubmatch(fileInfo.Name, -1)
	if len(matcherDefault) > 0 {
		for _, strMatched := range matcherDefault {
			if len(strMatched) == 6 {
				fileInfo.BusinessSpace = strMatched[1]
				major, err := strconv.ParseInt(strMatched[2], 10, 64)
				if err != nil {
					return err
				}
				fileInfo.MajorVersion = major
				minor, err := strconv.ParseInt(strMatched[3], 10, 64)
				if err != nil {
					return err
				}
				fileInfo.MinorVersion = minor
				patch, err := strconv.ParseInt(strMatched[4], 10, 64)
				if err != nil {
					return err
				}
				fileInfo.PatchVersion = patch
				fileInfo.CustomName = strMatched[5]
				fileInfo.Version = fmt.Sprintf("%s_V%d.%d.%d", fileInfo.BusinessSpace, major, minor, patch)
				return nil
			}
		}
	} else {
		matcherExtend := PTN_SCRIPT_NAME_EXTEND.FindAllStringSubmatch(fileInfo.Name, -1)
		for _, strMatched := range matcherExtend {
			if len(strMatched) == 7 {
				fileInfo.BusinessSpace = strMatched[1]
				major, err := strconv.ParseInt(strMatched[2], 10, 64)
				if err != nil {
					return err
				}
				fileInfo.MajorVersion = major
				minor, err := strconv.ParseInt(strMatched[3], 10, 64)
				if err != nil {
					return err
				}
				fileInfo.MinorVersion = minor
				patch, err := strconv.ParseInt(strMatched[4], 10, 64)
				if err != nil {
					return err
				}
				fileInfo.PatchVersion = patch
				extend, err := strconv.ParseInt(strMatched[5], 10, 64)
				if err != nil {
					return err
				}
				fileInfo.ExtendVersion = extend
				fileInfo.CustomName = strMatched[6]
				fileInfo.Version = fmt.Sprintf("%s_V%d.%d.%d.%d", fileInfo.BusinessSpace, major, minor, patch, extend)
				return nil
			}
		}
	}

	return fmt.Errorf("sqlFileName未能正确匹配正则表达式: %s", fileInfo.Name)
}
