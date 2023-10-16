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
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets`
--

LOCK TABLES `assets` WRITE;
/*!40000 ALTER TABLE `assets` DISABLE KEYS */;
INSERT INTO `assets` VALUES (100,'3a633829-a4e6-33ba-9cbc-8b56cca6c1f5','H2238975697-1-1001','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','cvs',1000,0,'新疆玛基特县防风固沙生态林建设基地项目','CVS002','issued','cn','植树-造林','CVS',1,1001,'CVS-CN-XJ-1-1000-CVS002-2020-2022',2020,0,'2023-10-08 02:34:48');
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
                               `country` varchar(5) NOT NULL COMMENT '资产所属国家',
                               `collection_wallet_id` varchar(255) NOT NULL COMMENT '收款的钱包 id',
                               `user_id` varchar(255) NOT NULL COMMENT '资产所属用户',
                               `source` varchar(255) NOT NULL COMMENT '来源资产',
                               `amount` double NOT NULL COMMENT 'number(数量)/amount(金额)',
                               `number` int NOT NULL COMMENT '数量',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易创建时间',
                               `end_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易结束时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets_sell`
--

LOCK TABLES `assets_sell` WRITE;
/*!40000 ALTER TABLE `assets_sell` DISABLE KEYS */;
INSERT INTO `assets_sell` VALUES (23,'b99d624b-9971-3deb-b5f1-e2e685023361','17fc4e0b-71fb-309d-b0b4-e01ef098ccf5','jp','11c8e330-c951-3bd1-a53b-0544c3b06515','70d68d1a-1c0f-3911-a8a7-917948225526','jcm',10,3,'2023-09-26 19:54:41','2023-09-28 10:13:30'),(24,'62e1414a-0acd-38fa-9357-1ecfd956826b','52fd4d5d-cb2a-38b5-a096-1e7892db1494','jp','11c8e330-c951-3bd1-a53b-0544c3b06515','70d68d1a-1c0f-3911-a8a7-917948225526','jcm',10,3,'2023-09-26 19:54:45','2023-09-28 10:13:33'),(25,'e94971d4-5496-3d46-a5c9-cf631e0596e5','21bf0af1-2bc6-3d1c-9720-70e540b949af','th','11c8e330-c951-3bd1-a53b-0544c3b06515','70d68d1a-1c0f-3911-a8a7-917948225526','jcm',10,3,'2023-09-26 19:54:47','2023-09-28 10:13:35'),(26,'27e8f406-46c7-3260-b10b-a55efcb2e4b2','c4297215-de43-3fa2-b250-b83a09c9c864','jp','11c8e330-c951-3bd1-a53b-0544c3b06515','70d68d1a-1c0f-3911-a8a7-917948225526','jcm',10,3,'2023-09-26 19:54:48','2023-09-28 10:13:36'),(27,'28c8615d-17a0-3a12-ab26-48801ed376e8','75ecf669-98b4-37ec-8165-ac0ab746c5fb','jp','11c8e330-c951-3bd1-a53b-0544c3b06515','70d68d1a-1c0f-3911-a8a7-917948225526','jcm',10,3,'2023-09-26 19:54:50','2023-09-28 10:13:38');
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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `retire`
--

LOCK TABLES `retire` WRITE;
/*!40000 ALTER TABLE `retire` DISABLE KEYS */;
INSERT INTO `retire` VALUES (15,'ff66ac4f-036c-3618-91d5-eeb35c701238','2165fd97-a90c-3f1b-9792-beca5b8ad30b','6bad7442-b629-380d-a6ea-ffa16e3b1aa8',1,1001,'67ac279a-2684-3c9c-b8a2-ad76c9c21ca2','2023-09-24 10:48:53'),(16,'27445bee-c45f-355e-9434-0ea5aeda24f4','c1227759-d470-340e-8c0b-d23d32bfe155','6bad7442-b629-380d-a6ea-ffa16e3b1aa8',1,1001,'63e53632-7eb0-3c1e-999a-e84ac04c8ca0','2023-09-24 12:00:36');
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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_wallet`
--

LOCK TABLES `user_wallet` WRITE;
/*!40000 ALTER TABLE `user_wallet` DISABLE KEYS */;
INSERT INTO `user_wallet` VALUES (12,'64e55f35-16e3-3b4c-a374-dcb585d1af89','6bad7442-b629-380d-a6ea-ffa16e3b1aa8','Coms','sb-cse47215937073@personal.example.com',8990,1,'2023-09-23 09:05:40','2023-09-23 09:05:40'),(13,'11c8e330-c951-3bd1-a53b-0544c3b06515','70d68d1a-1c0f-3911-a8a7-917948225526','COMS','1824461877@qq.com',8990,1,'2023-09-26 02:54:57','2023-09-26 02:54:57');
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

-- Dump completed on 2023-10-10  1:02:09