/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : hub

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 31/03/2023 23:26:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for plugin
-- ----------------------------
DROP TABLE IF EXISTS `plugin`;
CREATE TABLE `plugin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `domain` varchar(255) NOT NULL COMMENT 'plugin domain',
  `name` varchar(255) NOT NULL COMMENT 'plugin name',
  `description` varchar(255) NOT NULL COMMENT 'plugin description',
  `auth_type` varchar(128) NOT NULL COMMENT 'plugin auth type',
  `logo_url` varchar(255) NOT NULL COMMENT 'plugin logo url',
  `contact_email` varchar(255) NOT NULL COMMENT 'plugin contact email',
  `organization` varchar(255) NOT NULL COMMENT 'plugin organization',
  `api_type` varchar(255) NOT NULL COMMENT 'plugin api type',
  `api_url` varchar(255) NOT NULL COMMENT 'plugin api url',
  `label` json DEFAULT NULL COMMENT 'plugin label',
  `state` varchar(255) NOT NULL COMMENT 'plugin state,published or unpublished',
  `install_num` int NOT NULL DEFAULT '0' COMMENT 'plugin install num',
  `score` float NOT NULL DEFAULT '0' COMMENT 'plugin score',
  `heat` int NOT NULL DEFAULT '0' COMMENT 'plugin heat',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for plugin_score
-- ----------------------------
DROP TABLE IF EXISTS `plugin_score`;
CREATE TABLE `plugin_score` (
  `id` int NOT NULL AUTO_INCREMENT,
  `plugin_id` int NOT NULL COMMENT 'plugin id',
  `score` int NOT NULL COMMENT 'plugin score',
  `comments` varchar(1024) DEFAULT NULL COMMENT 'plugin score comments',
  `user_id` int DEFAULT NULL COMMENT 'user id for score',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
