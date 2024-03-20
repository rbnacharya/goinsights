CREATE TABLE `customer` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `active` tinyint(1) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
);

INSERT INTO `customer` VALUES (1,'Big News Media Corp',1),(2,'Online Mega Store',1),(3,'Nachoroo Delivery',0),(4,'Euro Telecom Group',1);
