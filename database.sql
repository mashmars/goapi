CREATE DATABASE  IF NOT EXISTS `sfadmin` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `sfadmin`;
-- MySQL dump 10.13  Distrib 8.0.23, for Linux (x86_64)
--
-- Host: localhost    Database: sfadmin
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
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','$argon2id$v=19$m=65536,t=4,p=1$QcESAIxW4hETavFWAFTu3g$weGomM1I5cAa6GVBVvNZ+CeGn9HVNzvbFni4HWSXDhM',1,'desc','2021-02-02 00:00:00',1,'127.0.0.1','2021-02-02 00:00:00'),(37,'mash','$argon2id$v=19$m=65536,t=4,p=1$SQIMdrDytXhkOrFprwOFlA$AVnEMpK8b076NZylm/G0P2Krcls9VkBBoWMHBkPrTAg',5,'1234456','2021-02-18 14:26:00',1,'','2021-02-18 14:26:00');
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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_action`
--

LOCK TABLES `admin_action` WRITE;
/*!40000 ALTER TABLE `admin_action` DISABLE KEYS */;
INSERT INTO `admin_action` VALUES (1,1,'菜单功能列表','admin.admin_action.index','admin.admin_action',1,100,'fi fi-record','App\\Controller\\Admin\\AdminActionController::index',1),(2,0,'菜单功能新增','admin.admin_action.add','admin.admin_action',0,100,'fi fi-record','App\\Controller\\Admin\\AdminActionController::add',1),(3,0,'菜单功能编辑','admin.admin_action.edit','admin.admin_action',0,100,'fi fi-record','App\\Controller\\Admin\\AdminActionController::edit',1),(4,0,'菜单功能更新','admin.admin_action.check','admin.admin_action',0,100,'fi fi-record','App\\Controller\\Admin\\AdminActionController::check',1),(5,1,'管理员列表','admin.admin.index','admin.admin',1,100,'fi fi-record','App\\Controller\\Admin\\AdminController::index',1),(6,0,'管理员新增','admin.admin.add','admin.admin',0,100,'fi fi-record','App\\Controller\\Admin\\AdminController::add',1),(7,0,'管理员编辑','admin.admin.edit','admin.admin',0,100,'fi fi-record','App\\Controller\\Admin\\AdminController::edit',1),(8,0,'管理员修改密码','admin.admin.password','admin.admin',0,100,'fi fi-record','App\\Controller\\Admin\\AdminController::password',1),(9,1,'菜单列表','admin.admin_menu.index','admin.admin_menu',1,100,'fi fi-record','App\\Controller\\Admin\\AdminMenuController::index',1),(10,0,'菜单新增','admin.admin_menu.add','admin.admin_menu',0,100,'fi fi-record','App\\Controller\\Admin\\AdminMenuController::add',1),(11,0,'菜单编辑','admin.admin_menu.edit','admin.admin_menu',0,100,'fi fi-record','App\\Controller\\Admin\\AdminMenuController::edit',1),(12,1,'管理员角色列表','admin.admin_role.index','admin.admin_role',1,100,'fi fi-record','App\\Controller\\Admin\\AdminRoleController::index',1),(13,0,'管理员角色新增','admin.admin_role.add','admin.admin_role',0,100,'fi fi-record','App\\Controller\\Admin\\AdminRoleController::add',1),(14,0,'管理员角色编辑','admin.admin_role.edit','admin.admin_role',0,100,'fi fi-record','App\\Controller\\Admin\\AdminRoleController::edit',1),(15,0,'功能api设置','admin.admin_action.api','admin.admin_action',0,100,'fi fi-record','App\\Controller\\Admin\\AdminActionController::api',1),(16,0,'子菜单功能','admin.admin.admin_menu.showAction','admin.admin.admin_menu',0,100,'fi fi-record','App\\Controller\\Admin\\AdminMenuController::show',1);
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
INSERT INTO `admin_action_api` VALUES (1,1,'','POST','/api/admin/action/add',100,'api/controller/adminactioncontroller.Add',1),(2,1,'','POST','/api/admin/action/edit/:id',100,'api/controller/adminactioncontroller.EditSave',1),(3,1,'','POST','/api/admin/action/status',100,'api/controller/adminactioncontroller.Status',1),(4,1,'','POST','/api/admin/action/delete',100,'api/controller/adminactioncontroller.Delete',1),(5,1,'','POST','/api/admin/action/menu',100,'api/controller/adminactioncontroller.SetActionMenu',1),(6,1,'','POST','/api/admin/action/collect',100,'api/controller/adminactioncontroller.CollectAction',1),(7,5,'','POST','/api/admin/add',100,'api/controller/admincontroller.Add',1),(8,12,'','POST','/api/admin/role/add',100,'api/controller/adminrolecontroller.Add',1),(9,12,'','POST','/api/admin/role/edit/:id',100,'api/controller/adminrolecontroller.EditSave',1),(10,12,'','POST','/api/admin/role/status',100,'api/controller/adminrolecontroller.Status',1),(11,12,'','POST','/api/admin/role/delete',100,'api/controller/adminrolecontroller.Delete',1),(12,9,'','POST','/api/admin/menu/add',100,'api/controller/adminmenucontroller.Add',1),(13,9,'','POST','/api/admin/menu/edit/:id',100,'api/controller/adminmenucontroller.EditSave',1),(14,9,'','POST','/api/admin/menu/status',100,'api/controller/adminmenucontroller.Status',1),(15,9,'','POST','/api/admin/menu/delete',100,'api/controller/adminmenucontroller.Delete',1),(16,0,'','POST','/api/admin/login',100,'api/controller/securitycontroller.Login',1),(17,5,'','POST','/api/admin/edit/:id',100,'api/controller/admincontroller.EditSave',1),(18,5,'','POST','/api/admin/delete',100,'api/controller/admincontroller.Delete',1),(19,5,'','POST','/api/admin/status',100,'api/controller/admincontroller.Status',1),(20,5,'','POST','/api/admin/password',100,'api/controller/admincontroller.Password',1),(21,9,'','GET','/api/admin/menu/all',100,'api/controller/adminmenucontroller.All',1),(22,9,'','GET','/api/admin/menu/action/:id',100,'api/controller/adminmenucontroller.MenuAction',1),(23,9,'','GET','/api/admin/menu/index',100,'api/controller/adminmenucontroller.Index',1),(24,9,'','GET','/api/admin/menu/edit/:id',100,'api/controller/adminmenucontroller.Edit',1),(25,12,'','GET','/api/admin/role/all',100,'api/controller/adminrolecontroller.All',1),(26,12,'','GET','/api/admin/role/index',100,'api/controller/adminrolecontroller.Index',1),(27,12,'','GET','/api/admin/role/edit/:id',100,'api/controller/adminrolecontroller.Edit',1),(28,1,'','GET','/api/admin/action/index',100,'api/controller/adminactioncontroller.Index',1),(29,1,'','GET','/api/admin/action/edit/:id',100,'api/controller/adminactioncontroller.Edit',1),(30,5,'','GET','/api/admin/index',100,'api/controller/admincontroller.Index',1),(31,5,'','GET','/api/admin/edit/:id',100,'api/controller/admincontroller.Edit',1),(32,1,'','POST','/api/admin/action/api/set',100,'api/controller/adminactionapicontroller.ApiSetAction',1),(33,12,'','POST','/api/admin/role/rbac/set/:id',100,'api/controller/adminrolerbaccontroller.RbacSet',1),(34,12,'','GET','/api/admin/role/rbac/:id',100,'api/controller/adminrolerbaccontroller.RbacInfo',1),(35,1,'','GET','/api/admin/action/api',100,'api/controller/adminactionapicontroller.Api',1);
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
INSERT INTO `admin_menu` VALUES (1,'权限管理','auth',100,'fas fa-users-cog',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role`
--

LOCK TABLES `admin_role` WRITE;
/*!40000 ALTER TABLE `admin_role` DISABLE KEYS */;
INSERT INTO `admin_role` VALUES (1,'超级管理员',1),(5,'测试角色',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role_action`
--

LOCK TABLES `admin_role_action` WRITE;
/*!40000 ALTER TABLE `admin_role_action` DISABLE KEYS */;
INSERT INTO `admin_role_action` VALUES (70,1,1),(71,1,5),(72,1,9),(73,1,12),(98,5,1),(99,5,5),(100,5,9),(101,5,12);
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
) ENGINE=InnoDB AUTO_INCREMENT=620 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role_action_api`
--

LOCK TABLES `admin_role_action_api` WRITE;
/*!40000 ALTER TABLE `admin_role_action_api` DISABLE KEYS */;
INSERT INTO `admin_role_action_api` VALUES (482,1,1,1),(483,1,1,2),(484,1,1,3),(485,1,1,4),(486,1,1,5),(487,1,1,6),(488,1,1,28),(489,1,1,29),(490,1,1,32),(491,1,1,35),(492,1,5,7),(493,1,5,17),(494,1,5,18),(495,1,5,19),(496,1,5,20),(497,1,5,30),(498,1,5,31),(499,1,9,12),(500,1,9,13),(501,1,9,14),(502,1,9,15),(503,1,9,21),(504,1,9,22),(505,1,9,23),(506,1,9,24),(507,1,12,8),(508,1,12,9),(509,1,12,10),(510,1,12,11),(511,1,12,25),(512,1,12,26),(513,1,12,27),(514,1,12,33),(515,1,12,34),(607,5,1,28),(608,5,5,7),(609,5,5,30),(610,5,5,31),(611,5,9,21),(612,5,9,23),(613,5,12,8),(614,5,12,10),(615,5,12,25),(616,5,12,26),(617,5,12,27),(618,5,12,33),(619,5,12,34);
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
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role_menu`
--

LOCK TABLES `admin_role_menu` WRITE;
/*!40000 ALTER TABLE `admin_role_menu` DISABLE KEYS */;
INSERT INTO `admin_role_menu` VALUES (20,1,1),(29,5,1);
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

-- Dump completed on 2021-02-26 15:18:38
