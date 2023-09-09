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
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets`
--

LOCK TABLES `assets` WRITE;
/*!40000 ALTER TABLE `assets` DISABLE KEYS */;
INSERT INTO `assets` VALUES (1,'8627361a-a3e6-3836-83de-7d1c8af690e2','H2930235125-85974-85974','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','gsf',3,2,'Houji太阳能炊具项目','GS7604','issued','中国','太阳热能','VER',85971,85976,'GS1-1-CN-GS7604-3-2020-21289',2020,0,'2023-09-07 01:34:03'),(23,'e331bcd8-d48f-3736-a494-ca7b5d2f2b3a','H2606604266-85978-85978','3cc4fc18-b22f-3121-a3aa-c110174b1fb0','gsf',1,0,'Houji太阳能炊具项目','GS7604','issued','中国','太阳热能','VER',85978,85978,'GS1-1-CN-GS7604-3-2020-21289',2020,0,'2023-09-07 17:55:59'),(25,'5724336a-7727-3d68-916f-3ad133101424','H3956430617-85977-85977','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','gsf',1,0,'Houji太阳能炊具项目','GS7604','issued','中国','太阳热能','VER',85977,85977,'GS1-1-CN-GS7604-3-2020-21289',2020,0,'2023-09-07 18:04:38');
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
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `retire`
--

LOCK TABLES `retire` WRITE;
/*!40000 ALTER TABLE `retire` DISABLE KEYS */;
INSERT INTO `retire` VALUES (13,'37cfaedf-bd6e-32dc-961f-c0b05a99e70a','0f4255a8-6f5a-3727-9648-e8aa27ec500f','6bad7442-b629-380d-a6ea-ffa16e3b1aa8',1,1001,'07529dbe-6559-3a44-86f1-60a150bd9588','2023-09-07 15:53:47'),(14,'5dc3d0b6-5b6f-33a4-9278-c0c8bf8a18ac','8627361a-a3e6-3836-83de-7d1c8af690e2','6bad7442-b629-380d-a6ea-ffa16e3b1aa8',1,1001,'fb28e3d5-8dfe-36a5-82c5-9cb645aededd','2023-09-07 16:11:51');
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
                               `cid` varchar(255) NOT NULL COMMENT '第三方钱包id',
                               `wallet_type` int NOT NULL COMMENT '钱包类型',
                               `default_collection` tinyint(1) NOT NULL COMMENT '默认收款',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `wallet_id_unique` (`wallet_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_wallet`
--

LOCK TABLES `user_wallet` WRITE;
/*!40000 ALTER TABLE `user_wallet` DISABLE KEYS */;
INSERT INTO `user_wallet` VALUES (9,'5aca447a-a2a2-3a83-984e-44e30380609c','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','C2','sb-rx5di26566864@personal.example.com',8990,1,'2023-09-07 04:27:54','2023-09-07 04:27:54');
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

-- Dump completed on 2023-09-08  3:04:18