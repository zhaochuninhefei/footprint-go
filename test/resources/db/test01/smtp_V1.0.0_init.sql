DROP TABLE IF EXISTS `rv_smtps`;
CREATE TABLE IF NOT EXISTS `rv_smtps` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'smtp服务器ID',
  `smtp_host` VARCHAR(200) NOT NULL COMMENT 'smtp服务器',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `smtp_host_UNIQUE` (`smtp_host` ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = 'smtp列表';

DROP TABLE IF EXISTS `rv_smtp_conninfos`;
CREATE TABLE IF NOT EXISTS `rv_smtp_conninfos` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'smtp服务器连接信息ID',
  `mail_address` varchar(100) NOT NULL COMMENT '注册邮箱地址',
  `smtp_host` varchar(200) NOT NULL COMMENT 'smtp服务器host',
  `smtp_port` INT NOT NULL DEFAULT 25 COMMENT 'smtp服务器port',
  `username` varchar(45) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mail_address_UNIQUE` (`mail_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = 'smtp服务器连接信息';

