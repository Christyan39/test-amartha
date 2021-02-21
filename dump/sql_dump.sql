CREATE DATABASE `christyan` /*!40100 DEFAULT CHARACTER SET latin1 */;

CREATE TABLE `shorten` (
  `code` varchar(45) NOT NULL,
  `url` text NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `last_seen_at` timestamp NULL DEFAULT NULL,
  `count` bigint(20) NOT NULL,
  UNIQUE KEY `code_UNIQUE` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
