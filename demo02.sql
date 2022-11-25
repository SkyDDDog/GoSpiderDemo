/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 8.0.27 : Database - demo02
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`demo02` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `demo02`;

/*Table structure for table `bili_comment` */

DROP TABLE IF EXISTS `bili_comment`;

CREATE TABLE `bili_comment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `del_flag` char(1) DEFAULT '0' COMMENT '逻辑删除',
  `create_date` datetime DEFAULT NULL COMMENT '创建时间',
  `update_date` datetime DEFAULT NULL COMMENT '更新时间',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `uid` bigint DEFAULT NULL COMMENT '用户uid',
  `message` longtext COMMENT '评论内容',
  `push_time` datetime DEFAULT NULL COMMENT '发布时间',
  `likes` int DEFAULT NULL COMMENT '点赞数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

/*Data for the table `bili_comment` */

/*Table structure for table `bili_user` */

DROP TABLE IF EXISTS `bili_user`;

CREATE TABLE `bili_user` (
  `del_flag` char(1) DEFAULT NULL COMMENT '逻辑删除',
  `create_date` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_date` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `space` varchar(255) DEFAULT NULL COMMENT '个人空间',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `name` varchar(255) DEFAULT NULL COMMENT '用户名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

/*Data for the table `bili_user` */

/*Table structure for table `fzu` */

DROP TABLE IF EXISTS `fzu`;

CREATE TABLE `fzu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `del_flag` char(1) DEFAULT NULL COMMENT '逻辑删除',
  `create_date` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_date` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `author` varchar(255) DEFAULT NULL COMMENT '作者',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `read_num` bigint DEFAULT NULL COMMENT '阅读量',
  `content` longtext COMMENT '正文',
  `publish_time` date DEFAULT NULL COMMENT '发布时间',
  `link` varchar(255) DEFAULT NULL COMMENT '链接',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


/*Table structure for table `sub_comment` */

DROP TABLE IF EXISTS `sub_comment`;

CREATE TABLE `sub_comment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `del_flag` char(1) DEFAULT '0' COMMENT '逻辑删除',
  `create_date` datetime DEFAULT NULL COMMENT '创建时间',
  `update_date` datetime DEFAULT NULL COMMENT '更新时间',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `cid` bigint DEFAULT NULL COMMENT '主评论id',
  `uid` bigint DEFAULT NULL COMMENT '用户uid',
  `message` longtext COMMENT '评论内容',
  `push_time` datetime DEFAULT NULL COMMENT '发布时间',
  `likes` int DEFAULT NULL COMMENT '点赞数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

/*Data for the table `sub_comment` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
