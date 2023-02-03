package versionctl

type EmbedSqlFileInfo struct {
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
