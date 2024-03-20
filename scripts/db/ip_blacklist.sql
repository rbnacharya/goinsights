CREATE TABLE `ip_blacklist` (
  `ip` int(11) unsigned NOT NULL,
  PRIMARY KEY (`ip`)
);

INSERT INTO `ip_blacklist` VALUES (0),(2130706433),(4294967295);
