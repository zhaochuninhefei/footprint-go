package utils

import "strings"

// ReadSqls 读取SQL脚本文件
//  @param sqlScript SQL脚本文件内容
//  @return []string SQL语句集合
func ReadSqls(sqlScript string) []string {
	sqls := make([]string, 0)
	// 按行读取SQL内容
	lines := splitLines(sqlScript)
	// sql构造器
	var sqlBuilder strings.Builder
	for _, line := range lines {
		// 去除首尾空白字符
		lineTrim := strings.TrimSpace(line)
		// 去除空行、注释行
		if lineTrim == "" || strings.HasPrefix(lineTrim, "--") {
			continue
		}
		if strings.HasSuffix(lineTrim, ";") {
			// 如果该行以";"结尾，则认为该条sql语句结束
			// 先去除末尾分号，将该行加入sql构造器
			sqlBuilder.WriteString(lineTrim[:len(lineTrim)-1])
			// 将sql构造器转为sql语句，加入sql语句集合
			sqls = append(sqls, sqlBuilder.String())
			// 重置sql构造器
			sqlBuilder.Reset()
		} else {
			// 如果该行没有以";"结尾，则认为该条sql语句尚未结束
			sqlBuilder.WriteString(lineTrim)
			sqlBuilder.WriteString(" \n")
		}
	}
	// 特殊场景处理：如果sql脚本最后一条sql语句没有写";"结尾，则需要将非空的sql构造器转为sql语句并加入sql语句集合。
	if sqlBuilder.Len() > 0 {
		sqls = append(sqls, strings.TrimSpace(sqlBuilder.String()))
		sqlBuilder.Reset()
	}
	return sqls
}

func splitLines(sqlScript string) []string {
	if sqlScript == "" {
		return make([]string, 0)
	}
	return strings.Split(strings.ReplaceAll(sqlScript, "\r\n", "\n"), "\n")
}
