-- MySQL dump 10.13  Distrib 8.0.32, for Linux (x86_64)
--
-- Host: localhost    Database: nocake
-- ------------------------------------------------------
-- Server version	8.0.32

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
-- Table structure for table `t_address`
--

DROP TABLE IF EXISTS `t_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_address` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(63) NOT NULL DEFAULT '' COMMENT '收货人名称',
  `open_id` varchar(63) NOT NULL DEFAULT '' COMMENT '用户表的用户',
  `province` varchar(63) NOT NULL DEFAULT '' COMMENT '行政区域表的省',
  `city` varchar(63) NOT NULL DEFAULT '' COMMENT '行政区域表的市',
  `county` varchar(63) NOT NULL DEFAULT '' COMMENT '行政区域表的区县',
  `detail` varchar(127) NOT NULL DEFAULT '' COMMENT '详细收货地址',
  `tel` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号码',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认地址',
  `deleted` tinyint(1) NOT NULL DEFAULT '1' COMMENT '逻辑删除 1默认 2删除',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_open_id` (`open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='微信商城小程序-收货地址表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_address`
--

LOCK TABLES `t_address` WRITE;
/*!40000 ALTER TABLE `t_address` DISABLE KEYS */;
INSERT INTO `t_address` VALUES (1,'张印','lainzhang','成都','','','','',1,1,'2023-03-13 09:50:57','2023-03-13 09:50:57',NULL),(2,'张印','owvNO46BJBnAfD977z7D-Vu3k7gU','北京市','北京市','东城区','2704','131',1,0,'2023-03-15 06:28:56','2023-03-15 13:07:47','0000-00-00 00:00:00'),(3,'1','owvNO46BJBnAfD977z7D-Vu3k7gU','北京市','北京市','东城区','31212','1',0,0,'2023-03-15 13:00:37','2023-03-15 13:07:46','0000-00-00 00:00:00');
/*!40000 ALTER TABLE `t_address` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_category`
--

DROP TABLE IF EXISTS `t_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(63) NOT NULL DEFAULT '' COMMENT '类目名称',
  `brief` varchar(255) NOT NULL DEFAULT '' COMMENT '类目简介',
  `keywords` varchar(1023) NOT NULL DEFAULT '' COMMENT '类目关键字，以JSON数组格式',
  `pid` int NOT NULL DEFAULT '0' COMMENT '父类目ID',
  `level` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类目层级',
  `icon_url` varchar(255) NOT NULL DEFAULT '' COMMENT '类目图标',
  `weight` tinyint NOT NULL DEFAULT '50' COMMENT '排序',
  `deleted` tinyint(1) NOT NULL DEFAULT '1' COMMENT '逻辑删除 1默认 2删除',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `uk_pid_level` (`pid`,`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='微信商城小程序-类目表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_category`
--

LOCK TABLES `t_category` WRITE;
/*!40000 ALTER TABLE `t_category` DISABLE KEYS */;
INSERT INTO `t_category` VALUES (1,'蛋糕','蛋糕，大家都喜欢','{\"受众\":\"大众\"}',0,1,'http://tubiao.png',50,1,'2023-03-13 09:50:56','2023-03-13 09:50:56',NULL),(2,'甜品','甜品，少数人喜欢','{\"受众\":\"小众\"}',0,1,'http://tubiao.png',50,1,'2023-03-13 09:50:56','2023-03-13 09:50:56',NULL),(3,'甜品1','甜品，少数人喜欢','{\"受众\":\"小众\"}',0,1,'http://tubiao.png',50,1,'2023-03-13 09:50:56','2023-03-13 09:50:56',NULL),(4,'蛋糕2','蛋糕，大家都喜欢','{\"受众\":\"大众\"}',0,1,'http://tubiao.png',50,1,'2023-03-13 09:50:56','2023-03-13 09:50:56',NULL),(5,'甜品3','甜品，少数人喜欢','{\"受众\":\"小众\"}',0,1,'http://tubiao.png',50,1,'2023-03-13 09:50:56','2023-03-13 09:50:56',NULL),(6,'甜品4','甜品，少数人喜欢','{\"受众\":\"小众\"}',0,1,'http://tubiao.png',50,1,'2023-03-13 09:50:56','2023-03-13 09:50:56',NULL);
/*!40000 ALTER TABLE `t_category` ENABLE KEYS */;
UNLOCK TABLES;

DROP TABLE IF EXISTS `t_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_goods` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(127) NOT NULL DEFAULT '' COMMENT '商品名称',
  `brief` varchar(1023) NOT NULL DEFAULT '' COMMENT '商品简介',
  `category_id` int NOT NULL DEFAULT '0' COMMENT '商品所属类目ID',
  `keywords` varchar(255) NOT NULL DEFAULT '' COMMENT '商品关键字，采用逗号间隔',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1未上架 2已上架',
  `is_hot` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1不推荐 2推荐',
  `weight` tinyint(1) NOT NULL DEFAULT '100',
  `pic_url` varchar(255) NOT NULL DEFAULT '' COMMENT '商品主图',
  `pic_urls` varchar(1023) NOT NULL DEFAULT '' COMMENT '商品附图JSON数组格式',
  `unit` varchar(31) NOT NULL DEFAULT '个' COMMENT '商品单位，例如件、盒',
  `quantity` int NOT NULL DEFAULT '1' COMMENT '商品库存',
  `price` decimal(10,2) NOT NULL DEFAULT '100000.00' COMMENT '价格',
  `options` varchar(1023) NOT NULL DEFAULT '{}' COMMENT '商品购买选项',
  `deleted` tinyint(1) NOT NULL DEFAULT '1' COMMENT '逻辑删除 1默认 2删除',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_weight` (`weight`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='微信商城小程序-商品基本信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_goods`
--

LOCK TABLES `t_goods` WRITE;
/*!40000 ALTER TABLE `t_goods` DISABLE KEYS */;
INSERT INTO `t_goods` VALUES (1,'水果蛋糕1','简介11111',1,'大众,水果,蛋糕',2,2,100,'https://www.nocake.cn/images/QQ浏览器截图20221202200134.png','[\"http://tubiao.png\",\"http://tubiao1.png\"]','个',1,128.00, '{"尺寸":["8寸","6寸"],"口味":["奥利奥","原味"]}',1,'2023-03-13 09:50:56','2023-03-15 13:58:17',NULL);
/*!40000 ALTER TABLE `t_goods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_order`
--

DROP TABLE IF EXISTS `t_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `open_id` varchar(63) NOT NULL DEFAULT '' COMMENT '用户表的用户ID',
  `options` varchar(1023) NOT NULL DEFAULT '' COMMENT '商品选项json',
  `status` smallint NOT NULL DEFAULT '1' COMMENT '订单状态 待付款1 已取消2 已付款3 配送中4 已完成5',
  `address` varchar(1023) NOT NULL DEFAULT '' COMMENT '收货具体地址',
  `message` varchar(512) NOT NULL DEFAULT '' COMMENT '用户订单留言',
  `goods_price` decimal(10,2) NOT NULL DEFAULT '100000.00' COMMENT '商品总费用',
  `goods_count` tinyint NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `coupon_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠券减免',
  `dis_price` decimal(10,2) NOT NULL DEFAULT '100000.00' COMMENT '配送费用',
  `integral_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '用户积分减免',
  `groupon_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '团购优惠价减免',
  `order_price` decimal(10,2) NOT NULL DEFAULT '100000.00' COMMENT '订单费用， = goods_price + dis_price - coupon_price',
  `actual_price` decimal(10,2) NOT NULL DEFAULT '100000.00' COMMENT '实付费用， = order_price - integral_price',
  `pay_id` varchar(63) NOT NULL DEFAULT '' COMMENT '微信付款编号',
  `pay_time` timestamp NULL DEFAULT NULL COMMENT '微信付款时间',
  `ship_sn` varchar(63) NOT NULL DEFAULT '' COMMENT '外卖订单',
  `ship_channel` tinyint NOT NULL DEFAULT '0' COMMENT '外卖平台',
  `ship_time` timestamp NULL DEFAULT NULL COMMENT '发货开始时间',
  `confirm_time` timestamp NULL DEFAULT NULL COMMENT '用户确认收货时间',
  `end_time` timestamp NULL DEFAULT NULL COMMENT '订单关闭时间',
  `deleted` tinyint(1) NOT NULL DEFAULT '1' COMMENT '逻辑删除 1默认 2删除',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_open_id` (`open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='微信商城小程序-订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_order`
--

LOCK TABLES `t_order` WRITE;
/*!40000 ALTER TABLE `t_order` DISABLE KEYS */;
INSERT INTO `t_order` VALUES (1,'owvNO46BJBnAfD977z7D-Vu3k7gU','{\"2\":\"5\"}',0,0,'','1231231',640.00,5,0.00,0.00,0.00,0.00,0.00,0.00,'','0000-00-00 00:00:00','',0,'0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',0,'2023-03-15 13:00:48','0000-00-00 00:00:00','0000-00-00 00:00:00'),(2,'owvNO46BJBnAfD977z7D-Vu3k7gU','{\"2\":\"4\",\"8\":\"2\"}',0,0,'','321312',768.00,6,0.00,0.00,0.00,0.00,0.00,0.00,'','0000-00-00 00:00:00','',0,'0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',0,'2023-03-15 13:02:20','0000-00-00 00:00:00','0000-00-00 00:00:00'),(3,'owvNO46BJBnAfD977z7D-Vu3k7gU','{\"13\":\"1\",\"7\":\"2\"}',0,0,'{\"name\":\"张印\",\"open_id\":\"owvNO46BJBnAfD977z7D-Vu3k7gU\",\"province\":\"北京市\",\"city\":\"北京市\",\"county\":\"东城区\",\"detail\":\"2704\",\"area_code\":\"\",\"postal_code\":\"\",\"tel\":\"131\",\"is_default\":0}','1231212',384.00,3,0.00,0.00,0.00,0.00,0.00,0.00,'','0000-00-00 00:00:00','',0,'0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',0,'2023-03-15 13:02:49','0000-00-00 00:00:00','0000-00-00 00:00:00'),(4,'owvNO46BJBnAfD977z7D-Vu3k7gU','{\"1\":\"1\",\"3\":\"3\"}',0,0,'{\"name\":\"张印\",\"open_id\":\"owvNO46BJBnAfD977z7D-Vu3k7gU\",\"province\":\"北京市\",\"city\":\"北京市\",\"county\":\"东城区\",\"detail\":\"2704\",\"area_code\":\"\",\"postal_code\":\"\",\"tel\":\"131\",\"is_default\":0}','dfhgsg ',512.00,4,0.00,0.00,0.00,0.00,0.00,0.00,'','0000-00-00 00:00:00','',0,'0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',0,'2023-03-15 13:06:00','0000-00-00 00:00:00','0000-00-00 00:00:00'),(5,'owvNO46BJBnAfD977z7D-Vu3k7gU','{\"2\":\"2\",\"6\":\"3\"}',0,0,'{\"name\":\"张印\",\"open_id\":\"owvNO46BJBnAfD977z7D-Vu3k7gU\",\"province\":\"北京市\",\"city\":\"北京市\",\"county\":\"东城区\",\"detail\":\"2704\",\"area_code\":\"\",\"postal_code\":\"\",\"tel\":\"131\",\"is_default\":0}','Ddddd',640.00,5,0.00,0.00,0.00,0.00,0.00,0.00,'','0000-00-00 00:00:00','',0,'0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',0,'2023-03-15 13:07:28','0000-00-00 00:00:00','0000-00-00 00:00:00'),(6,'owvNO46BJBnAfD977z7D-Vu3k7gU','{\"5\":\"1\",\"6\":\"1\"}',0,0,'{\"name\":\"1\",\"open_id\":\"owvNO46BJBnAfD977z7D-Vu3k7gU\",\"province\":\"北京市\",\"city\":\"北京市\",\"county\":\"东城区\",\"detail\":\"31212\",\"area_code\":\"\",\"postal_code\":\"\",\"tel\":\"1\",\"is_default\":0}','吧吧吧吧吧',256.00,2,0.00,0.00,0.00,0.00,0.00,0.00,'','0000-00-00 00:00:00','',0,'0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',0,'2023-03-15 14:04:04','0000-00-00 00:00:00','0000-00-00 00:00:00');
/*!40000 ALTER TABLE `t_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_user`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `open_id` varchar(63) NOT NULL DEFAULT '' COMMENT '微信登录openid',
  `nickname` varchar(63) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像图片',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别：0 未知， 1男， 1 女',
  `birthday` varchar(20) NOT NULL DEFAULT '0' COMMENT '生日',
  `user_level` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 普通用户，1 VIP用户，2 高级VIP用户',
  `tel` varchar(20) NOT NULL DEFAULT '' COMMENT '用户手机号码',
  `deleted` tinyint(1) NOT NULL DEFAULT '1' COMMENT '逻辑删除 1默认 2删除',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_open_id` (`open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='微信商城小程序-用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_user`
--

LOCK TABLES `t_user` WRITE;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
INSERT INTO `t_user` VALUES (1,'owvNO46BJBnAfD977z7D-Vu3k7gU','','',0,'',0,'',0,'2023-03-15 06:22:05','0000-00-00 00:00:00','0000-00-00 00:00:00'),(2,'owvNO4-N4MByzdIg3QiDkn70Gezs','','',0,'',0,'',0,'2023-03-15 13:17:56','0000-00-00 00:00:00','0000-00-00 00:00:00');
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_web_user`
--

DROP TABLE IF EXISTS `t_web_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_web_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(63) NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` varchar(63) NOT NULL DEFAULT '' COMMENT '用户密码',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像图片',
  `deleted` tinyint(1) NOT NULL DEFAULT '1' COMMENT '逻辑删除 1默认 2删除',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='微信商城web后台-用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_web_user`
--

LOCK TABLES `t_web_user` WRITE;
/*!40000 ALTER TABLE `t_web_user` DISABLE KEYS */;
INSERT INTO `t_web_user` VALUES (1,'zy','123456','',1,'2023-03-13 09:50:57','2023-03-13 09:50:57',NULL),(2,'gly','123456','',1,'2023-03-13 09:50:57','2023-03-13 09:50:57',NULL);
/*!40000 ALTER TABLE `t_web_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-15 22:06:49
