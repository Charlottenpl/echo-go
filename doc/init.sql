-- 创建数据库
CREATE DATABASE echo;
USE echo;

-- 创建blog-tag链接表
CREATE TABLE `blog_tag`  (
  `blog_id` int NOT NULL,
  `tag_id` int NOT NULL,
  PRIMARY KEY (`blog_id`, `tag_id`)
);

-- 创建blog-type链接表
CREATE TABLE `blog_type`  (
  `blog_id` int NOT NULL,
  `type_id` int NOT NULL,
  PRIMARY KEY (`blog_id`, `type_id`)
);


-- 创建表格
CREATE TABLE `blog`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NULL,
  `pic` varchar(255) NULL,
  `content` TEXT NULL,
  `type` varchar(255) NULL,
  `create_time` datetime NULL,
  `update_time` datetime NULL,
  `click_num` int NOT NULL DEFAULT 0,
  `status` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

-- 创建友链表
CREATE TABLE `links`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NULL,
  `pic` varchar(255) NULL,
  `url` varchar(255) NULL,
  `nick_name` varchar(255) NULL,
  `status` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

CREATE TABLE `log`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user` varchar(255) NULL,
  `article_id` int NULL,
  `tag_id` int NULL,
  `type_id` int NULL,
  `time` datetime NULL,
  `level` varchar(255) NULL,
  `log` varchar(255) NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `tag`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NULL,
  `status` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

CREATE TABLE `type`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NULL,
  `status` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '唯一外键',
  `nick_name` varchar(255) NULL,
  `pic` varchar(255) NULL,
  `password` varchar(255) NULL,
  `roles` varchar(255) NULL,
  `create_at` varchar(255) NULL,
  `status` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);


CREATE TRIGGER blog_create
BEFORE INSERT
ON blog FOR EACH ROW
BEGIN
SET new.create_time = CURRENT_TIMESTAMP;
SET new.update_time = CURRENT_TIMESTAMP;
END;



CREATE TRIGGER blog_update
BEFORE UPDATE
ON blog FOR EACH ROW
SET new.update_time = CURRENT_TIMESTAMP;


-- DROP TRIGGER blog_update
SELECT * FROM blog WHERE id=1