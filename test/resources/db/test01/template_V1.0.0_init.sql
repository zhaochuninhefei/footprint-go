DROP TABLE IF EXISTS `rv_mail_templates`;
CREATE TABLE IF NOT EXISTS `rv_mail_templates` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '邮件模板ID',
  `name` VARCHAR(255) NOT NULL COMMENT '模板名称',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '邮件模板';

DROP TABLE IF EXISTS `rv_imgs`;
CREATE TABLE IF NOT EXISTS `rv_imgs` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '邮件图片ID',
  `img_cid` VARCHAR(45) NOT NULL COMMENT '图片cid',
  `img_file_name` VARCHAR(255) NOT NULL COMMENT '图片文件名',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `rv_imgs_unique_001` (`img_cid` ASC),
  UNIQUE INDEX `rv_imgs_unique_002` (`img_file_name` ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '邮件图片表';

DROP TABLE IF EXISTS `rv_temp_img_map`;
CREATE TABLE IF NOT EXISTS `rv_temp_img_map` (
  `tmp_id` INT UNSIGNED NOT NULL COMMENT '邮件模板ID',
  `img_id` INT UNSIGNED NOT NULL COMMENT '邮件图片ID',
  PRIMARY KEY (`tmp_id`, `img_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '邮件模板图片映射表';

INSERT INTO `rv_mail_templates` (`name`) VALUES ('warpgate');
