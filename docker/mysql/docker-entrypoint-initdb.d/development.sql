/*
Navicat MySQL Data Transfer

Source Server         : docker
Source Server Version : 50646
Source Host           : 192.168.99.100:3306
Source Database       : development

Target Server Type    : MYSQL
Target Server Version : 50646
File Encoding         : 65001

Date: 2019-10-31 10:06:40
*/

SET FOREIGN_KEY_CHECKS=0;
CREATE DATABASE `development` CHARACTER SET 'utf8';
 
use development;
-- ----------------------------
-- Table structure for house_info
-- ----------------------------
DROP TABLE IF EXISTS `house_info`;
CREATE TABLE `house_info` (
  `id` int(11) NOT NULL COMMENT 'id',
  `url` varchar(100) NOT NULL COMMENT 'url',
  `title` varchar(100) DEFAULT NULL COMMENT '标题',
  `sub_title` varchar(255) DEFAULT NULL COMMENT '副标题',
  `region` varchar(50) DEFAULT NULL COMMENT '小区',
  `layout` varchar(50) DEFAULT NULL COMMENT '房屋户型',
  `floor` varchar(50) DEFAULT NULL COMMENT '所在楼层',
  `area_build` varchar(50) DEFAULT NULL COMMENT '建筑面积',
  `struct_house` varchar(50) DEFAULT NULL COMMENT '户型结构',
  `area_inner` varchar(50) DEFAULT NULL COMMENT '套内面积',
  `build_type` varchar(50) DEFAULT NULL COMMENT '建筑类型',
  `face` varchar(50) DEFAULT NULL COMMENT '房屋朝向',
  `struct_build` varchar(50) DEFAULT NULL COMMENT '建筑结构',
  `decoration` varchar(50) DEFAULT NULL COMMENT '装修情况',
  `elevator_ratio` varchar(50) DEFAULT NULL COMMENT '梯户比例',
  `elevator` varchar(50) DEFAULT NULL COMMENT '配备电梯',
  `property` varchar(50) DEFAULT NULL COMMENT '产权年限',
  `listing_time` varchar(50) DEFAULT NULL COMMENT '挂牌时间',
  `trading_authority` varchar(50) DEFAULT NULL COMMENT '交易权属',
  `last_transaction` varchar(50) DEFAULT NULL COMMENT '上次交易',
  `housing_purposes` varchar(50) DEFAULT NULL COMMENT '房屋用途',
  `house_year` varchar(50) DEFAULT NULL COMMENT '房屋年限',
  `property_rights` varchar(50) DEFAULT NULL COMMENT '产权所属',
  `mortgage_info` varchar(50) DEFAULT NULL COMMENT '抵押信息',
  `document_photo` varchar(50) DEFAULT NULL COMMENT '房本备件',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
