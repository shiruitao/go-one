CREATE TABLE IF NOT EXISTS `user` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NULL,
  `sex`  varchar(20) NOT NULL,
  `age` INT(8) NOT NULL,
  `area` VARCHAR(128),
  `numberid` VARCHAR(128) NOT NULL UNIQUE,
  `address` VARCHAR(256) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`, `numberid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `economic` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `year` int(20),
  `area` VARCHAR(128),
  `count` int(20) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `industrial` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `manufacture` int(20),
  `build` VARCHAR(128),
  `retail` int(20) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `company` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128),
  `type` VARCHAR(128),
  `Capital` int(20) NOT NULL,
  `managescope` VARCHAR(128),
  `address` VARCHAR(128),
  `area` VARCHAR(128),
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `region` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128),
  `address` VARCHAR(128),
  `acreage` VARCHAR(128),
  `developers` VARCHAR(128),
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `house` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128),
  `layer` INT(4) NOT NULL,
  `high` FLOAT(20,2) NOT NULL,
  `acreage` FLOAT(128,2),
  `address` VARCHAR(128),
  `date` VARCHAR(128),
  `region` VARCHAR(128),
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `notice` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(128),
  `content` VARCHAR(256),
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
