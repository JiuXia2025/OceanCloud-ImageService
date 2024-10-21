/*
 Navicat Premium Dump SQL

 Source Server         : 本地3306
 Source Server Type    : MySQL
 Source Server Version : 50726 (5.7.26)
 Source Host           : localhost:3306
 Source Schema         : ocean

 Target Server Type    : MySQL
 Target Server Version : 50726 (5.7.26)
 File Encoding         : 65001

 Date: 20/10/2024 22:19:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_image
-- ----------------------------
DROP TABLE IF EXISTS `app_image`;
CREATE TABLE `app_image`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of app_image
-- ----------------------------
INSERT INTO `app_image` VALUES (1, 'CentOS7', 'centos', '/opt/iaas/centos7.qcow');
INSERT INTO `app_image` VALUES (2, 'Windows10', 'windows', '/opt/iaas/windows.img');
INSERT INTO `app_image` VALUES (3, 'centos9系统', 'centos', '/opt/iaas/centos9.qcow2');

-- ----------------------------
-- Table structure for app_user
-- ----------------------------
DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user`  (
  `id` int(11) NOT NULL,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of app_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
