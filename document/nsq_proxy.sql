/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3310
 Source Schema         : nsq_proxy

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 08/02/2022 14:56:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for nsqproxy_consume_config
-- ----------------------------
DROP TABLE IF EXISTS `nsqproxy_consume_config`;
CREATE TABLE `nsqproxy_consume_config`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `topic` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'topic名',
  `channel` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'channel名',
  `description` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `owner` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '责任人',
  `monitor_threshold` int(11) NOT NULL DEFAULT 50000 COMMENT '报警监控的阈值, 0是白名单',
  `handle_num` int(11) NOT NULL DEFAULT 2 COMMENT '消费者的并发量',
  `max_in_flight` int(11) NOT NULL DEFAULT 2 COMMENT '未返回时nsqd最大的可推送量',
  `is_requeue` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否需要重新入队，0是不需要，1是需要',
  `timeout_dial` int(11) NOT NULL DEFAULT 3590 COMMENT '超时时间',
  `timeout_read` int(11) NOT NULL DEFAULT 3590 COMMENT '超时时间-读',
  `timeout_write` int(11) NOT NULL DEFAULT 3590 COMMENT '超时时间-写',
  `invalid` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否有效，0是有效',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '添加时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniq_topic_channel`(`topic`, `channel`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '消费者的配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for nsqproxy_consume_server_map
-- ----------------------------
DROP TABLE IF EXISTS `nsqproxy_consume_server_map`;
CREATE TABLE `nsqproxy_consume_server_map`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `consumeid` int(11) NOT NULL COMMENT '消费者id',
  `serverid` int(11) NOT NULL COMMENT '服务器id',
  `weight` int(11) NULL DEFAULT 0 COMMENT '权重',
  `invalid` tinyint(4) NULL DEFAULT 0 COMMENT '是否有效，0是有效',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '添加时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `index_uq_cid_sid`(`consumeid`, `serverid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '消费者和可消费的服务器之间的关联关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for nsqproxy_message
-- ----------------------------
DROP TABLE IF EXISTS `nsqproxy_message`;
CREATE TABLE `nsqproxy_message`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `message_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `topic` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `argument` json NULL,
  `delay` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '延迟时间',
  `status` tinyint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '结果标识 0-待执行 1-执行成功 2-执行失败',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消息类型  1-及时消息  2-延迟消息',
  `response` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `update_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `indx_message_id`(`message_id`) USING BTREE,
  INDEX `index_topic`(`topic`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for nsqproxy_platform
-- ----------------------------
DROP TABLE IF EXISTS `nsqproxy_platform`;
CREATE TABLE `nsqproxy_platform`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `app_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `app_secret` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `topic` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `remark` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `update_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for nsqproxy_work_server
-- ----------------------------
DROP TABLE IF EXISTS `nsqproxy_work_server`;
CREATE TABLE `nsqproxy_work_server`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `addr` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `protocol` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'CBNSQ' COMMENT '使用的协议，支持HTTP、FastCGI、CBNSQ',
  `extra` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '扩展字段，比如协议是FastCGI时，需要传入PHP-FPM的执行的PHP文件的路径',
  `description` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '描述',
  `owner` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '责任人',
  `invalid` tinyint(4) NULL DEFAULT 0 COMMENT '是否有效，0是有效',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '添加时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `index_uq_addr`(`addr`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '所有的可以消费的服务器列表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
