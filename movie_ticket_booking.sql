-- Adminer 4.7.0 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `booked_seats`;
CREATE TABLE `booked_seats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `seat_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  `booked_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `seat_id` (`seat_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `booked_seats_ibfk_1` FOREIGN KEY (`seat_id`) REFERENCES `seats` (`id`),
  CONSTRAINT `booked_seats_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `seats`;
CREATE TABLE `seats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `seat_number` varchar(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `seats` (`id`, `seat_number`) VALUES
(1,	'1A'),
(2,	'1B'),
(3,	'1C'),
(4,	'1D'),
(5,	'1E'),
(6,	'1F'),
(7,	'1G'),
(8,	'1H'),
(9,	'1I'),
(10,	'1J'),
(11,	'1K'),
(12,	'1L'),
(13,	'1M'),
(14,	'1N'),
(15,	'1O'),
(16,	'1P'),
(17,	'1Q'),
(18,	'1R'),
(19,	'1S'),
(20,	'1T'),
(21,	'1U'),
(22,	'1V'),
(23,	'1W'),
(24,	'1X'),
(25,	'1Y'),
(26,	'1Z'),
(27,	'2B'),
(28,	'2C');

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `name` varchar(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- 2020-01-31 12:36:03
