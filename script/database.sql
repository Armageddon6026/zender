-- --------------------------------------------------------
-- 主機:                           127.0.0.1
-- 伺服器版本:                        11.3.2-MariaDB - mariadb.org binary distribution
-- 伺服器作業系統:                      Win64
-- HeidiSQL 版本:                  12.8.0.6917
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 傾印 zender 的資料庫結構
CREATE DATABASE IF NOT EXISTS `zender` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `zender`;

-- 傾印  資料表 zender.service_group 結構
CREATE TABLE IF NOT EXISTS `service_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `note` text DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在傾印表格  zender.service_group 的資料：~8 rows (近似值)
DELETE FROM `service_group`;
INSERT INTO `service_group` (`id`, `name`, `note`) VALUES
	(1, 'K99', '雄安新村'),
	(2, 'CAT', '貓貓'),
	(4, 'DOG', '狗'),
	(5, 'JOJO', 'I reject my humanity Jojo'),
	(7, 'Apple', '高雄淹大水發大財'),
	(9, 'GroupZZ', '測試用\n群組'),
	(12, 'dd', 'aaaa'),
	(14, 'TEST2s', 'TTT2');

-- 傾印  資料表 zender.service_infomation 結構
CREATE TABLE IF NOT EXISTS `service_infomation` (
  `sn` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL DEFAULT 'Unknown',
  `privateKey` varchar(50) NOT NULL,
  `groupId` int(11) NOT NULL DEFAULT -1,
  PRIMARY KEY (`sn`) USING BTREE,
  UNIQUE KEY `UUID` (`uuid`,`privateKey`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在傾印表格  zender.service_infomation 的資料：~9 rows (近似值)
DELETE FROM `service_infomation`;
INSERT INTO `service_infomation` (`sn`, `uuid`, `name`, `privateKey`, `groupId`) VALUES
	(1, '52d1fdaa-743a-45a8-b3b3-927b78399e64', 'YuXi', '50,L-V|.pUr1Dyl+e48S', -1),
	(34, 'dc002e16-7961-49c6-bf1c-d1702e89358f', 'TEST0', 'P8i2g@QE{*9q_|4UVo3b', -1),
	(66, 'a7dc3769-cc9c-4487-ada0-ce956a3686e0', 'TEST1', ')@F40YH>y15z9;sTd*qV', -1),
	(67, '57c3b9b1-cce8-425a-afb4-5381d1ec61bb', 'TEST2', 'ek9![50UyO{2fFPG%h&7', 2),
	(68, 'd2fe4cf3-c4d4-431f-838b-d4edcee39302', 'TEST3', 'g4$^iw8]Q30JnY+E6R*q', 1),
	(69, 'da62552b-1498-4f3c-9af8-1d791ad2bbb5', 'TEST4', '6y$5*j3J{IdAm2U|K[h9', 2),
	(70, '568bbfbc-095e-4a22-a408-f6971de152b7', 'TEST6', 'iz7a62UyF#4,<LGR@0n*', 1),
	(72, '427d2df1-8e7c-467d-a0e6-1f11df067314', '77', ']x5qRr49$7UJ1_CfQ&:c', -1),
	(74, 'df709d23-9140-4bf0-8952-df057249bc14', 'TEST97', '1=K3d9*Hf27>F+Qzp&bL', -1);

-- 傾印  資料表 zender.users 結構
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'Unknown',
  `date` timestamp NOT NULL DEFAULT current_timestamp(),
  `role` tinyint(4) NOT NULL DEFAULT 0,
  `auth` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account` (`account`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- 正在傾印表格  zender.users 的資料：~3 rows (近似值)
DELETE FROM `users`;
INSERT INTO `users` (`id`, `account`, `password`, `name`, `date`, `role`, `auth`) VALUES
	(1, 'Zak', '12345', 'Zak Liu', '2024-05-17 11:06:02', 1, 1),
	(2, 'Rudi', 'test2', 'Rudi Parsen', '2024-05-17 11:06:02', 0, 1),
	(3, 'test', '55677', '測試', '2024-05-17 11:06:02', 0, 0);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
