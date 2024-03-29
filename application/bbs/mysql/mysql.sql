CREATE TABLE IF NOT EXISTS `user` (
  `id` INT(64) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL UNIQUE,
  `nickname` VARCHAR(128) NOT NULL,
  `avatar` VARCHAR(128),
  `sex` VARCHAR(16),
  `school` VARCHAR(128) NOT NULL,
  `realname` VARCHAR(128) NOT NULL,
  `password` VARCHAR(128) NOT NULL,
  `isadmin` BOOLEAN NOT NULL,
  `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `lastlogin` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `isactive` BOOLEAN DEFAULT TRUE,
  PRIMARY KEY (`id`, `name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `article` (
  `id` INT(64) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(256) NOT NULL,
  `content` TEXT,
  `authorid` INT(64) NOT NULL,
  `image1` VARCHAR(256),
  `image2` VARCHAR(256),
  `image3` VARCHAR(256),
  `video`  VARCHAR(256),
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `isactive` BOOLEAN DEFAULT TRUE,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `comment` (
  `id` INT(64) UNSIGNED NOT NULL AUTO_INCREMENT,
  `artid` VARCHAR(256) NOT NULL,
  `content` TEXT,
  `avatar` VARCHAR(256),
  `creatorid` INT(64) NOT NULL,
  `creator` VARCHAR(256),
  `repliedid` VARCHAR(256),
  `replied` VARCHAR(256),
  `file`  VARCHAR(256),
  `repfile` VARCHAR(256),
  `repcontent` VARCHAR(256),
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `isactive` BOOLEAN DEFAULT TRUE,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
