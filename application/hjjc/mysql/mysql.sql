CREATE TABLE IF NOT EXISTS `user` (
  `UserID` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `UserName` varchar(20),
  `Userpwd`  varchar(128) NOT NULL,
  `RealName` varchar(128),
  `UserRole` VARCHAR(128),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `company` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20),
  `divisions`  varchar(128) NOT NULL,
  `jd` decimal(10, 7) NOT NULL,
  `wd` decimal(10, 7) NOT NULL,
  `area` VARCHAR(256),
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

CREATE TABLE IF NOT EXISTS `devices` (
  `id` INT(32) unsigned NOT NULL AUTO_INCREMENT,
  `DviceID` INT(32),
  `SetupDate` VARCHAR(128),
  `DeviceName` VARCHAR(128),
  `DivisionID` INT(32),
  `DeviceAddr` VARCHAR(128),
  `CompanyName` VARCHAR(128),
  `DeviceType` VARCHAR(128),
  `LastDataTime` VARCHAR(128),
  `CompanyID` INT(32),
  `phaseACurrent` FLOAT(7,3),
  `phaseBCurrent` FLOAT(7,3),
  `phaseCCurrent` FLOAT(7,3),
  `CTRatio` FLOAT(7,3),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `historydata` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT,
  `DeviceID` INT(32)
  `DataTimeHour`  varchar(128),
  `CurrentA` FLOAT(7,3),
  `CurrentB` FLOAT(7,3),
  `CurrentC` FLOAT(7,3),
  `VoltA` FLOAT(7,3),
  `VoltB` FLOAT(7,3),
  `VoltC` FLOAT(7,3),
  `isAvailable` BOOLEAN DEFAULT TRUE,
  `isCalculate` BOOLEAN DEFAULT FALSE,
  `EnergyA` DEFAULT NULL,
  `EnergyB` DEFAULT NULL,
  `EnergyC` DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
