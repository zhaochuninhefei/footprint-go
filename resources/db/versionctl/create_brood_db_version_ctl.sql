CREATE TABLE IF NOT EXISTS `brood_db_version_ctl` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据库版本ID',
    `business_space` VARCHAR(50) NOT NULL COMMENT '业务空间',
    `major_version` INT NOT NULL COMMENT '主版本号',
    `minor_version` INT NOT NULL COMMENT '次版本号',
    `patch_version` INT NOT NULL COMMENT '补丁版本号',
    `extend_version` INT NOT NULL DEFAULT 0 COMMENT '扩展版本号',
    `version` VARCHAR(50) NOT NULL COMMENT '版本号,V[major].[minor].[patch].[extend_version]',
    `custom_name` VARCHAR(50) NOT NULL DEFAULT 'none' COMMENT '脚本自定义名称',
    `version_type` VARCHAR(10) NOT NULL COMMENT '版本类型:SQL/BaseLine',
    `script_file_name` VARCHAR(200) NOT NULL DEFAULT 'none' COMMENT '脚本文件名',
    `script_digest_hex` VARCHAR(200) NOT NULL DEFAULT 'none' COMMENT '脚本内容摘要(16进制)',
    `success` TINYINT NOT NULL COMMENT '是否执行成功',
    `execution_time` INT NOT NULL COMMENT '脚本安装耗时',
    `install_time` VARCHAR(19) NOT NULL COMMENT '脚本安装时间,格式:[yyyy-MM-dd HH:mm:ss]',
    `install_user` VARCHAR(100) NOT NULL COMMENT '脚本安装用户',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `brood_db_version_ctl_unique01` (`business_space`, `major_version`, `minor_version`, `patch_version`, `extend_version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '数据库版本控制表';
