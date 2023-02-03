package versionctl

type ScriptResourceMode string

//goland:noinspection GoUnusedConst
const (
	EMBEDFS    ScriptResourceMode = "embedfs"
	FILESYSTEM ScriptResourceMode = "filesystem"
)

type DbVersionCtlProps struct {
	// ScriptResourceMode sql脚本资源类型, embedfs/filesystem, 默认embedfs
	ScriptResourceMode ScriptResourceMode `json:"script_resource_mode" yaml:"script_resource_mode" mapstructure:"script_resource_mode"`

	// ScriptDirs sql脚本文件目录，多个时用","连接。例如："embedfs:db/raven/,embedfs:db/sentry/"
	ScriptDirs string `json:"script_dirs" yaml:"script_dirs" mapstructure:"script_dirs"`

	// BaselineBusinessSpaceAndVersions 数据库非空但首次使用数据库版本管理时，指定生成版本基线的业务空间及其基线版本，多个业务空间时使用逗号连接。
	// 例如:"raven_V1.0.0,sentry_V1.1.2" `json:"script_resource_mode" yaml:"script_resource_mode" mapstructure:"script_resource_mode"`
	BaselineBusinessSpaceAndVersions string `json:"baseline_business_space_and_versions" yaml:"baseline_business_space_and_versions" mapstructure:"baseline_business_space_and_versions"`

	// DbVersionTableName 数据库版本管理表，默认"brood_db_version_ctl"
	DbVersionTableName string `json:"db_version_table_name" yaml:"db_version_table_name" mapstructure:"db_version_table_name"`

	// DbVersionTableCreateSqlPath 数据库版本管理表建表文路径，默认 embedfs:db/versionctl/create_brood_db_version_ctl.sql
	DbVersionTableCreateSqlPath string `json:"db_version_table_create_sql_path" yaml:"db_version_table_create_sql_path" mapstructure:"db_version_table_create_sql_path"`

	// DriverClassName 数据库驱动名, 目前只支持mysql
	DriverClassName string `json:"driver_class_name" yaml:"driver_class_name" mapstructure:"driver_class_name"`

	//// Url 数据库连接URL
	//Url string `json:"url" yaml:"url" mapstructure:"url"`

	// Host 数据库Host
	Host string `json:"host" yaml:"host" mapstructure:"host"`

	// Port 数据库端口
	Port string `json:"port" yaml:"port" mapstructure:"port"`

	// Database 目标数据库
	Database string `json:"database" yaml:"database" mapstructure:"database"`

	// Username 数据库连接用户
	Username string `json:"username" yaml:"username" mapstructure:"username"`

	// Password 数据库连接用户密码
	Password string `json:"password" yaml:"password" mapstructure:"password"`

	// ExistTblQuerySql 查看当前database所有表的sql，默认"show tables"
	ExistTblQuerySql string `json:"exist_tbl_query_sql" yaml:"exist_tbl_query_sql" mapstructure:"exist_tbl_query_sql"`

	// BaselineReset 是否重置数据库基线版本(y/n)
	BaselineReset string `json:"baseline_reset" yaml:"baseline_reset" mapstructure:"baseline_reset"`

	// BaselineResetConditionSql 数据库基线版本重置条件SQL，只有[baselineReset]设置为"y"，且该SQL查询结果非空，才会进行数据库基线版本重置操作
	// 通常建议使用时间戳字段[install_time]作为查询SQL的条件，这样只会生效一次，
	// 以后升级版本时，即使忘记将【baselineReset】属性清除或设置为"n"也不会导致数据库基线版本被误重置。
	BaselineResetConditionSql string `json:"baseline_reset_condition_sql" yaml:"baseline_reset_condition_sql" mapstructure:"baseline_reset_condition_sql"`

	// ModifyDbVersionTable 是否修改DbVersionTable结构(y/n)
	ModifyDbVersionTable string `json:"modify_db_version_table" yaml:"modify_db_version_table" mapstructure:"modify_db_version_table"`

	// ModifyDbVersionTableSqlPath 修改DbVersionTable的SQL
	ModifyDbVersionTableSqlPath string `json:"modify_db_version_table_sql_path" yaml:"modify_db_version_table_sql_path" mapstructure:"modify_db_version_table_sql_path"`
}
