-- MySQL dump 10.13  Distrib 8.0.32, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: carbon
-- ------------------------------------------------------
-- Server version       8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `assets_sell`
--
USE carbon;
DROP TABLE IF EXISTS `assets_sell`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assets_sell` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `ex_id` varchar(255) NOT NULL COMMENT '交易单号',
                               `ass_id` varchar(255) NOT NULL COMMENT '资产编号',
                               `country` varchar(5) NOT NULL COMMENT '资产所属国家',
                               `collection_wallet_id` varchar(255) NOT NULL COMMENT '收款的钱包 id',
                               `user_id` varchar(255) NOT NULL COMMENT '资产所属用户',
                               `source` varchar(255) NOT NULL COMMENT '来源资产',
                               `amount` double NOT NULL COMMENT 'number(数量)/amount(金额)',
                               `number` int NOT NULL COMMENT '数量',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易创建时间',
                               `end_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易结束时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets_sell`
--

LOCK TABLES `assets_sell` WRITE;
/*!40000 ALTER TABLE `assets_sell` DISABLE KEYS */;
/*!40000 ALTER TABLE `assets_sell` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-09-25 21:16:13