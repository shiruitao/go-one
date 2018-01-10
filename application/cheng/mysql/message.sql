CREATE DATABASE IF NOT EXISTS `blog`;
USE `blog`;

CREATE TABLE IF NOT EXISTS `message` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL,
  `content` text,
  `state` int NOT NULL DEFAULT '1',
  `label` varchar(64) NOT NULL DEFAULT '',
  `created` datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARACTER SET latin1 AUTO_INCREMENT=1000;

CREATE TABLE IF NOT EXISTS `admin` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) UNIQUE  NOT NULL,
  `password` varchar(128) NOT NULL,
  `state` int(8) DEFAULT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARACTER SET latin1 AUTO_INCREMENT=1;