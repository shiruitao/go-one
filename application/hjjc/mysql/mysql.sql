CREATE TABLE IF NOT EXISTS `user` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NULL,
  `password`  varchar(20) NOT NULL,
  `realname` VARCHAR(20),
  `userrole` VARCHAR(128) NOT NULL UNIQUE,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `company` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20),
  `divisions`  varchar(128) NOT NULL,
  `jd` VARCHAR(128) NOT NULL,
  `wd` VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `device` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20),
  `device`  VARCHAR(128) NOT NULL,
  `datatime` VARCHAR(128) NOT NULL,
  `Aa` FLOAT(7, 3),
  `Ab` FLOAT(7, 3),
  `Ac` FLOAT(7, 3),
  `Va` FLOAT(7, 3),
  `Vb` FLOAT(7, 3),
  `Vc` FLOAT(7, 3),
  `Wa` FLOAT(7, 4),
  `Wb` FLOAT(7, 4),
  `Wc` FLOAT(7, 4),
  `area` VARCHAR(256),
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
