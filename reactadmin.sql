-- MySQL dump 10.13  Distrib 8.0.23, for Linux (x86_64)
--
-- Host: localhost    Database: reactadmin
-- ------------------------------------------------------
-- Server version	8.0.23-0ubuntu0.20.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(45) NOT NULL,
  `password` varchar(200) NOT NULL DEFAULT '',
  `role_id` int NOT NULL DEFAULT '0',
  `descript` varchar(100) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_enabled` smallint NOT NULL DEFAULT '1',
  `last_login_ip` varchar(15) NOT NULL DEFAULT '',
  `last_login_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`),
  KEY `123` (`username`,`password`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','$argon2id$v=19$m=65536,t=4,p=1$SQIMdrDytXhkOrFprwOFlA$AVnEMpK8b076NZylm/G0P2Krcls9VkBBoWMHBkPrTAg',1,'123','2021-03-04 16:42:41',1,'','2021-03-04 16:42:41'),(37,'mash','$argon2id$v=19$m=65536,t=4,p=1$iYR1yQgkvI2lXerRqxmIkg$AX7DWgYMFizIoEC3jPdw1hwwVl3ruxApz7iB7f1iArk',5,'1234456','2021-02-18 14:26:00',1,'','2021-02-18 14:26:00');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_action`
--

DROP TABLE IF EXISTS `admin_action`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_action` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `menu_id` int DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `router_name` varchar(100) DEFAULT NULL,
  `router_short_name` varchar(100) DEFAULT NULL,
  `is_sub_menu` tinyint DEFAULT '0',
  `sorted_by` smallint DEFAULT '100',
  `icon` varchar(100) DEFAULT 'fi fi-record',
  `controller_action` varchar(150) DEFAULT NULL,
  `is_enabled` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_admin_action_router_name` (`router_name`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_action`
--

LOCK TABLES `admin_action` WRITE;
/*!40000 ALTER TABLE `admin_action` DISABLE KEYS */;
INSERT INTO `admin_action` VALUES (1,0,'角色新增','/admin/role/add','/admin/role',0,100,'far fa-file','',1),(2,0,'角色编辑','/admin/role/edit/:id','/admin/role',0,100,'far fa-file','',1),(3,1,'角色列表','/admin/role/index','/admin/role',1,100,'far fa-file','',1),(4,0,'管理员新增','/admin/admin/add','/admin/admin',0,100,'far fa-file','',1),(5,0,'管理员编辑','/admin/admin/edit/:id','/admin/admin',0,100,'far fa-file','',1),(6,1,'管理员列表','/admin/admin/index','/admin/admin',1,100,'far fa-file','',1),(7,0,'菜单编辑','/admin/menu/edit/:id','/admin/menu',0,100,'far fa-file','',1),(8,1,'菜单列表','/admin/menu/index','/admin/menu',1,100,'far fa-file','',1),(9,0,'查看功能','/admin/menu/edit/action/:id','/admin/menu',0,100,'fi fi-record','',1),(10,0,'菜单功能编辑','/admin/action/edit/:id','/admin/action',0,100,'fi fi-record','',1),(11,0,'设置api','/admin/action/api/:id','/admin/action',0,100,'fi fi-record','',1),(12,1,'菜单功能列表','/admin/action/index','/admin/action',1,100,'fi fi-record','',1),(13,0,'角色授权','/admin/role/auth/:id','/admin/role',0,100,'fi fi-record','',1);
/*!40000 ALTER TABLE `admin_action` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_action_api`
--

DROP TABLE IF EXISTS `admin_action_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_action_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `action_id` int DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `method` varchar(10) DEFAULT NULL,
  `path` varchar(100) DEFAULT NULL,
  `sorted_by` smallint DEFAULT '100',
  `controller_action` varchar(150) DEFAULT NULL,
  `is_enabled` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_action_api`
--

LOCK TABLES `admin_action_api` WRITE;
/*!40000 ALTER TABLE `admin_action_api` DISABLE KEYS */;
INSERT INTO `admin_action_api` VALUES (1,12,'','POST','/api/admin/action/add',100,'api/controller/adminactioncontroller.Add',1),(2,12,'','POST','/api/admin/action/edit/:id',100,'api/controller/adminactioncontroller.EditSave',1),(3,12,'','POST','/api/admin/action/status',100,'api/controller/adminactioncontroller.Status',1),(4,12,'','POST','/api/admin/action/delete',100,'api/controller/adminactioncontroller.Delete',1),(5,12,'','POST','/api/admin/action/menu',100,'api/controller/adminactioncontroller.SetActionMenu',1),(6,12,'','POST','/api/admin/action/collect',100,'api/controller/adminactioncontroller.CollectAction',1),(7,6,'','POST','/api/admin/add',100,'api/controller/admincontroller.Add',1),(8,3,'','POST','/api/admin/role/add',100,'api/controller/adminrolecontroller.Add',1),(9,3,'','POST','/api/admin/role/edit/:id',100,'api/controller/adminrolecontroller.EditSave',1),(10,3,'','POST','/api/admin/role/status',100,'api/controller/adminrolecontroller.Status',1),(11,3,'','POST','/api/admin/role/delete',100,'api/controller/adminrolecontroller.Delete',1),(12,8,'','POST','/api/admin/menu/add',100,'api/controller/adminmenucontroller.Add',1),(13,8,'','POST','/api/admin/menu/edit/:id',100,'api/controller/adminmenucontroller.EditSave',1),(14,8,'','POST','/api/admin/menu/status',100,'api/controller/adminmenucontroller.Status',1),(15,8,'','POST','/api/admin/menu/delete',100,'api/controller/adminmenucontroller.Delete',1),(16,0,'','POST','/api/admin/login',100,'api/controller/securitycontroller.Login',1),(17,6,'','POST','/api/admin/edit/:id',100,'api/controller/admincontroller.EditSave',1),(18,6,'','POST','/api/admin/delete',100,'api/controller/admincontroller.Delete',1),(19,6,'','POST','/api/admin/status',100,'api/controller/admincontroller.Status',1),(20,6,'','POST','/api/admin/password',100,'api/controller/admincontroller.Password',1),(21,8,'','GET','/api/admin/menu/all',100,'api/controller/adminmenucontroller.All',1),(22,8,'','GET','/api/admin/menu/action/:id',100,'api/controller/adminmenucontroller.MenuAction',1),(23,8,'','GET','/api/admin/menu/index',100,'api/controller/adminmenucontroller.Index',1),(24,8,'','GET','/api/admin/menu/edit/:id',100,'api/controller/adminmenucontroller.Edit',1),(25,3,'','GET','/api/admin/role/all',100,'api/controller/adminrolecontroller.All',1),(26,3,'','GET','/api/admin/role/index',100,'api/controller/adminrolecontroller.Index',1),(27,3,'','GET','/api/admin/role/edit/:id',100,'api/controller/adminrolecontroller.Edit',1),(28,12,'','GET','/api/admin/action/index',100,'api/controller/adminactioncontroller.Index',1),(29,12,'','GET','/api/admin/action/edit/:id',100,'api/controller/adminactioncontroller.Edit',1),(30,6,'','GET','/api/admin/index',100,'api/controller/admincontroller.Index',1),(31,6,'','GET','/api/admin/edit/:id',100,'api/controller/admincontroller.Edit',1),(32,12,'','POST','/api/admin/action/api/set',100,'api/controller/adminactionapicontroller.ApiSetAction',1),(33,3,'','POST','/api/admin/role/rbac/set/:id',100,'api/controller/adminrolerbaccontroller.RbacSet',1),(34,3,'','GET','/api/admin/role/rbac/:id',100,'api/controller/adminrolerbaccontroller.RbacInfo',1),(35,12,'','GET','/api/admin/action/api',100,'api/controller/adminactionapicontroller.Api',1);
/*!40000 ALTER TABLE `admin_action_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_menu`
--

DROP TABLE IF EXISTS `admin_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `sign` varchar(20) DEFAULT NULL COMMENT '唯一标识',
  `sorted_by` smallint DEFAULT '100',
  `icon` varchar(30) DEFAULT 'fab fa-symfony',
  `is_enabled` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_admin_menu_sign` (`sign`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_menu`
--

LOCK TABLES `admin_menu` WRITE;
/*!40000 ALTER TABLE `admin_menu` DISABLE KEYS */;
INSERT INTO `admin_menu` VALUES (1,'权限管理','auth',100,'fab fa-symfony',1);
/*!40000 ALTER TABLE `admin_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_role`
--

DROP TABLE IF EXISTS `admin_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `is_enabled` smallint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role`
--

LOCK TABLES `admin_role` WRITE;
/*!40000 ALTER TABLE `admin_role` DISABLE KEYS */;
INSERT INTO `admin_role` VALUES (1,'超级管理员',1),(5,'测试角色',1),(12,'3ss121',0),(13,'dd',1),(16,'33',1),(17,'44',1),(18,'123',1);
/*!40000 ALTER TABLE `admin_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_role_action`
--

DROP TABLE IF EXISTS `admin_role_action`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role_action` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int DEFAULT NULL,
  `action_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role_action`
--

LOCK TABLES `admin_role_action` WRITE;
/*!40000 ALTER TABLE `admin_role_action` DISABLE KEYS */;
INSERT INTO `admin_role_action` VALUES (12,18,12),(13,18,8),(14,1,3),(15,1,6),(16,1,8),(17,1,12);
/*!40000 ALTER TABLE `admin_role_action` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_role_action_api`
--

DROP TABLE IF EXISTS `admin_role_action_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role_action_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int DEFAULT NULL,
  `action_id` int DEFAULT NULL,
  `api_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role_action_api`
--

LOCK TABLES `admin_role_action_api` WRITE;
/*!40000 ALTER TABLE `admin_role_action_api` DISABLE KEYS */;
INSERT INTO `admin_role_action_api` VALUES (101,18,12,1),(102,18,12,2),(103,18,12,3),(104,18,12,4),(105,18,12,5),(106,18,12,6),(107,18,12,28),(108,18,12,29),(109,18,12,32),(110,18,12,35),(111,18,8,12),(112,18,8,13),(113,18,8,14),(114,18,8,15),(115,18,8,21),(116,18,8,22),(117,18,8,23),(118,18,8,24),(119,1,3,8),(120,1,3,9),(121,1,3,10),(122,1,3,11),(123,1,3,25),(124,1,3,26),(125,1,3,27),(126,1,3,33),(127,1,3,34),(128,1,6,7),(129,1,6,17),(130,1,6,18),(131,1,6,19),(132,1,6,20),(133,1,6,30),(134,1,6,31),(135,1,8,12),(136,1,8,13),(137,1,8,14),(138,1,8,15),(139,1,8,21),(140,1,8,22),(141,1,8,23),(142,1,8,24),(143,1,12,1),(144,1,12,2),(145,1,12,3),(146,1,12,4),(147,1,12,5),(148,1,12,6),(149,1,12,28),(150,1,12,29),(151,1,12,32),(152,1,12,35);
/*!40000 ALTER TABLE `admin_role_action_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_role_menu`
--

DROP TABLE IF EXISTS `admin_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int DEFAULT NULL,
  `menu_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role_menu`
--

LOCK TABLES `admin_role_menu` WRITE;
/*!40000 ALTER TABLE `admin_role_menu` DISABLE KEYS */;
INSERT INTO `admin_role_menu` VALUES (8,18,1),(9,1,1);
/*!40000 ALTER TABLE `admin_role_menu` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-11 17:18:10
