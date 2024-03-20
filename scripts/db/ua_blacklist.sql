CREATE TABLE `ua_blacklist` (
  `ua` varchar(255) NOT NULL,
  PRIMARY KEY (`ua`)
);

INSERT INTO `ua_blacklist` VALUES ('A6-Indexer'),('Googlebot-News'),('Googlebot');
