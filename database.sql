CREATE TABLE `Product` (
  `id` INTEGER NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50),
  `price` DECIMAL(12,2),
  PRIMARY KEY (`id`)
);
INSERT INTO `Product` (`name`, `price`) VALUES ('Mobile', 100);
INSERT INTO `Product` (`name`, `price`) VALUES ('Tablet', 200);
INSERT INTO `Product` (`name`, `price`) VALUES ('Labtop', 300.00);
INSERT INTO `Product` (`name`, `price`) VALUES ('Desktop', 400);
INSERT INTO `Product` (`name`, `price`) VALUES ('Server', 500);