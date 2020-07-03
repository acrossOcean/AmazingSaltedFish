/*
 Navicat Premium Data Transfer

 Source Server         : local-AmazingSaltedFish
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 192.168.137.128:3306
 Source Schema         : amazing_salted_fish

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 03/07/2020 17:28:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for core_generator_params
-- ----------------------------
DROP TABLE IF EXISTS `core_generator_params`;
CREATE TABLE `core_generator_params`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `generator_id` int(11) NOT NULL COMMENT '所属生成器id',
  `location` int(11) NOT NULL COMMENT '位置, 1.入参 2.出参',
  `param_type` int(11) NOT NULL COMMENT '字段类型',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字段名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字段注释',
  `is_list` tinyint(1) NOT NULL COMMENT '是否为 列表',
  `sort` int(11) DEFAULT NULL COMMENT '排序位置',
  `struct_id` int(11) DEFAULT NULL COMMENT '对应结构体ID, 当 FType 为 ParamTypeStruct 时 有用',
  `map_key_param_id` int(11) DEFAULT NULL COMMENT '当ParamType 为 map 时',
  `map_value_param_id` int(11) DEFAULT NULL COMMENT '当ParamType 为 map 时',
  `is_const` tinyint(1) NOT NULL COMMENT '是否为固定值',
  `const_data` binary(1) DEFAULT NULL COMMENT '固定值数据',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for core_generators
-- ----------------------------
DROP TABLE IF EXISTS `core_generators`;
CREATE TABLE `core_generators`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `language` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '实现所用语言',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名字',
  `version` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'Version. 实现对应版本',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '注释',
  `data` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '实现内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for field_info
-- ----------------------------
DROP TABLE IF EXISTS `field_info`;
CREATE TABLE `field_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '对应 父级 ID (可能还是字段, 也可能是结构体)',
  `field_type` int(11) NOT NULL COMMENT '字段类型',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字段名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字段注释',
  `is_list` tinyint(1) NOT NULL COMMENT '是否为 列表',
  `sort` int(11) NOT NULL COMMENT '排序位置',
  `struct_id` int(11) DEFAULT NULL COMMENT '对应结构体ID, 当 FType 为 FieldTypeStruct 时 有用',
  `map_key_field_id` int(11) DEFAULT NULL COMMENT '当fieldType 为 map 时',
  `map_value_field_id` int(11) DEFAULT NULL COMMENT '当fieldType 为 map 时',
  `is_const` tinyint(1) NOT NULL COMMENT '如果是固定值, 那么记录',
  `const_data` binary(1) DEFAULT NULL COMMENT '如果是固定值, 那么记录',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for link_define
-- ----------------------------
DROP TABLE IF EXISTS `link_define`;
CREATE TABLE `link_define`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_define_id` int(11) DEFAULT NULL COMMENT '所属工程ID',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '链名称',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '注释',
  `is_shared` tinyint(1) NOT NULL COMMENT '是否共享, 不共享的话 只能当前工程可用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for link_instance
-- ----------------------------
DROP TABLE IF EXISTS `link_instance`;
CREATE TABLE `link_instance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_instance_id` int(11) DEFAULT NULL COMMENT '所属工程ID',
  `define_id` int(11) NOT NULL COMMENT '所用 链 定义 id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '链名称',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '注释',
  `is_shared` tinyint(1) NOT NULL COMMENT '是否共享, 不共享的话 只能当前工程可用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for link_param_define
-- ----------------------------
DROP TABLE IF EXISTS `link_param_define`;
CREATE TABLE `link_param_define`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `link_define_id` int(11) NOT NULL COMMENT '所属 链 id',
  `location` int(11) NOT NULL COMMENT '位置, 1.入参 2.出参',
  `param_type` int(11) NOT NULL COMMENT '字段类型',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字段名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字段注释',
  `is_list` tinyint(1) NOT NULL COMMENT '是否为 列表',
  `sort` int(11) DEFAULT NULL COMMENT '排序位置',
  `struct_id` int(11) DEFAULT NULL COMMENT '对应结构体ID, 当 FType 为 ParamTypeStruct 时 有用',
  `map_key_param_id` int(11) DEFAULT NULL COMMENT '当ParamType 为 map 时',
  `map_value_param_id` int(11) DEFAULT NULL COMMENT '当ParamType 为 map 时',
  `is_const` tinyint(1) NOT NULL COMMENT '如果是固定值, 那么记录',
  `const_data` binary(1) DEFAULT NULL COMMENT '如果是固定值, 那么记录',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for link_param_instance
-- ----------------------------
DROP TABLE IF EXISTS `link_param_instance`;
CREATE TABLE `link_param_instance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `define_id` int(11) NOT NULL COMMENT '所用 链参数 定义 id',
  `link_instance_id` int(11) NOT NULL COMMENT '所属 链 instance id',
  `location` int(11) NOT NULL COMMENT '位置, 1.入参 2.出参',
  `param_type` int(11) NOT NULL COMMENT '字段类型',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字段名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字段注释',
  `input_type` int(11) DEFAULT NULL COMMENT '入参来源',
  `input_const` binary(1) DEFAULT NULL COMMENT '如果是来自确定值, 那么记录这个值',
  `input_var_is_link` tinyint(1) DEFAULT NULL COMMENT '如果来自变量, 那么记录是来自 其它链还是其它Node',
  `input_var_define_id` int(11) DEFAULT NULL COMMENT '如果来自变量, 那么记录来源',
  `input_var_instance_id` int(11) DEFAULT NULL COMMENT '如果来自变量, 那么记录来源',
  `output_type` int(11) DEFAULT NULL COMMENT '返回值类型 (如果返回变量, 那么不需要记录, 到时候直接用就可以)',
  `output_const` binary(1) DEFAULT NULL COMMENT '如果是确定值, 那么记录这个值',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for node_define
-- ----------------------------
DROP TABLE IF EXISTS `node_define`;
CREATE TABLE `node_define`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `belong_link_define_id` int(11) DEFAULT NULL COMMENT '所属 链 ID',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '方法名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '注释',
  `pre_node_define_id` int(11) NOT NULL COMMENT '前一个NodeID, 第一个写-1',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for node_instance
-- ----------------------------
DROP TABLE IF EXISTS `node_instance`;
CREATE TABLE `node_instance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `belong_link_instance_id` int(11) DEFAULT NULL COMMENT '所属 链 ID',
  `define_id` int(11) NOT NULL COMMENT '所用 node define id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '方法名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '注释',
  `is_link` tinyint(1) NOT NULL COMMENT '是否是链, 如果不是链, 那么就是节点, 可以是链, 也可以是节点, 如果是链的话,对应记录下 链的ID, 如果是节点,那么记录节点的生成代码实现方式',
  `link_instance_id` int(11) DEFAULT NULL COMMENT '如果是链, 那么需要记录 link instance 的ID',
  `generator_id` int(11) DEFAULT NULL COMMENT '实现方式ID, 对应 \"core_generator\" id',
  `generator_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '实现方式名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for node_param_define
-- ----------------------------
DROP TABLE IF EXISTS `node_param_define`;
CREATE TABLE `node_param_define`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `node_define_id` int(11) NOT NULL COMMENT '所属节点id',
  `location` int(11) NOT NULL COMMENT '位置, 1.入参 2.出参',
  `param_type` int(11) NOT NULL COMMENT '字段类型',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字段名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字段注释',
  `is_list` tinyint(1) NOT NULL COMMENT '是否为 列表',
  `sort` int(11) DEFAULT NULL COMMENT '排序位置',
  `struct_id` int(11) DEFAULT NULL COMMENT '对应结构体ID, 当 FType 为 ParamTypeStruct 时 有用',
  `map_key_param_id` int(11) DEFAULT NULL COMMENT '当ParamType 为 map 时',
  `map_value_param_id` int(11) DEFAULT NULL COMMENT '当ParamType 为 map 时',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for node_param_instance
-- ----------------------------
DROP TABLE IF EXISTS `node_param_instance`;
CREATE TABLE `node_param_instance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `define_id` int(11) NOT NULL COMMENT '对应定义ID',
  `node_define_id` int(11) NOT NULL COMMENT '所属节点定义id',
  `node_instance_id` int(11) NOT NULL COMMENT '所属节点id',
  `location` int(11) NOT NULL COMMENT '位置, 1.入参 2.出参',
  `param_type` int(11) NOT NULL COMMENT '字段类型',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字段名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字段注释',
  `input_type` int(11) DEFAULT NULL COMMENT '入参来源',
  `input_const` binary(1) DEFAULT NULL COMMENT '如果是来自确定值, 那么记录这个值',
  `input_var_is_link` tinyint(1) DEFAULT NULL COMMENT '如果来自变量, 那么记录是来自 其它link还是其它Node',
  `input_var_define_id` int(11) DEFAULT NULL COMMENT '如果来自变量, 那么记录来源',
  `input_var_instance_id` int(11) DEFAULT NULL COMMENT '如果来自变量, 那么记录来源',
  `output_type` int(11) DEFAULT NULL COMMENT '返回值类型 (如果返回变量, 那么不需要记录, 到时候直接用就可以)',
  `output_const` binary(1) DEFAULT NULL COMMENT '如果是确定值, 那么记录这个值',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for project_define
-- ----------------------------
DROP TABLE IF EXISTS `project_define`;
CREATE TABLE `project_define`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '工程名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '注释',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for project_instance
-- ----------------------------
DROP TABLE IF EXISTS `project_instance`;
CREATE TABLE `project_instance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `define_id` int(11) NOT NULL COMMENT '对应 定义 id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '工程名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '注释',
  `generate_language` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '实现工程所用语言及版本',
  `first_link_define_id` int(11) DEFAULT NULL COMMENT '第一个 link 信息',
  `first_link_instance_id` int(11) DEFAULT NULL COMMENT '第一个 link 信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for project_param_define
-- ----------------------------
DROP TABLE IF EXISTS `project_param_define`;
CREATE TABLE `project_param_define`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_define_id` int(11) NOT NULL COMMENT '所属项目id',
  `location` int(11) NOT NULL COMMENT '位置, 1.入参 2.出参',
  `param_type` int(11) DEFAULT NULL COMMENT '字段类型',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字段名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字段注释',
  `is_list` tinyint(1) NOT NULL COMMENT '是否为 列表',
  `sort` int(11) DEFAULT NULL COMMENT '排序位置',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for project_param_instance
-- ----------------------------
DROP TABLE IF EXISTS `project_param_instance`;
CREATE TABLE `project_param_instance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_instance_id` int(11) NOT NULL COMMENT '所属项目定义ID',
  `location` int(11) NOT NULL COMMENT '位置, 1.入参 2.出参',
  `param_type` int(11) NOT NULL COMMENT '字段类型',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字段名',
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字段注释',
  `input_type` int(11) DEFAULT NULL COMMENT '入参来源',
  `input_const` binary(1) DEFAULT NULL COMMENT '如果是来自确定值, 那么记录这个值',
  `input_var_is_link` tinyint(1) DEFAULT NULL COMMENT '如果来自变量, 那么记录是来自 其它link还是其它Node',
  `input_var_define_id` int(11) DEFAULT NULL COMMENT '如果来自变量, 那么记录来源',
  `input_var_instance_id` int(11) DEFAULT NULL COMMENT '如果来自变量, 那么记录来源',
  `output_type` int(11) DEFAULT NULL COMMENT '返回值类型 (如果返回变量, 那么不需要记录, 到时候直接用就可以)',
  `output_const` binary(1) DEFAULT NULL COMMENT '如果是确定值, 那么记录这个值',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for struct_info
-- ----------------------------
DROP TABLE IF EXISTS `struct_info`;
CREATE TABLE `struct_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
