package resources

import "embed"

// DBFiles 嵌入数据库版本控制的SQL
//go:embed db
var DBFiles embed.FS
