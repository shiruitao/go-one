CREATE DATABASE IF NOT EXISTS `blog`;
USE `blog`;
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `message` (
  `id` int(64) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL,
  `content` text,
  `state` int NOT NULL DEFAULT '1',
  `label1` varchar(64) NOT NULL DEFAULT '',
  `label2` varchar(64) NOT NULL DEFAULT '',
  `created` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `admin` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) UNIQUE  NOT NULL,
  `password` varchar(128) NOT NULL,
  `state` int(8) DEFAULT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `comments` (
  `id` int(64) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(16) unsigned NOT NULL,
  `message_id` int(64) unsigned NOT NULL,
  `comment` text NOT NULL,
  `status` int(8),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) UNIQUE DEFAULT NULL,
  `password` varchar(128) NOT NULL,
  `sex` TINYINT(1) DEFAULT NULL COMMENT '0:男;1:女',
  `phone` int(16) UNIQUE NOT NULL,
  `email` varchar(64),
  `status` int(8) NOT NULL,
  `created` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ---------------------------------------------------------------------
