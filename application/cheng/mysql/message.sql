--
CREATE DATABASE IF NOT EXISTS `blog`;
USE `blog`;
-- ----------------------------------------------

CREATE TABLE IF NOT EXISTS `message` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL,
  `content` text,
  `state` int NOT NULL DEFAULT '1',
  `label` varchar(64) NOT NULL DEFAULT '',
  `created` datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ----------------------------------------------