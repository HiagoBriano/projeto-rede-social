CREATE DATABASE IF NOT EXISTS `Devbook`
USE `Devbook`;

DROP TABLE IF EXISTS `Users`;

CREATE TABLE `Users` (
  `id` int AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(50) NOT NULL,
  `nickname` varchar(50) NOT NULL UNIQUE,
  `email` varchar(255) NOT NULL UNIQUE,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP()
) ENGINE=InnoDB;