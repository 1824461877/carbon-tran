-- MySQL dump 10.13  Distrib 8.0.32, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: trade
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
-- Table structure for table `trade_order`
--

DROP TABLE IF EXISTS `trade_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `trade_order` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `trade_order_id` varchar(255) NOT NULL COMMENT '交易订单id',
                               `pay_order_id` varchar(255) NOT NULL COMMENT '支付订单id',
                               `exchange_asset_id` varchar(255) NOT NULL COMMENT '交易所订单 id',
                               `carbon_asset_id` varchar(255) NOT NULL COMMENT '用户密码',
                               `collection_id` varchar(255) NOT NULL COMMENT '收款账号',
                               `initiator` varchar(255) NOT NULL COMMENT '交易的发起者',
                               `recipient` varchar(255) NOT NULL COMMENT '交易接受者',
                               `trade_status` int NOT NULL COMMENT '交易状态',
                               `number` int NOT NULL COMMENT '数量',
                               `initiator_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               `finish_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `trade_order_id_unique` (`trade_order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trade_order`
--

LOCK TABLES `trade_order` WRITE;
/*!40000 ALTER TABLE `trade_order` DISABLE KEYS */;
INSERT INTO `trade_order` VALUES (37,'349dd3a8-238e-3296-8304-de0607b3a1e7','23b9defa-53a4-37d4-8dc6-7b353cc4b44b','64a5d77e-4525-3f55-a61c-530214ae3b79','faf9092a-5446-315b-9132-4b4947b4afc6','5aca447a-a2a2-3a83-984e-44e30380609c','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','6bad7442-b629-380d-a6ea-ffa16e3b1aa8',1002,1,'2023-09-09 07:05:31','2023-09-09 07:05:31');
/*!40000 ALTER TABLE `trade_order` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-09-08  3:04:32