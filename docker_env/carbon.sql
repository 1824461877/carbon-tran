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
-- Table structure for table `assets`
--
CREATE DATABASE carbon;
USE carbon;
DROP TABLE IF EXISTS `assets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assets` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `ass_id` varchar(255) NOT NULL COMMENT '用于api操作',
                          `hid` varchar(255) NOT NULL COMMENT '资产唯一标识',
                          `user_id` varchar(255) NOT NULL COMMENT '资产所属用户',
                          `source` varchar(255) NOT NULL COMMENT '来源资产',
                          `number` int NOT NULL COMMENT '数量',
                          `retire_number` int NOT NULL COMMENT '注销数量',
                          `project` varchar(255) NOT NULL COMMENT '项目名称',
                          `gs_id` varchar(20) NOT NULL COMMENT 'gs_id',
                          `status` varchar(20) NOT NULL COMMENT 'status',
                          `country` varchar(20) NOT NULL COMMENT 'country',
                          `project_type` varchar(20) NOT NULL COMMENT 'country',
                          `product` varchar(20) NOT NULL COMMENT 'product',
                          `vers_head` int NOT NULL COMMENT 'vers 序号区间 （head）',
                          `vers_tail` int NOT NULL COMMENT 'vers 序号区间 （tail）',
                          `serial_number` varchar(255) NOT NULL COMMENT '资产编号',
                          `day` int NOT NULL COMMENT '日期时间',
                          `listing` tinyint(1) NOT NULL COMMENT '挂牌',
                          `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '资产创建时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets`
--

LOCK TABLES `assets` WRITE;
/*!40000 ALTER TABLE `assets` DISABLE KEYS */;
INSERT INTO `assets` VALUES (1,'ab082c70-644a-3711-ac28-d4755b79d26a','H102771824-85975-85975','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','gsf',3,2,'Houji太阳\n能炊具项目','GS7604','issued','中国','太阳热能','VER',85975,85979,'GS1-1-CN-GS7604-3-2020-21289',2020,0,'2023-06-10 15:53:13');
/*!40000 ALTER TABLE `assets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `assets_sell`
--

DROP TABLE IF EXISTS `assets_sell`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assets_sell` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `ex_id` varchar(255) NOT NULL COMMENT '交易单号',
                               `ass_id` varchar(255) NOT NULL COMMENT '资产编号',
                               `collection_wallet_id` varchar(255) NOT NULL COMMENT '收款的钱包 id',
                               `user_id` varchar(255) NOT NULL COMMENT '资产所属用户',
                               `source` varchar(255) NOT NULL COMMENT '来源资产',
                               `amount` double NOT NULL COMMENT 'number(数量)/amount(金额)',
                               `number` int NOT NULL COMMENT '数量',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易创建时间',
                               `end_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易结束时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets_sell`
--

LOCK TABLES `assets_sell` WRITE;
/*!40000 ALTER TABLE `assets_sell` DISABLE KEYS */;
/*!40000 ALTER TABLE `assets_sell` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `retire`
--

DROP TABLE IF EXISTS `retire`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `retire` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `r_id` varchar(255) NOT NULL COMMENT '注销单号ID',
                          `ass_id` varchar(255) NOT NULL COMMENT '注销的资产id',
                          `user_id` varchar(255) NOT NULL COMMENT '资产注销用户id',
                          `number` int NOT NULL COMMENT '数量',
                          `status` int NOT NULL COMMENT '注销状态',
                          `certificate_link` varchar(255) NOT NULL COMMENT '证书访问link',
                          `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注销时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `retire`
--

LOCK TABLES `retire` WRITE;
/*!40000 ALTER TABLE `retire` DISABLE KEYS */;
INSERT INTO `retire` VALUES (1,'f3af4ed4-69d5-3cef-bea2-f19c9d8c4300','ab082c70-644a-3711-ac28-d4755b79d26a','6bad7442-b629-380d-a6ea-ffa16e3b1aa8',2,1001,'xxx','2023-08-15 10:15:20');
/*!40000 ALTER TABLE `retire` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_wallet`
--

DROP TABLE IF EXISTS `user_wallet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_wallet` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `wallet_id` varchar(255) NOT NULL COMMENT '钱包ID',
                               `user_id` varchar(255) NOT NULL COMMENT '钱包所属用户ID',
                               `name` varchar(255) NOT NULL COMMENT '钱包名',
                               `password` varchar(255) NOT NULL COMMENT '用户密码',
                               `amount` double NOT NULL COMMENT '钱包余额',
                               `salt` varchar(255) NOT NULL COMMENT '密码salt',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `wallet_id_unique` (`wallet_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_wallet`
--

LOCK TABLES `user_wallet` WRITE;
/*!40000 ALTER TABLE `user_wallet` DISABLE KEYS */;
INSERT INTO `user_wallet` VALUES (1,'d87e60be-5e75-3d17-81c6-58da08f1e7c9','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','钱包1','dc9ebbd924a85429dbffab825ab6bd73c3a75d444147264a6199dcfe7034afbd',490.1,'7c36344a-a198-3329-9b99-a9814f5d70401686411923445321500','2023-06-10 11:56:05','2023-06-10 16:22:34'),(2,'f25dacb8-1d10-3bd0-88c6-dcdcf7bb90df','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','钱包2','e05564178142b35a0f7cac5e6bec637cc906ecce2703a327fc23aa9903a568e8',3820.3,'21781fe3-b040-3d74-97f2-140726cc222f1686416419192786800','2023-06-10 13:17:15','2023-06-10 16:22:34'),(4,'836580f0-6d7f-3773-b835-f4adf57bf2dc','3cc4fc18-b22f-3121-a3aa-c110174b1fb0','Jos钱包1','e1f4f1620b1006da9e7a186778f5839792d9e4954e7608727ae93a73ec24a379',4800.9,'c0b93017-ce13-331b-9011-4ce21099a19b1686439205085937600','2023-06-10 16:08:11','2023-06-10 16:22:34');
/*!40000 ALTER TABLE `user_wallet` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-15 20:22:52
