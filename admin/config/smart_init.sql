/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : smart

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 30/08/2024 15:32:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `account` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登陆账号',
  `encryption` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码密文',
  `salt_key` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码盐值（AES加密KEY）',
  `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '姓名',
  `dept_id` bigint(0) NOT NULL COMMENT '部门ID',
  `is_leader` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门职位，枚举：0_部门员工 1_部门领导',
  `permission_type` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '权限类型，枚举：0_继承部门 1_本部门 2_本人 3_全局',
  `pwd_exp_time` datetime(0) NULL DEFAULT NULL COMMENT '密码过期时间，创建即过期',
  `last_login_time` datetime(0) NULL DEFAULT NULL COMMENT '最后登陆时间，为空表示创建',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识，枚举：0_可变更 1_禁止变更',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `account_uni`(`account`) USING BTREE COMMENT '登陆账号唯一',
  INDEX `dept_query`(`dept_id`) USING BTREE COMMENT '部门查询'
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登陆账号' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account
-- ----------------------------
INSERT INTO `account` VALUES (2, 'admin', 'PUNce7DqTrAzapC9UizQnA==', 'tDdVGjQK7nDXxjxP', '管理员', 1, '1', '3', '2024-10-29 09:38:01', NULL, '0', '0', '2024-07-08 08:56:33', '2024-07-31 09:38:01');
INSERT INTO `account` VALUES (3, 'test1', 'DEidKlqe7jYrqnlthxjbJA==', 'GaNnzTVHodLqigTl', '', 1, '0', '0', '2024-10-29 10:24:04', NULL, '0', '0', '2024-07-08 10:19:33', '2024-07-31 21:27:28');
INSERT INTO `account` VALUES (4, 'CC', 'qEXB5Jx3YtvyXAkOl44vZA==', 'Q1FfGM3ZSeyByeed', '', 2, '0', '0', '2024-07-08 10:45:47', NULL, '0', '0', '2024-07-08 10:45:47', '2024-07-31 21:27:33');

-- ----------------------------
-- Table structure for account_role
-- ----------------------------
DROP TABLE IF EXISTS `account_role`;
CREATE TABLE `account_role`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `account_id` bigint(0) NOT NULL COMMENT '账号ID',
  `role_id` bigint(0) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `account_role_uni`(`account_id`, `role_id`) USING BTREE COMMENT '账号角色唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '账号角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account_role
-- ----------------------------
INSERT INTO `account_role` VALUES (5, 2, 5);
INSERT INTO `account_role` VALUES (6, 3, 5);
INSERT INTO `account_role` VALUES (7, 3, 6);
INSERT INTO `account_role` VALUES (8, 4, 5);
INSERT INTO `account_role` VALUES (9, 4, 6);

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
  `pid` bigint(0) NOT NULL DEFAULT 0 COMMENT '父级部门ID，租户创建时默认创建根部门，父级ID=0',
  `sort` int(0) NOT NULL DEFAULT 0 COMMENT '同级部门排序',
  `permission_type` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '权限类型，枚举：0_本部门与子部门 1_本部门 2_个人 3_全局 4_指定部门 5_指定人',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识，枚举：0_可变更 1_禁止变更',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `dept_uni`(`name`) USING BTREE COMMENT '部门名唯一',
  INDEX `dept_pid_query`(`pid`) USING BTREE COMMENT '部门树构建'
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '集团部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dept
-- ----------------------------
INSERT INTO `dept` VALUES (1, '站线集团', 0, 0, '0', '0', '0', NULL, NULL);
INSERT INTO `dept` VALUES (2, '战斗士传播公司', 1, 1, '1', '0', '0', '2024-06-29 12:13:26', '2024-06-29 12:13:35');
INSERT INTO `dept` VALUES (5, '极光线数据公司', 1, 0, '0', '0', '0', '2024-06-29 12:15:45', '2024-06-29 12:15:45');
INSERT INTO `dept` VALUES (7, '迁移A', 2, 0, '0', '0', '0', '2024-07-08 10:20:05', '2024-07-08 10:45:59');
INSERT INTO `dept` VALUES (8, '迁移B', 7, 0, '0', '0', '0', '2024-07-08 10:20:13', '2024-07-08 10:20:13');
INSERT INTO `dept` VALUES (9, '迁移C', 5, 0, '0', '0', '0', '2024-07-08 10:20:22', '2024-07-08 10:20:22');

-- ----------------------------
-- Table structure for dict
-- ----------------------------
DROP TABLE IF EXISTS `dict`;
CREATE TABLE `dict`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `group_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典分组KEY',
  `label` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段名称',
  `label_en` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段名称（英文）',
  `choose` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '是否可被选择 0可选择 1不可选择',
  `val` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典值（字符型）',
  `pid` bigint(0) NULL DEFAULT NULL COMMENT '父级字典ID 默认 0（根数据）',
  `sort` smallint(0) NULL DEFAULT NULL COMMENT '字典排序',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典描述',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `dict_uni`(`group_key`, `val`) USING BTREE COMMENT '同一分组值唯一',
  INDEX `dict_pgs`(`group_key`, `pid`, `sort`) USING BTREE COMMENT '父级ID，分组过滤，排序'
) ENGINE = InnoDB AUTO_INCREMENT = 55 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dict
-- ----------------------------
INSERT INTO `dict` VALUES (1, 'accountPermission', '继承部门', 'Same department', '0', '0', 1, 0, '账号权限：继承部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (2, 'accountPermission', '本部门', 'Main department', '0', '1', 1, 0, '账号权限：本部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (3, 'accountPermission', '本人', 'Personal', '0', '2', 1, 0, '账号权限：本人', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (4, 'accountPermission', '全部', 'All', '0', '3', 1, 0, '账号权限：全部', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (5, 'accountStatus', '正常', 'Normal', '0', '0', 1, 0, '账号状态：正常', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (6, 'accountStatus', '锁定', 'Locked', '0', '1', 1, 0, '账号状态：锁定', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (7, 'deptAccount', '部门员工', 'Department staff', '0', '0', 1, 0, '部门成员：部门员工', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (8, 'deptAccount', '部门领导', 'Department heads', '0', '1', 1, 0, '部门成员：部门领导', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (9, 'deptPermission', '本部门及子部门', 'Main and sub departments', '0', '0', 1, 0, '部门权限：本部门与子部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (10, 'deptPermission', '本部门', 'Main department', '0', '1', 1, 0, '部门权限：本部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (11, 'deptPermission', '本人', 'Personal', '0', '2', 1, 0, '部门权限：本人', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (12, 'deptPermission', '全部', 'All', '0', '3', 1, 0, '字典分组：部门权限类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (13, 'deptPermission', '指定部门', 'Choose department', '1', '4', 1, 0, '字典分组：部门权限类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (14, 'deptPermission', '指定人', 'Choose person', '1', '5', 1, 0, '字典分组：部门权限类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (15, 'deptToType', '部门合入', 'Fit into', '0', '0', 1, 0, '合并模式：部门合入', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (16, 'deptToType', '部门并入', 'Merge into', '0', '1', 1, 0, '合并模式：部门并入', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (17, 'dictGroup', '账号数据权限', 'Account data permissions', '0', 'accountPermission', 1, 0, '字典分组：账号数据权限', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (18, 'dictGroup', '账号状态', 'Account status', '0', 'accountStatus', 1, 0, '字典分组：账号状态', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (19, 'dictGroup', '部门账号类型', 'Department account type', '0', 'deptAccount', 1, 0, '字典分组：部门账号类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (20, 'dictGroup', '部门数据权限', 'deptPermission', '0', 'deptPermission', 1, 0, '字典分组：部门数据权限', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (21, 'dictGroup', '部门合并类型', 'Department merger type', '0', 'deptToType', 1, 0, '字典分组：部门合并类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (22, 'dictGroup', '字典分组', 'Dict Group', '0', 'dictGroup', 1, 0, '字典分组：字典分组', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (23, 'dictGroup', '启用状态', 'Open Status', '0', 'openStatus', 1, 2, '字典分组：启用状态', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (24, 'dictGroup', '权限等级', 'Permission Level', '0', 'permissionLevel', 1, 3, '字典分组：权限等级', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (25, 'dictGroup', '响应类型', 'Response Type', '0', 'responseType', 1, 4, '字典分组：响应类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (26, 'dictGroup', '路由类型', 'Router Type', '0', 'routerType', 1, 5, '字典分组：路由类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (27, 'dictGroup', '业务编码', 'Service code', '0', 'serviceCode', 1, 1, '字典分组：业务编码', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (28, 'dictGroup', '时间周期', 'Time unit', '0', 'timeUnit', 1, 6, '字典分组：时间周期', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (29, 'openStatus', '启用', 'Enable', '0', '0', 1, 0, '启用状态：启用', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (30, 'openStatus', '禁用', 'Disable', '0', '1', 1, 0, '启用状态：禁用', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (31, 'permissionLevel', '系统', 'System', '1', '0', 1, 0, '权限等级：系统', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (32, 'permissionLevel', '模块', 'Model', '1', '1', 1, 0, '权限等级：模块', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (33, 'permissionLevel', '页面', 'Page', '1', '2', 1, 0, '权限等级：页面', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (34, 'permissionLevel', '按钮', 'Button', '1', '3', 1, 0, '权限等级：按钮', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (35, 'responseType', '异常', 'Error', '0', 'E', 1, 3, '响应类型：异常', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (36, 'responseType', '失败', 'Fail', '0', 'F', 1, 2, '响应类型：失败', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (37, 'responseType', '成功', 'Success', '0', 'S', 1, 0, '响应类型：成功', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (38, 'routerType', '鉴权路由', 'Authentication Router', '0', '0', 1, 0, '路由类型：需要登陆授权', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (39, 'routerType', '白名单路由', 'Whitelist Router', '0', '1', 1, 0, '路由类型：白名单', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (40, 'serviceCode', '框架', 'Frame', '0', '0', 1, 0, '模块：基础框架', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (41, 'serviceCode', '字典', 'Dict', '0', '1', 1, 7, '模块：字典', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (42, 'serviceCode', '响应', 'Response', '0', '2', 1, 6, '模块：响应配置', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (43, 'serviceCode', '路由', 'Router', '0', '3', 1, 5, '模块：接口路由', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (44, 'serviceCode', '权限', 'Permissions', '0', '4', 1, 4, '模块：权限集', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (45, 'serviceCode', '账号', 'Account', '0', '5', 1, 3, '模块：账号', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (46, 'serviceCode', '角色', 'Role', '0', '6', 1, 2, '模块：角色', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (47, 'serviceCode', '部门', 'Dept', '0', '7', 1, 1, '模块：部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (48, 'serviceCode', '配置', 'Config', '0', '8', 1, 8, '模块：配置', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (49, 'timeUnit', '默认(毫秒)', 'Default (milliseconds)', '0', '0', 1, 0, '时间单元：默认(毫秒)', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (50, 'timeUnit', '秒', 'Second', '0', '1', 1, 0, '时间单元：秒', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (51, 'timeUnit', '分', 'Point', '0', '2', 1, 0, '时间单元：分', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (52, 'timeUnit', '时', 'Hour', '0', '3', 1, 0, '时间单元：时', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (53, 'timeUnit', '日', 'Day', '0', '4', 1, 0, '时间单元：日', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (54, 'timeUnit', '月', 'Month', '0', '5', 1, 0, '时间单元：月', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (55, 'timeUnit', '年', 'Year', '0', '6', 1, 0, '时间单元：年', '1', '0', NULL, NULL);

-- ----------------------------
-- Table structure for login_record
-- ----------------------------
DROP TABLE IF EXISTS `login_record`;
CREATE TABLE `login_record`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `account_id` bigint(0) NOT NULL COMMENT '账号ID',
  `login_type` smallint(0) NULL DEFAULT NULL COMMENT '登陆类型 1平台账号登录',
  `login_time` datetime(0) NULL DEFAULT NULL COMMENT '登陆时间',
  `token` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆Token',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0登陆成功 1主动登出 2被动登出',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `login_time_query`(`mark`, `account_id`, `login_time`) USING BTREE COMMENT '登陆时间定时任务批量查询'
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登陆记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of login_record
-- ----------------------------
INSERT INTO `login_record` VALUES (1, 1, 1, '2024-07-10 08:33:10', NULL, '0', '0', NULL, NULL);
INSERT INTO `login_record` VALUES (2, 1, 1, '2024-07-09 08:33:13', NULL, '0', '0', NULL, NULL);
INSERT INTO `login_record` VALUES (3, 2, 1, '2024-07-18 19:42:40', 'ozt4sSNKUZNeDYqDIlPqRgxMDyLtqAuj', '2', '0', '2024-07-18 19:42:40', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (4, 2, 1, '2024-07-18 19:45:35', 'MeO30ygKMWfaMizC6qNbVsttyrq4080O', '2', '0', '2024-07-18 19:45:35', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (5, 2, 1, '2024-07-18 19:49:26', 'ABHJcocP5ZP37NjWMboN8OSjCqtAO8Ki', '2', '0', '2024-07-18 19:49:26', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (6, 2, 1, '2024-07-18 20:26:07', 'GS7hK2Mc2l86o5ao5nkP6ZUUVBSzUa85', '2', '0', '2024-07-18 20:26:07', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (7, 2, 1, '2024-07-18 20:32:58', 'DH8v19wNIqqycu0FUIg09b4fFQ6pT0UU', '2', '0', '2024-07-18 20:32:58', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (8, 2, 1, '2024-07-23 15:25:23', '2Y7H3yOJAzq5pm8Qd5jXown6awCZ4EGp', '2', '0', '2024-07-23 15:25:23', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (9, 2, 1, '2024-07-23 15:27:00', 'ScQ8pCoqGzB9a6dHX6vhzjYqiKOqN8IZ', '2', '0', '2024-07-23 15:27:00', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (10, 2, 1, '2024-07-23 15:28:01', 'Kbe6xcJXVlxcizNkZ1vUk21KjJGxf4YX', '2', '0', '2024-07-23 15:28:01', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (11, 2, 1, '2024-07-23 15:28:28', 'shsDG0RSTxJ92PILCXV1bhoBefBlokQd', '2', '0', '2024-07-23 15:28:28', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (12, 2, 1, '2024-07-23 15:31:42', 'EKBKUycyJyMw4nQWonX8QaTICQJq1RBE', '2', '0', '2024-07-23 15:31:42', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (13, 2, 1, '2024-07-23 15:31:59', 'nukjDaeDp2pwuzJgKVmP4AEyaY3gGyb2', '2', '0', '2024-07-23 15:31:59', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (14, 2, 1, '2024-07-23 15:32:55', 'D7EohMeIaPspseJFb3A0thkUCcF9cGX2', '2', '0', '2024-07-23 15:32:55', '2024-07-23 15:34:11');
INSERT INTO `login_record` VALUES (15, 2, 1, '2024-07-23 15:34:11', 'IHIxRBQFzhdWGfY0ICsAjXVFih3tbia2', '2', '0', '2024-07-23 15:34:11', '2024-07-23 15:36:56');
INSERT INTO `login_record` VALUES (16, 2, 1, '2024-07-23 15:36:56', 'msm8itWudw6qy3GHXVgZWbTxKGmDa3a5', '2', '0', '2024-07-23 15:36:56', '2024-07-23 15:41:35');
INSERT INTO `login_record` VALUES (17, 2, 1, '2024-07-23 15:41:35', 'MKz68Ojoh8yUrYwkbMCcPDHx8lcfGEig', '2', '0', '2024-07-23 15:41:35', '2024-07-23 17:56:09');
INSERT INTO `login_record` VALUES (18, 2, 1, '2024-07-23 17:56:09', 'RQUB81vefbGdOtCWjx84z1aH8LSMihGc', '2', '0', '2024-07-23 17:56:09', '2024-07-30 09:31:02');
INSERT INTO `login_record` VALUES (19, 2, 1, '2024-07-30 09:31:02', 'EjWCAi9kHAIIAcH5m66Aze1MgFrwcm5R', '2', '0', '2024-07-30 09:31:02', '2024-07-30 09:35:22');
INSERT INTO `login_record` VALUES (20, 2, 1, '2024-07-30 09:35:22', 'hFl2kG2iHLkb0K3z9zuzHl1XB7VfmdC1', '2', '0', '2024-07-30 09:35:22', '2024-07-30 09:46:13');
INSERT INTO `login_record` VALUES (21, 2, 1, '2024-07-30 09:46:13', '53kKZQncWSOBYND22U2xOfr1UKD9gjRw', '2', '0', '2024-07-30 09:46:13', '2024-07-30 09:51:48');
INSERT INTO `login_record` VALUES (22, 2, 1, '2024-07-30 09:51:48', 'fx3SldxOpJ7A6QnI8AT5V9rBeaDGhTAX', '2', '0', '2024-07-30 09:51:48', '2024-07-31 09:37:35');
INSERT INTO `login_record` VALUES (23, 2, 1, '2024-07-31 09:37:35', 'u8jqN7u5pf4N7nHOqdQG9OE6usaGMB5h', '2', '0', '2024-07-31 09:37:35', '2024-07-31 09:39:27');
INSERT INTO `login_record` VALUES (24, 2, 1, '2024-07-31 09:39:27', 'jPjHTXSUa4I0SoSNLH8ozidib462w25g', '2', '0', '2024-07-31 09:39:27', '2024-07-31 09:43:22');
INSERT INTO `login_record` VALUES (25, 2, 1, '2024-07-31 09:43:22', 'obSa4MoDC4LT1CAGhDIBuIJaSU1wzvKl', '2', '0', '2024-07-31 09:43:22', '2024-07-31 10:33:21');
INSERT INTO `login_record` VALUES (26, 3, 1, '2024-07-31 10:20:30', 'dOLoFoyd27F5PG7FLVai09FDFH1hWxMu', '0', '0', '2024-07-31 10:20:30', '2024-07-31 10:20:30');
INSERT INTO `login_record` VALUES (27, 2, 1, '2024-07-31 10:33:21', 'ln15EUqOHYcrhR343nnkhbmUgjvnL3CG', '2', '0', '2024-07-31 10:33:21', '2024-07-31 10:34:12');
INSERT INTO `login_record` VALUES (28, 2, 1, '2024-07-31 10:34:12', 'W0tQ2zHxGC66URJNhQMeF0UnjwKKMH6d', '2', '0', '2024-07-31 10:34:12', '2024-07-31 10:35:53');
INSERT INTO `login_record` VALUES (29, 2, 1, '2024-07-31 10:35:53', 'igSguijMK0HVdS20FnryftcMla4jit4V', '2', '0', '2024-07-31 10:35:53', '2024-07-31 10:36:22');
INSERT INTO `login_record` VALUES (30, 2, 1, '2024-07-31 10:36:22', 'CfTb5EsTclkObdDbuxljM1oEoJviJKos', '2', '0', '2024-07-31 10:36:22', '2024-07-31 10:36:36');
INSERT INTO `login_record` VALUES (31, 2, 1, '2024-07-31 10:36:36', 'Hwmo1GlGg9zEWncJ1CIm1l1KUexxXkft', '2', '0', '2024-07-31 10:36:36', '2024-07-31 14:54:32');
INSERT INTO `login_record` VALUES (32, 2, 1, '2024-07-31 14:54:32', '5H7Lhuxd9GZdcgPruPLcBG5l7mJxuSTD', '2', '0', '2024-07-31 14:54:32', '2024-07-31 16:44:48');
INSERT INTO `login_record` VALUES (33, 2, 1, '2024-07-31 16:44:48', 'BN0ZR30rdZsJRFH86p9kFJGqbmdqBgXw', '2', '0', '2024-07-31 16:44:48', '2024-07-31 21:18:08');
INSERT INTO `login_record` VALUES (34, 2, 1, '2024-07-31 21:18:08', 'ISVMc7dQx4QHLBidHRGy4phZ2Io6vppQ', '2', '0', '2024-07-31 21:18:08', '2024-07-31 21:20:44');
INSERT INTO `login_record` VALUES (35, 2, 1, '2024-07-31 21:20:44', 'Nh4wSC8ZXdeHGsuNskBCvr5JY3aQEBGF', '2', '0', '2024-07-31 21:20:44', '2024-07-31 23:01:35');
INSERT INTO `login_record` VALUES (36, 2, 1, '2024-07-31 23:01:35', 'pqNY4GdiWx4uGiwOv7JpRSe8cWq1L5M0', '2', '0', '2024-07-31 23:01:35', '2024-08-16 17:51:46');
INSERT INTO `login_record` VALUES (37, 2, 1, '2024-08-16 17:51:46', 'PKkLBbWLLGwtMPer6WRhIl8mCBa8vP8R', '2', '0', '2024-08-16 17:51:46', '2024-08-16 18:16:03');
INSERT INTO `login_record` VALUES (38, 2, 1, '2024-08-16 18:16:03', 'GK9ITvCcY7ZUOvHS8deUkIbQwZSI4GPC', '2', '0', '2024-08-16 18:16:03', '2024-08-23 14:05:51');
INSERT INTO `login_record` VALUES (39, 2, 1, '2024-08-23 14:05:51', 'ZYqebBpbqdEfs2LZisJEq1gdRWgeq2Z4', '2', '0', '2024-08-23 14:05:51', '2024-08-23 15:19:25');
INSERT INTO `login_record` VALUES (40, 2, 1, '2024-08-23 15:19:25', 'ljcGPmRL1MIvkQ94dIqw5SMC2utB2ctg', '2', '0', '2024-08-23 15:19:25', '2024-08-23 15:22:53');
INSERT INTO `login_record` VALUES (41, 2, 1, '2024-08-23 15:22:53', 'UqHiWkx7Kp3bJASLDQ16XEZKCEqtbs8L', '2', '0', '2024-08-23 15:22:53', '2024-08-23 15:36:57');
INSERT INTO `login_record` VALUES (42, 2, 1, '2024-08-23 15:36:57', 'pIjAfbMFjAbbsdUHcxTCEDSlYJQHWnmo', '2', '0', '2024-08-23 15:36:57', '2024-08-23 15:37:31');
INSERT INTO `login_record` VALUES (43, 2, 1, '2024-08-23 15:37:31', 'oAZLYOihoy1U2IydCRLBwjJuZGPrRT8e', '2', '0', '2024-08-23 15:37:31', '2024-08-23 15:37:46');
INSERT INTO `login_record` VALUES (44, 2, 1, '2024-08-23 15:37:46', 'E8UAv9ZcwVprumnJB6v9eRcwGJpiJEUe', '2', '0', '2024-08-23 15:37:46', '2024-08-23 15:38:20');
INSERT INTO `login_record` VALUES (45, 2, 1, '2024-08-23 15:38:20', '9UIkq0q3l9LmPMiyuA0jkKYqCDW6pxOu', '2', '0', '2024-08-23 15:38:20', '2024-08-23 15:39:45');
INSERT INTO `login_record` VALUES (46, 2, 1, '2024-08-23 15:39:45', 'Vz3MIMqInAU7XorbmoaWFEb71SkgY62M', '2', '0', '2024-08-23 15:39:45', '2024-08-23 15:40:58');
INSERT INTO `login_record` VALUES (47, 2, 1, '2024-08-23 15:40:58', 'jui62RtdAgzFQxpxQR20R2pjtRx5Ed27', '2', '0', '2024-08-23 15:40:58', '2024-08-23 16:14:11');
INSERT INTO `login_record` VALUES (48, 2, 1, '2024-08-23 16:14:11', 'W94YL5CYqWdy8nJmGqSzYjjQNTt5RRCJ', '2', '0', '2024-08-23 16:14:11', '2024-08-23 16:15:36');
INSERT INTO `login_record` VALUES (49, 2, 1, '2024-08-23 16:15:36', 'eJZMpqABjTRCTpeAVxmoN58wBTbrPHFq', '2', '0', '2024-08-23 16:15:36', '2024-08-23 16:49:06');
INSERT INTO `login_record` VALUES (50, 2, 1, '2024-08-23 16:49:06', 'jgcPzjfVSRUpPNP0Va1tci7qax5TnTfc', '2', '0', '2024-08-23 16:49:06', '2024-08-23 16:49:36');
INSERT INTO `login_record` VALUES (51, 2, 1, '2024-08-23 16:49:36', 'I4Ved59rfWNUMIyNwV3IhZXuWCCI7TI0', '2', '0', '2024-08-23 16:49:36', '2024-08-23 17:03:46');
INSERT INTO `login_record` VALUES (52, 2, 1, '2024-08-23 17:03:46', 'sxSPF5BFWUj6FzRzkblL3cqmwXQLqg6u', '2', '0', '2024-08-23 17:03:46', '2024-08-23 17:05:56');
INSERT INTO `login_record` VALUES (53, 2, 1, '2024-08-23 17:05:56', 'y5o56Y9tN89CPcKnf0NiG624JuDmU4Is', '2', '0', '2024-08-23 17:05:56', '2024-08-23 17:21:45');
INSERT INTO `login_record` VALUES (54, 2, 1, '2024-08-23 17:21:45', 'TLI4ykg6yjAsNfcJBRPv7Zi4Bjxly8ez', '2', '0', '2024-08-23 17:21:45', '2024-08-30 15:31:15');
INSERT INTO `login_record` VALUES (55, 2, 1, '2024-08-30 15:31:15', 'DQNiakNC9Yj55rDWlFICHgO51K0dfL6x', '0', '0', '2024-08-30 15:31:15', '2024-08-30 15:31:15');

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限名称，界面展示，建议与界面导航一致',
  `alias` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限别名，英文，规范如下：sys，sysAccount sysAccountAdd',
  `level` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限等级 1分组（一级导航）2模块（页面）3功能（按钮）第四级路由不在本表中体现',
  `pid` bigint(0) NOT NULL DEFAULT 0 COMMENT '父级ID，默认为1',
  `sort` int(0) NOT NULL DEFAULT 0 COMMENT '字典排序',
  `static` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '默认启用权限，0 启用 1 不启，启用后，该权限默认被分配，不可去勾',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `alias_uni`(`alias`) USING BTREE COMMENT '权限别名唯一',
  UNIQUE INDEX `name_uni`(`name`) USING BTREE COMMENT '权限名唯一',
  INDEX `permission_pid_query`(`pid`, `sort`) USING BTREE COMMENT '权限父级ID查询索引'
) ENGINE = InnoDB AUTO_INCREMENT = 52 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, '根权限', 'ROOT', '0', 0, 1, '0', '1', '0', NULL, NULL);
INSERT INTO `permission` VALUES (2, '通用接口', 'Open', '1', 1, 0, '0', '1', '0', NULL, '2024-08-23 14:07:20');
INSERT INTO `permission` VALUES (3, '主页', 'Center', '2', 1, 0, '0', '1', '0', NULL, '2024-07-31 08:39:09');
INSERT INTO `permission` VALUES (4, '平台管理', 'Plat', '1', 1, 1, '1', '1', '0', NULL, NULL);
INSERT INTO `permission` VALUES (7, '角色', 'PlatRole', '2', 4, 2, '1', '0', '0', '2024-05-28 08:38:24', '2024-07-31 21:23:30');
INSERT INTO `permission` VALUES (9, '账号', 'PlatAccount', '2', 4, 0, '1', '0', '0', '2024-05-28 08:50:32', '2024-07-31 14:57:57');
INSERT INTO `permission` VALUES (10, '权限', 'PlatPermission', '2', 4, 3, '1', '0', '0', '2024-05-28 08:50:44', '2024-07-31 09:01:57');
INSERT INTO `permission` VALUES (11, '工作中心', 'CenterIndex', '3', 3, 0, '1', '0', '0', '2024-06-13 16:17:38', '2024-07-31 08:37:39');
INSERT INTO `permission` VALUES (12, '部门', 'PlatDept', '2', 4, 1, '1', '0', '0', '2024-06-21 18:24:48', '2024-07-31 08:54:46');
INSERT INTO `permission` VALUES (13, '账号新建', 'PlatAccountAdd', '3', 9, 1, '1', '0', '0', '2024-07-31 08:47:53', '2024-07-31 21:27:54');
INSERT INTO `permission` VALUES (14, '账号编辑', 'PlatAccountEdit', '3', 9, 2, '1', '0', '0', '2024-07-31 08:49:01', '2024-07-31 08:49:01');
INSERT INTO `permission` VALUES (15, '账号删除', 'PlatAccountDel', '3', 9, 3, '1', '0', '0', '2024-07-31 08:49:24', '2024-07-31 08:49:24');
INSERT INTO `permission` VALUES (16, '密码重置', 'PlatAccountReset', '3', 9, 4, '1', '0', '0', '2024-07-31 08:49:45', '2024-07-31 08:49:45');
INSERT INTO `permission` VALUES (17, '部门新建', 'PlatDeptAdd', '3', 12, 1, '1', '0', '0', '2024-07-31 08:55:56', '2024-07-31 08:55:56');
INSERT INTO `permission` VALUES (18, '部门编辑', 'PlatDeptEdit', '3', 12, 2, '1', '0', '0', '2024-07-31 08:56:18', '2024-07-31 08:56:18');
INSERT INTO `permission` VALUES (19, '部门删除', 'PlatDeptDel', '3', 12, 3, '1', '0', '0', '2024-07-31 08:56:43', '2024-07-31 08:56:43');
INSERT INTO `permission` VALUES (20, '部门排序', 'PlatDeptSort', '3', 12, 5, '1', '0', '0', '2024-07-31 08:57:15', '2024-07-31 09:04:10');
INSERT INTO `permission` VALUES (21, '部门迁移', 'PlatDeptMerge', '3', 12, 4, '1', '0', '0', '2024-07-31 08:57:46', '2024-07-31 08:57:46');
INSERT INTO `permission` VALUES (22, '角色新建', 'PlatRoleAdd', '3', 7, 1, '1', '0', '0', '2024-07-31 08:59:57', '2024-07-31 08:59:57');
INSERT INTO `permission` VALUES (23, '角色编辑', 'PlatRoleEdit', '3', 7, 2, '1', '0', '0', '2024-07-31 09:00:14', '2024-07-31 09:00:14');
INSERT INTO `permission` VALUES (24, '角色删除', 'PlatRoleDel', '3', 7, 3, '1', '0', '0', '2024-07-31 09:00:38', '2024-07-31 09:00:38');
INSERT INTO `permission` VALUES (25, '权限新建', 'PlatPermissionAdd', '3', 10, 1, '1', '0', '0', '2024-07-31 09:02:36', '2024-07-31 09:02:46');
INSERT INTO `permission` VALUES (26, '权限编辑', 'PlatPermissionEdit', '3', 10, 2, '1', '0', '0', '2024-07-31 09:03:13', '2024-07-31 09:03:13');
INSERT INTO `permission` VALUES (27, '权限删除', 'PlatPermissionDel', '3', 10, 3, '1', '0', '0', '2024-07-31 09:03:36', '2024-07-31 09:03:36');
INSERT INTO `permission` VALUES (28, '权限排序', 'PlatPermissionSort', '3', 10, 4, '1', '0', '0', '2024-07-31 09:04:02', '2024-07-31 09:04:02');
INSERT INTO `permission` VALUES (29, '路由', 'PlatRouter', '2', 4, 4, '1', '0', '0', '2024-07-31 09:05:03', '2024-07-31 15:03:18');
INSERT INTO `permission` VALUES (30, '路由新建', 'PlatRouterAdd', '3', 29, 1, '1', '0', '0', '2024-07-31 09:05:31', '2024-07-31 09:05:31');
INSERT INTO `permission` VALUES (31, '路由编辑', 'PlatRouterEdit', '3', 29, 2, '1', '0', '0', '2024-07-31 09:05:50', '2024-07-31 09:05:50');
INSERT INTO `permission` VALUES (32, '路由删除', 'PlatRouterDel', '3', 29, 3, '1', '0', '0', '2024-07-31 09:06:12', '2024-07-31 09:06:12');
INSERT INTO `permission` VALUES (33, '响应', 'PlatResponse', '2', 4, 5, '1', '0', '0', '2024-07-31 09:06:57', '2024-07-31 09:06:57');
INSERT INTO `permission` VALUES (34, '字典', 'PlatDict', '2', 4, 6, '1', '0', '0', '2024-07-31 09:07:38', '2024-07-31 09:07:38');
INSERT INTO `permission` VALUES (35, '配置', 'PlatConfig', '2', 4, 7, '1', '0', '0', '2024-07-31 09:08:13', '2024-08-23 16:14:41');
INSERT INTO `permission` VALUES (36, '响应码新建', 'PlatResponseAdd', '3', 33, 1, '1', '0', '0', '2024-07-31 09:09:09', '2024-07-31 09:09:09');
INSERT INTO `permission` VALUES (37, '响应码编辑', 'PlatResponseEdit', '3', 33, 2, '1', '0', '0', '2024-07-31 09:09:31', '2024-07-31 21:18:39');
INSERT INTO `permission` VALUES (38, '响应码删除', 'PlatResponseDel', '3', 33, 3, '1', '0', '0', '2024-07-31 09:09:49', '2024-07-31 09:09:49');
INSERT INTO `permission` VALUES (39, '字典新建', 'PlatDictAdd', '3', 34, 1, '1', '0', '0', '2024-07-31 09:10:37', '2024-07-31 09:10:37');
INSERT INTO `permission` VALUES (40, '字典编辑', 'PlatDictEdit', '3', 34, 2, '1', '0', '0', '2024-07-31 09:10:59', '2024-07-31 09:10:59');
INSERT INTO `permission` VALUES (41, '字典删除', 'PlatDictDel', '3', 34, 3, '1', '0', '0', '2024-07-31 09:11:18', '2024-07-31 09:11:18');
INSERT INTO `permission` VALUES (42, '字典排序', 'PlatDictSort', '3', 34, 4, '1', '0', '0', '2024-07-31 09:11:43', '2024-07-31 09:11:43');
INSERT INTO `permission` VALUES (43, '配置编辑', 'PlatConfigEdit', '3', 35, 1, '1', '0', '0', '2024-07-31 09:12:12', '2024-07-31 09:12:12');
INSERT INTO `permission` VALUES (44, '账号详情', 'PlatAccountView', '3', 9, 0, '1', '0', '0', '2024-07-31 14:55:20', '2024-07-31 14:55:20');
INSERT INTO `permission` VALUES (45, '部门详情', 'PlatDeptView', '3', 12, 0, '1', '0', '0', '2024-07-31 14:58:40', '2024-07-31 14:58:40');
INSERT INTO `permission` VALUES (46, '角色详情', 'PlatRoleView', '3', 7, 0, '1', '0', '0', '2024-07-31 14:59:29', '2024-07-31 14:59:29');
INSERT INTO `permission` VALUES (47, '权限详情', 'PlatPermissionView', '3', 10, 0, '1', '0', '0', '2024-07-31 15:01:36', '2024-07-31 15:01:36');
INSERT INTO `permission` VALUES (48, '路由详情', 'PlatRouterView', '3', 29, 0, '1', '0', '0', '2024-07-31 15:02:12', '2024-07-31 15:02:12');
INSERT INTO `permission` VALUES (49, '响应码详情', 'PlatResponseView', '3', 33, 0, '1', '0', '0', '2024-07-31 15:02:51', '2024-07-31 15:02:51');
INSERT INTO `permission` VALUES (50, '字典详情', 'PlatDictView', '3', 34, 0, '1', '0', '0', '2024-07-31 15:03:55', '2024-07-31 15:03:55');
INSERT INTO `permission` VALUES (52, '配置详情', 'PlatConfigView', '3', 35, 0, '1', '0', '0', '2024-07-31 21:24:18', '2024-07-31 21:24:18');

-- ----------------------------
-- Table structure for permission_router
-- ----------------------------
DROP TABLE IF EXISTS `permission_router`;
CREATE TABLE `permission_router`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `permission_id` bigint(0) NOT NULL COMMENT '权限ID',
  `router_id` bigint(0) NOT NULL COMMENT '路由ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `permission_router_uni`(`permission_id`, `router_id`) USING BTREE COMMENT '权限和路由组合唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 94 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限路由' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission_router
-- ----------------------------
INSERT INTO `permission_router` VALUES (92, 2, 2);
INSERT INTO `permission_router` VALUES (93, 2, 54);
INSERT INTO `permission_router` VALUES (88, 7, 45);
INSERT INTO `permission_router` VALUES (72, 9, 8);
INSERT INTO `permission_router` VALUES (73, 9, 17);
INSERT INTO `permission_router` VALUES (74, 9, 25);
INSERT INTO `permission_router` VALUES (75, 9, 44);
INSERT INTO `permission_router` VALUES (36, 10, 25);
INSERT INTO `permission_router` VALUES (38, 10, 31);
INSERT INTO `permission_router` VALUES (37, 10, 33);
INSERT INTO `permission_router` VALUES (22, 12, 14);
INSERT INTO `permission_router` VALUES (23, 12, 17);
INSERT INTO `permission_router` VALUES (24, 12, 25);
INSERT INTO `permission_router` VALUES (90, 13, 4);
INSERT INTO `permission_router` VALUES (16, 14, 6);
INSERT INTO `permission_router` VALUES (17, 15, 5);
INSERT INTO `permission_router` VALUES (18, 16, 9);
INSERT INTO `permission_router` VALUES (25, 17, 10);
INSERT INTO `permission_router` VALUES (26, 18, 13);
INSERT INTO `permission_router` VALUES (27, 19, 12);
INSERT INTO `permission_router` VALUES (45, 20, 11);
INSERT INTO `permission_router` VALUES (46, 20, 15);
INSERT INTO `permission_router` VALUES (30, 21, 16);
INSERT INTO `permission_router` VALUES (33, 22, 40);
INSERT INTO `permission_router` VALUES (34, 23, 42);
INSERT INTO `permission_router` VALUES (35, 24, 41);
INSERT INTO `permission_router` VALUES (40, 25, 27);
INSERT INTO `permission_router` VALUES (41, 26, 30);
INSERT INTO `permission_router` VALUES (42, 27, 29);
INSERT INTO `permission_router` VALUES (44, 28, 28);
INSERT INTO `permission_router` VALUES (43, 28, 32);
INSERT INTO `permission_router` VALUES (82, 29, 25);
INSERT INTO `permission_router` VALUES (83, 29, 50);
INSERT INTO `permission_router` VALUES (50, 30, 46);
INSERT INTO `permission_router` VALUES (51, 31, 48);
INSERT INTO `permission_router` VALUES (52, 32, 47);
INSERT INTO `permission_router` VALUES (53, 33, 25);
INSERT INTO `permission_router` VALUES (54, 33, 39);
INSERT INTO `permission_router` VALUES (57, 34, 22);
INSERT INTO `permission_router` VALUES (56, 34, 24);
INSERT INTO `permission_router` VALUES (55, 34, 25);
INSERT INTO `permission_router` VALUES (94, 35, 25);
INSERT INTO `permission_router` VALUES (61, 36, 34);
INSERT INTO `permission_router` VALUES (60, 36, 38);
INSERT INTO `permission_router` VALUES (85, 37, 36);
INSERT INTO `permission_router` VALUES (63, 38, 35);
INSERT INTO `permission_router` VALUES (65, 39, 18);
INSERT INTO `permission_router` VALUES (64, 39, 23);
INSERT INTO `permission_router` VALUES (66, 40, 21);
INSERT INTO `permission_router` VALUES (67, 41, 20);
INSERT INTO `permission_router` VALUES (68, 42, 19);
INSERT INTO `permission_router` VALUES (69, 42, 26);
INSERT INTO `permission_router` VALUES (70, 43, 51);
INSERT INTO `permission_router` VALUES (71, 44, 7);
INSERT INTO `permission_router` VALUES (76, 45, 14);
INSERT INTO `permission_router` VALUES (77, 46, 43);
INSERT INTO `permission_router` VALUES (79, 47, 31);
INSERT INTO `permission_router` VALUES (80, 48, 49);
INSERT INTO `permission_router` VALUES (81, 49, 37);
INSERT INTO `permission_router` VALUES (84, 50, 22);
INSERT INTO `permission_router` VALUES (89, 52, 52);

-- ----------------------------
-- Table structure for response_code
-- ----------------------------
DROP TABLE IF EXISTS `response_code`;
CREATE TABLE `response_code`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '响应码',
  `service_code` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务ID，来源于字典，指定响应码归属业务',
  `type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '响应类型，该字段用于筛选，可配置2和5',
  `zh_cn` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '中文响应文言',
  `en_us` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '英文响应文言',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '其他备注信息',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code_uni`(`code`) USING BTREE COMMENT '响应码唯一',
  INDEX `query_filter`(`service_code`, `type`) USING BTREE COMMENT '平台筛选查询索引'
) ENGINE = InnoDB AUTO_INCREMENT = 77 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '响应码配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of response_code
-- ----------------------------
INSERT INTO `response_code` VALUES (1, 'E000', '0', 'E', '系统异常', 'System exception', '系统异常（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (2, 'E001', '0', 'E', '参数非法', 'Illegal parameters', '参数非法（默认）（免翻译）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (3, 'E002', '0', 'E', '尚未登陆', 'Not logged in yet', '尚未登陆（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (4, 'E003', '0', 'E', '无权访问', 'No access rights', '无权访问（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (5, 'E004', '0', 'E', '路径不存在', 'Path does not exist', '路径不存在（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (6, 'S000', '0', 'S', '处理成功', 'Processed successfully', '处理成功（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (7, 'S001', '0', 'S', '登陆成功', 'Landed successfully', '登陆成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (8, 'S002', '0', 'S', '密码重置成功', 'Password reset successful', '密码重置成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (9, 'F000', '0', 'F', '处理失败', 'Processing failed', '处理失败（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (10, 'F001', '0', 'F', '登陆失败，请联系管理员', 'Login failed, please contact the administrator', '登陆失败，请联系管理员', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (11, 'F002', '0', 'F', '异常登陆，请联系管理员', 'Abnormal login, please contact the administrator', '异常登陆，请联系管理员', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (12, 'F003', '0', 'F', '密码重置失败，请联系管理员', 'Password reset failed, please contact the administrator', '密码重置失败，请联系管理员', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (13, 'S100', '1', 'S', '字典创建成功', 'Dictionary creation successful', '字典创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (14, 'S101', '1', 'S', '字典编辑成功', 'Dictionary editing successful', '字典编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (15, 'S102', '1', 'S', '字典排序成功', 'Dictionary sorting successful', '字典排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (16, 'S103', '1', 'S', '字典封存成功', 'Dictionary sealing successful', '字典封存成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (17, 'F100', '1', 'F', '字典查询失败', 'Dictionary query failed', '字典查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (18, 'F101', '1', 'F', '字典分组下字典值唯一', 'Dictionary value is unique under dictionary group', '字典分组下字典值唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (19, 'F102', '1', 'F', '内置字典禁止刪除', 'Built-in dictionary cannot be deleted', '内置字典禁止刪除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (20, 'F103', '1', 'F', '字典排序失败', 'Dictionary sorting failed', '字典排序失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (21, 'S200', '2', 'S', '响应码创建成功', 'Response code creation successful', '响应码创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (22, 'S201', '2', 'S', '响应码创建成功,实际响应码为{{code}}', 'Response code creation successful, actual response code is {{code}}', '响应码创建成功,实际响应码为{{code}}', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (23, 'S202', '2', 'S', '响应码编辑成功', 'Response code editing successful', '响应码编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (24, 'S203', '2', 'S', '响应码封存成功', 'Response code sealing successful', '响应码封存成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (25, 'F200', '2', 'F', '响应码查询失败', 'Response code query failed', '响应码查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (26, 'F201', '2', 'F', '响应码全局唯一', 'Response code is globally unique', '响应码全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (27, 'F202', '2', 'F', '内置响应码禁止删除', 'Built-in response code cannot be deleted', '内置响应码禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (28, 'S300', '3', 'S', '路由创建成功', 'Route creation successful', '路由创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (29, 'S301', '3', 'S', '路由编辑成功', 'Route editing successful', '路由编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (30, 'S302', '3', 'S', '路由删除成功', 'Route deletion successful', '路由删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (31, 'F300', '3', 'F', '路由查询失败', 'Route query failed', '路由查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (32, 'F301', '3', 'F', '路由地址全局唯一', 'Route address is globally unique', '路由地址全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (33, 'F302', '3', 'F', '路由名称全局唯一', 'Route name is globally unique', '路由名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (34, 'F303', '3', 'F', '内置路由禁止删除', 'Built-in route cannot be deleted', '内置路由禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (35, 'F304', '3', 'F', '路由删除', 'Route deletion', '路由删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (36, 'S400', '4', 'S', '权限创建成功', 'Permission creation successful', '权限创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (37, 'S401', '4', 'S', '权限编辑成功', 'Permission editing successful', '权限编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (38, 'S402', '4', 'S', '权限删除成功', 'Permission deletion successful', '权限删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (39, 'S403', '4', 'S', '权限排序成功', 'Permissions sorted successfully', '权限排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (40, 'F400', '4', 'F', '权限查询失败', 'Permission query failed', '权限查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (41, 'F401', '4', 'F', '权限别名全局唯一', 'Permission alias is globally unique', '权限别名全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (42, 'F402', '4', 'F', '权限名称全局唯一', 'Permission name is globally unique', '权限名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (43, 'F403', '4', 'F', '内置权限禁止删除', 'Built-in permission cannot be deleted', '内置权限禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (44, 'F404', '4', 'F', '权限配置路由同步失败', 'Permission configuration route synchronization failed', '权限配置路由同步失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (45, 'F405', '4', 'F', '权限配置删除失败', 'Permission configuration deletion failed', '权限配置删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (46, 'S500', '5', 'S', '角色创建成功', 'Role creation successful', '角色创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (47, 'S501', '5', 'S', '角色编辑成功', 'Role editing successful', '角色编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (48, 'S502', '5', 'S', '角色删除失败', 'Role deletion failed', '角色删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (49, 'F500', '5', 'F', '角色查询失败', 'Role query failed', '角色查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (50, 'F501', '5', 'F', '内置角色禁止编辑', 'Built-in roles cannot be edited', '内置角色禁止编辑', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (51, 'F502', '5', 'F', '角色名全局唯一', 'Role name is globally unique', '角色名全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (52, 'F503', '5', 'F', '角色权限配置失败', 'Role permission configuration failed', '角色权限配置失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (53, 'F504', '5', 'F', '角色删除失败', 'Role deletion failed', '角色删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (54, 'S600', '6', 'S', '部门创建成功', 'Department created successfully', '集团部门创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (55, 'S601', '6', 'S', '部门编辑成功', 'Department editing successful', '集团部门编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (56, 'S602', '6', 'S', '部门删除成功', 'Department deleted successfully', '集团部门删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (57, 'S603', '6', 'S', '部门排序成功', 'Department sorting successful', '集团部门排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (58, 'S604', '6', 'S', '部门迁移成功', 'Department migration successful', '集团部门迁移成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (59, 'F600', '6', 'F', '部门查询失败', 'Department query failed', '集团部门查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (60, 'F601', '6', 'F', '部门名称全局唯一', 'Department name Globally unique', '集团部门名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (61, 'F602', '6', 'F', '内置部门禁止删除', 'Deletion of built-in departments is prohibited', '内置集团部门禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (62, 'F603', '6', 'F', '部门删除失败', 'Department deletion failed', '集团部门删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (63, 'F604', '6', 'F', '部门存在子部门禁止删除', 'Department has sub-departments and is prohibited from deletion', '集团部门存在子部门禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (64, 'F605', '6', 'F', '部门存在成员禁止删除', 'Department members cannot be deleted if they exist', '集团部门存在成员禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (65, 'F606', '6', 'F', '部门迁移失败', 'Department migration failed', '集团部门迁移失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (66, 'S700', '7', 'S', '登陆账号创建成功', 'Login account created successfully', '登陆账号创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (67, 'S701', '7', 'S', '登陆账号编辑成功', 'Login account edited successfully', '登陆账号编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (68, 'S702', '7', 'S', '登陆账号删除成功', 'Login account deleted successfully', '登陆账号删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (69, 'S703', '7', 'S', '登陆账号重置成功', 'Login account reset successfully', '登陆账号重置成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (70, 'F700', '7', 'F', '登陆账号查询失败', 'Login account query failed', '登陆账号查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (71, 'F701', '7', 'F', '登陆账号全局唯一', 'Login account is globally unique', '登陆账号全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (72, 'F702', '7', 'F', '账号角色同步失败', 'Account role synchronization failed', '账号角色同步失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (73, 'F703', '7', 'F', '特殊账号禁止编辑', 'Editing of special accounts is prohibited', '特殊账号禁止编辑', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (74, 'F704', '7', 'F', '登陆账号删除失败', 'Login account deletion failed', '登陆账号删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (75, 'F705', '7', 'F', '登陆账号重置失败', 'Login account reset failed', '登陆账号重置失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (76, 'S800', '8', 'S', '系统配置编辑成功', 'System configuration editing successful', '系统配置编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (77, 'F800', '8', 'F', '系统配置查询失败', 'System configuration query failed', '系统配置查询失败', '1', '0', NULL, NULL);

-- ----------------------------
-- Table structure for response_code_copy1
-- ----------------------------
DROP TABLE IF EXISTS `response_code_copy1`;
CREATE TABLE `response_code_copy1`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '响应码',
  `service_code` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务ID，来源于字典，指定响应码归属业务',
  `type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '响应类型，该字段用于筛选，可配置2和5',
  `zh_cn` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '中文响应文言',
  `en_us` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '英文响应文言',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '其他备注信息',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code_uni`(`code`) USING BTREE COMMENT '响应码唯一',
  INDEX `query_filter`(`service_code`, `type`) USING BTREE COMMENT '平台筛选查询索引'
) ENGINE = InnoDB AUTO_INCREMENT = 72 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '响应码配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of response_code_copy1
-- ----------------------------
INSERT INTO `response_code_copy1` VALUES (1, 'S000', '0', 'S', '处理成功', 'Successful processing', '处理成功（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (2, 'F000', '0', 'F', '处理失败', 'Failed processing', '处理失败（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (3, 'E000', '0', 'E', '系统异常', 'System exception', '系统异常（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (4, 'E001', '0', 'E', '参数非法', 'Illegal parameters', '参数非法（默认）（免翻译）', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (5, 'E002', '0', 'E', '尚未登陆', 'Not logged in yet', '尚未登陆（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (6, 'E003', '0', 'E', '无权访问', 'No access rights', '无权访问（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (7, 'E004', '0', 'E', '路径不存在', 'Path does not exist', '路径不存在（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (8, 'S100', '1', 'S', '字典创建成功', 'Dictionary creation successful', '字典创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (9, 'S101', '1', 'S', '字典编辑成功', 'Dictionary editing successful', '字典编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (10, 'S102', '1', 'S', '字典排序成功', 'Dictionary sorting successful', '字典排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (11, 'S103', '1', 'S', '字典封存成功', 'Dictionary sealing successful', '字典封存成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (12, 'F100', '1', 'F', '字典查询失败', 'Dictionary query failed', '字典查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (13, 'F101', '1', 'F', '字典分组下字典值唯一', 'Dictionary value is unique under dictionary group', '字典分组下字典值唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (14, 'F102', '1', 'F', '内置字典禁止刪除', 'Built-in dictionary cannot be deleted', '内置字典禁止刪除', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (15, 'F103', '1', 'F', '字典排序失败', 'Dictionary sorting failed', '字典排序失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (16, 'S200', '2', 'S', '响应码创建成功', 'Response code creation successful', '响应码创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (17, 'S201', '2', 'S', '响应码创建成功,实际响应码为{{code}}', 'Response code creation successful, actual response code is {{code}}', '响应码创建成功,实际响应码为{{code}}', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (18, 'S202', '2', 'S', '响应码编辑成功', 'Response code editing successful', '响应码编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (19, 'S203', '2', 'S', '响应码封存成功', 'Response code sealing successful', '响应码封存成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (20, 'F200', '2', 'F', '响应码查询失败', 'Response code query failed', '响应码查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (21, 'F201', '2', 'F', '响应码全局唯一', 'Response code is globally unique', '响应码全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (22, 'F202', '2', 'F', '内置响应码禁止删除', 'Built-in response code cannot be deleted', '内置响应码禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (23, 'S300', '3', 'S', '路由创建成功', 'Route creation successful', '路由创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (24, 'S301', '3', 'S', '路由编辑成功', 'Route editing successful', '路由编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (25, 'S302', '3', 'S', '路由删除成功', 'Route deletion successful', '路由删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (26, 'F300', '3', 'F', '路由查询失败', 'Route query failed', '路由查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (27, 'F301', '3', 'F', '路由地址全局唯一', 'Route address is globally unique', '路由地址全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (28, 'F302', '3', 'F', '路由名称全局唯一', 'Route name is globally unique', '路由名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (29, 'F303', '3', 'F', '内置路由禁止删除', 'Built-in route cannot be deleted', '内置路由禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (30, 'F304', '3', 'F', '路由删除', 'Route deletion', '路由删除', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (31, 'S400', '4', 'S', '权限创建成功', 'Permission creation successful', '权限创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (32, 'S401', '4', 'S', '权限编辑成功', 'Permission editing successful', '权限编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (33, 'S402', '4', 'S', '权限删除成功', 'Permission deletion successful', '权限删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (34, 'S403', '4', 'S', '权限排序成功', 'Permissions sorted successfully', '权限排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (35, 'F400', '4', 'F', '权限查询失败', 'Permission query failed', '权限查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (36, 'F401', '4', 'F', '权限别名全局唯一', 'Permission alias is globally unique', '权限别名全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (37, 'F402', '4', 'F', '权限名称全局唯一', 'Permission name is globally unique', '权限名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (38, 'F403', '4', 'F', '内置权限禁止删除', 'Built-in permission cannot be deleted', '内置权限禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (39, 'F404', '4', 'F', '权限配置路由同步失败', 'Permission configuration route synchronization failed', '权限配置路由同步失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (40, 'F405', '4', 'F', '权限配置删除失败', 'Permission configuration deletion failed', '权限配置删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (41, 'S500', '5', 'S', '角色创建成功', 'Role creation successful', '角色创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (42, 'S501', '5', 'S', '角色编辑成功', 'Role editing successful', '角色编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (43, 'S502', '5', 'S', '角色删除失败', 'Role deletion failed', '角色删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (44, 'F500', '5', 'F', '角色查询失败', 'Role query failed', '角色查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (45, 'F501', '5', 'F', '内置角色禁止编辑', 'Built-in roles cannot be edited', '内置角色禁止编辑', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (46, 'F502', '5', 'F', '角色名全局唯一', 'Role name is globally unique', '角色名全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (47, 'F503', '5', 'F', '角色权限配置失败', 'Role permission configuration failed', '角色权限配置失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (48, 'F504', '5', 'F', '角色删除失败', 'Role deletion failed', '角色删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (49, 'S600', '6', 'S', '部门创建成功', 'Department created successfully', '集团部门创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (50, 'S601', '6', 'S', '部门编辑成功', 'Department editing successful', '集团部门编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (51, 'S602', '6', 'S', '部门删除成功', 'Department deleted successfully', '集团部门删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (52, 'S603', '6', 'S', '部门排序成功', 'Department sorting successful', '集团部门排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (53, 'S604', '6', 'S', '部门迁移成功', 'Department migration successful', '集团部门迁移成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (54, 'F600', '6', 'F', '部门查询失败', 'Department query failed', '集团部门查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (55, 'F601', '6', 'F', '部门名称全局唯一', 'Department name Globally unique', '集团部门名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (56, 'F602', '6', 'F', '内置部门禁止删除', 'Deletion of built-in departments is prohibited', '内置集团部门禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (57, 'F603', '6', 'F', '部门删除失败', 'Department deletion failed', '集团部门删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (58, 'F604', '6', 'F', '部门存在子部门禁止删除', 'Department has sub-departments and is prohibited from deletion', '集团部门存在子部门禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (59, 'F605', '6', 'F', '部门存在成员禁止删除', 'Department members cannot be deleted if they exist', '集团部门存在成员禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (60, 'F606', '6', 'F', '部门迁移失败', 'Department migration failed', '集团部门迁移失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (61, 'S700', '7', 'S', '登陆账号创建成功', 'Login account created successfully', '登陆账号创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (62, 'S701', '7', 'S', '登陆账号编辑成功', 'Login account edited successfully', '登陆账号编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (63, 'S702', '7', 'S', '登陆账号删除成功', 'Login account deleted successfully', '登陆账号删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (64, 'S703', '7', 'S', '登陆账号重置成功', 'Login account reset successfully', '登陆账号重置成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (65, 'F700', '7', 'F', '登陆账号查询失败', 'Login account query failed', '登陆账号查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (66, 'F701', '7', 'F', '登陆账号全局唯一', 'Login account is globally unique', '登陆账号全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (67, 'F702', '7', 'F', '账号角色同步失败', 'Account role synchronization failed', '账号角色同步失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (68, 'F703', '7', 'F', '特殊账号禁止编辑', 'Editing of special accounts is prohibited', '特殊账号禁止编辑', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (69, 'F704', '7', 'F', '登陆账号删除失败', 'Login account deletion failed', '登陆账号删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (70, 'F705', '7', 'F', '登陆账号重置失败', 'Login account reset failed', '登陆账号重置失败', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (71, 'S800', '8', 'S', '系统配置编辑成功', 'System configuration editing successful', '系统配置编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code_copy1` VALUES (72, 'F800', '8', 'F', '系统配置查询失败', 'System configuration query failed', '系统配置查询失败', '1', '0', NULL, NULL);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `remark` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色备注',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name_uni`(`name`) USING BTREE COMMENT '角色名称唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (5, '管理员', '', '0', '0', '2024-06-29 13:03:06', '2024-08-23 16:49:17');
INSERT INTO `role` VALUES (6, 'Test', '', '0', '0', '2024-07-31 21:22:03', '2024-07-31 21:22:03');

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `role_id` bigint(0) NOT NULL COMMENT '角色ID',
  `permission_id` bigint(0) NOT NULL COMMENT '权限ID',
  `check_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '选择类型 0 check 1 halfCheck',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `role_permission_uni`(`role_id`, `permission_id`) USING BTREE COMMENT '角色权限全局唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 952 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO `role_permission` VALUES (84, 6, 2, '0');
INSERT INTO `role_permission` VALUES (85, 6, 3, '0');
INSERT INTO `role_permission` VALUES (86, 6, 44, '0');
INSERT INTO `role_permission` VALUES (87, 6, 15, '0');
INSERT INTO `role_permission` VALUES (88, 6, 16, '0');
INSERT INTO `role_permission` VALUES (89, 6, 29, '0');
INSERT INTO `role_permission` VALUES (90, 6, 35, '0');
INSERT INTO `role_permission` VALUES (91, 6, 45, '0');
INSERT INTO `role_permission` VALUES (92, 6, 17, '0');
INSERT INTO `role_permission` VALUES (93, 6, 18, '0');
INSERT INTO `role_permission` VALUES (94, 6, 20, '0');
INSERT INTO `role_permission` VALUES (95, 6, 47, '0');
INSERT INTO `role_permission` VALUES (96, 6, 27, '0');
INSERT INTO `role_permission` VALUES (97, 6, 28, '0');
INSERT INTO `role_permission` VALUES (98, 6, 49, '0');
INSERT INTO `role_permission` VALUES (99, 6, 36, '0');
INSERT INTO `role_permission` VALUES (100, 6, 50, '0');
INSERT INTO `role_permission` VALUES (101, 6, 39, '0');
INSERT INTO `role_permission` VALUES (102, 6, 42, '0');
INSERT INTO `role_permission` VALUES (103, 6, 46, '0');
INSERT INTO `role_permission` VALUES (104, 6, 24, '0');
INSERT INTO `role_permission` VALUES (105, 6, 11, '0');
INSERT INTO `role_permission` VALUES (106, 6, 48, '0');
INSERT INTO `role_permission` VALUES (107, 6, 30, '0');
INSERT INTO `role_permission` VALUES (108, 6, 31, '0');
INSERT INTO `role_permission` VALUES (109, 6, 32, '0');
INSERT INTO `role_permission` VALUES (110, 6, 43, '0');
INSERT INTO `role_permission` VALUES (111, 6, 1, '1');
INSERT INTO `role_permission` VALUES (112, 6, 9, '1');
INSERT INTO `role_permission` VALUES (113, 6, 4, '1');
INSERT INTO `role_permission` VALUES (114, 6, 12, '1');
INSERT INTO `role_permission` VALUES (115, 6, 10, '1');
INSERT INTO `role_permission` VALUES (116, 6, 33, '1');
INSERT INTO `role_permission` VALUES (117, 6, 34, '1');
INSERT INTO `role_permission` VALUES (118, 6, 7, '1');
INSERT INTO `role_permission` VALUES (905, 5, 2, '0');
INSERT INTO `role_permission` VALUES (906, 5, 3, '0');
INSERT INTO `role_permission` VALUES (907, 5, 7, '0');
INSERT INTO `role_permission` VALUES (908, 5, 10, '0');
INSERT INTO `role_permission` VALUES (909, 5, 12, '0');
INSERT INTO `role_permission` VALUES (910, 5, 13, '0');
INSERT INTO `role_permission` VALUES (911, 5, 14, '0');
INSERT INTO `role_permission` VALUES (912, 5, 15, '0');
INSERT INTO `role_permission` VALUES (913, 5, 16, '0');
INSERT INTO `role_permission` VALUES (914, 5, 29, '0');
INSERT INTO `role_permission` VALUES (915, 5, 33, '0');
INSERT INTO `role_permission` VALUES (916, 5, 34, '0');
INSERT INTO `role_permission` VALUES (917, 5, 35, '0');
INSERT INTO `role_permission` VALUES (918, 5, 11, '0');
INSERT INTO `role_permission` VALUES (919, 5, 46, '0');
INSERT INTO `role_permission` VALUES (920, 5, 22, '0');
INSERT INTO `role_permission` VALUES (921, 5, 23, '0');
INSERT INTO `role_permission` VALUES (922, 5, 24, '0');
INSERT INTO `role_permission` VALUES (923, 5, 47, '0');
INSERT INTO `role_permission` VALUES (924, 5, 25, '0');
INSERT INTO `role_permission` VALUES (925, 5, 26, '0');
INSERT INTO `role_permission` VALUES (926, 5, 27, '0');
INSERT INTO `role_permission` VALUES (927, 5, 28, '0');
INSERT INTO `role_permission` VALUES (928, 5, 45, '0');
INSERT INTO `role_permission` VALUES (929, 5, 17, '0');
INSERT INTO `role_permission` VALUES (930, 5, 18, '0');
INSERT INTO `role_permission` VALUES (931, 5, 19, '0');
INSERT INTO `role_permission` VALUES (932, 5, 21, '0');
INSERT INTO `role_permission` VALUES (933, 5, 20, '0');
INSERT INTO `role_permission` VALUES (934, 5, 48, '0');
INSERT INTO `role_permission` VALUES (935, 5, 30, '0');
INSERT INTO `role_permission` VALUES (936, 5, 31, '0');
INSERT INTO `role_permission` VALUES (937, 5, 32, '0');
INSERT INTO `role_permission` VALUES (938, 5, 49, '0');
INSERT INTO `role_permission` VALUES (939, 5, 36, '0');
INSERT INTO `role_permission` VALUES (940, 5, 37, '0');
INSERT INTO `role_permission` VALUES (941, 5, 38, '0');
INSERT INTO `role_permission` VALUES (942, 5, 50, '0');
INSERT INTO `role_permission` VALUES (943, 5, 39, '0');
INSERT INTO `role_permission` VALUES (944, 5, 40, '0');
INSERT INTO `role_permission` VALUES (945, 5, 41, '0');
INSERT INTO `role_permission` VALUES (946, 5, 42, '0');
INSERT INTO `role_permission` VALUES (947, 5, 52, '0');
INSERT INTO `role_permission` VALUES (948, 5, 43, '0');
INSERT INTO `role_permission` VALUES (949, 5, 44, '0');
INSERT INTO `role_permission` VALUES (950, 5, 9, '0');
INSERT INTO `role_permission` VALUES (951, 5, 4, '0');
INSERT INTO `role_permission` VALUES (952, 5, 1, '0');

-- ----------------------------
-- Table structure for router
-- ----------------------------
DROP TABLE IF EXISTS `router`;
CREATE TABLE `router`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由名称，用于界面展示，与权限关联',
  `url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由地址，后端访问URL，后端不在URL中携带参数，统一Post处理内容',
  `type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '免授权路由 0 授权 1 免授权',
  `service_code` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务编码（字典），为接口分组',
  `log_in_db` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '日志入库 0 入库 1 不入库',
  `req_log_print` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '请求日志打印 0 打印 1 不打印',
  `req_log_secure` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求日志脱敏字段，逗号分隔',
  `res_log_print` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '响应日志打印 0 打印 1 不打印',
  `res_log_secure` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '响应日志脱敏字段，逗号分隔',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `url_uni`(`url`) USING BTREE COMMENT '路由地址唯一',
  UNIQUE INDEX `name_uni`(`name`) USING BTREE COMMENT '路由名称唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口路由' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of router
-- ----------------------------
INSERT INTO `router` VALUES (1, '账号密码登陆', '/auth/login', '1', '0', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (2, '账号密码重置', '/auth/reset', '0', '0', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (3, '通用API示例', '/docs/sample', '0', '0', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (4, '登陆账号新建', '/plat/account/add', '0', '5', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (5, '登陆账号移除', '/plat/account/del', '0', '5', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (6, '登陆账号编辑', '/plat/account/edit', '0', '5', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (7, '登陆账号详情', '/plat/account/get', '0', '5', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:33:05');
INSERT INTO `router` VALUES (8, '登陆账号分页', '/plat/account/page', '0', '5', '1', '0', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (9, '登陆账号重置', '/plat/account/reset', '0', '5', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (10, '集团部门新建', '/plat/dept/add', '0', '7', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (11, '集团同级部门', '/plat/dept/bro', '0', '7', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (12, '集团部门移除', '/plat/dept/del', '0', '7', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (13, '集团部门编辑', '/plat/dept/edit', '0', '7', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (14, '集团部门详情', '/plat/dept/get', '0', '7', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:59');
INSERT INTO `router` VALUES (15, '集团部门排序', '/plat/dept/sort', '0', '7', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (16, '集团部门迁移', '/plat/dept/to', '0', '7', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (17, '集团部门树', '/plat/dept/tree', '0', '7', '1', '0', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (18, '字典新建', '/plat/dict/add', '0', '1', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (19, '字典同级查询', '/plat/dict/bro', '0', '1', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (20, '字典封存', '/plat/dict/del', '0', '1', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (21, '字典编辑', '/plat/dict/edit', '0', '1', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (22, '字典详情', '/plat/dict/get', '0', '1', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:44');
INSERT INTO `router` VALUES (23, '字典NextVal建议', '/plat/dict/nextVal', '0', '1', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (24, '字典分页', '/plat/dict/page', '0', '1', '1', '0', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (25, '读取字典', '/plat/dict/read', '0', '1', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:33:42');
INSERT INTO `router` VALUES (26, '排序处理', '/plat/dict/sort', '0', '1', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (27, '权限新建', '/plat/permission/add', '0', '4', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (28, '同级权限', '/plat/permission/bro', '0', '4', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (29, '权限封存', '/plat/permission/del', '0', '4', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (30, '权限编辑', '/plat/permission/edit', '0', '4', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (31, '权限详情', '/plat/permission/get', '0', '4', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:40');
INSERT INTO `router` VALUES (32, '权限排序', '/plat/permission/sort', '0', '4', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (33, '权限树', '/plat/permission/tree', '0', '4', '1', '0', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (34, '响应码新建', '/plat/response/add', '0', '2', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (35, '响应码封存', '/plat/response/del', '0', '2', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (36, '响应码编辑', '/plat/response/edit', '0', '2', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (37, '响应码详情', '/plat/response/get', '0', '2', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:34');
INSERT INTO `router` VALUES (38, '响应码NextVal建议', '/plat/response/nextVal', '0', '2', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (39, '响应码分页', '/plat/response/page', '0', '2', '1', '0', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (40, '角色新建', '/plat/role/add', '0', '6', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (41, '角色删除', '/plat/role/del', '0', '6', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (42, '角色编辑', '/plat/role/edit', '0', '6', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (43, '角色详情', '/plat/role/get', '0', '6', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:29');
INSERT INTO `router` VALUES (44, '角色列表', '/plat/role/list', '0', '6', '1', '0', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (45, '角色分页', '/plat/role/page', '0', '6', '1', '0', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (46, '路由接口新建', '/plat/router/add', '0', '3', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (47, '路由接口封存', '/plat/router/del', '0', '3', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (48, '路由接口编辑', '/plat/router/edit', '0', '3', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (49, '路由接口详情', '/plat/router/get', '0', '3', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:23');
INSERT INTO `router` VALUES (50, '路由接口分页', '/plat/router/page', '0', '3', '1', '0', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (51, '系统配置编辑', '/plat/sysConfig/edit', '0', '8', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (52, '系统配置详情', '/plat/sysConfig/get', '0', '8', '1', '0', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:17');
INSERT INTO `router` VALUES (53, '账号登出', '/auth/logout', '1', '0', '1', '0', NULL, '0', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (54, '账号权限信息', '/auth/mine', '0', '0', '1', '0', '', '1', '', '', '0', '0', '2024-08-23 14:06:48', '2024-08-23 14:06:48');

-- ----------------------------
-- Table structure for router_log
-- ----------------------------
DROP TABLE IF EXISTS `router_log`;
CREATE TABLE `router_log`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `app_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用名',
  `app_node` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用节点',
  `app_trace_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用节点TraceID',
  `req_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '来源IP',
  `req_url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求地址',
  `req_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求报文',
  `req_at` datetime(0) NULL DEFAULT NULL COMMENT '请求时间',
  `res_status` smallint(0) NULL DEFAULT NULL COMMENT '响应状态',
  `res_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应报文',
  `res_time` datetime(0) NULL DEFAULT NULL COMMENT '响应时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `log_uni`(`app_name`, `app_node`, `app_trace_id`) USING BTREE COMMENT '日志唯一性'
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '路由接口日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of router_log
-- ----------------------------
INSERT INTO `router_log` VALUES (1, 'Smart', 'APP01', 'HBVBCMKXQZ', '127.0.0.1:53772', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":0,\"showTotal\":true}', '2024-04-19 19:49:51', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/router/page\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/docs/sample\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-19 19:49:51');
INSERT INTO `router_log` VALUES (2, 'Smart', 'APP01', 'QNOOLTTANQ', '127.0.0.1:52832', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":0,\"showTotal\":true}', '2024-04-24 08:05:47', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/router/page\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/docs/sample\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:05:47');
INSERT INTO `router_log` VALUES (3, 'Smart', 'APP01', 'CJEGTISQOP', '127.0.0.1:53016', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:19:35', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/router/page\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/docs/sample\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:20:09');
INSERT INTO `router_log` VALUES (4, 'Smart', 'APP01', 'AYTHSTEIBV', '127.0.0.1:53043', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":0,\"showTotal\":true}', '2024-04-24 08:21:07', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/router/page\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/docs/sample\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:21:09');
INSERT INTO `router_log` VALUES (5, 'Smart', 'APP01', 'SCWPFBBPRD', '127.0.0.1:53068', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":0,\"showTotal\":true}', '2024-04-24 08:22:12', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/router/page\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/docs/sample\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:22:12');
INSERT INTO `router_log` VALUES (6, 'Smart', 'APP01', 'JGEVKAWNYU', '127.0.0.1:53130', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:22:58', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/router/page\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/docs/sample\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:22:58');
INSERT INTO `router_log` VALUES (7, 'Smart', 'APP01', 'JVSFCXBFUR', '127.0.0.1:53139', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":0,\"showTotal\":true}', '2024-04-24 08:23:06', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/router/page\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/docs/sample\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:23:06');
INSERT INTO `router_log` VALUES (8, 'Smart', 'APP01', 'PEWIGFFVPK', '127.0.0.1:53156', '/plat/router/page', '{(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)u(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)r(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)l(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855):(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855),\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:23:50', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/*******page(ba3c2168d1583ee6531ebd3caca23c0d05b1ba4afca96e13efb1d60a2f132965)\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"url\",\"resLogPrint\":\"0\",\"resLogSecure\":\"name,remark,url\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/do*****mple(2bac71247800d08d526fc863e8472714d97e2752ce9f28c292eed08a25cf8c73)\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:23:50');
INSERT INTO `router_log` VALUES (9, 'Smart', 'APP01', 'XFJUMKACAP', '127.0.0.1:53160', '/plat/router/page', '{(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)u(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)r(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)l(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855):(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855),\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:24:00', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/*******page(ba3c2168d1583ee6531ebd3caca23c0d05b1ba4afca96e13efb1d60a2f132965)\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"url\",\"resLogPrint\":\"0\",\"resLogSecure\":\"name,remark,url\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/do*****mple(2bac71247800d08d526fc863e8472714d97e2752ce9f28c292eed08a25cf8c73)\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:24:00');
INSERT INTO `router_log` VALUES (10, 'Smart', 'APP01', 'ADODKQEQQU', '127.0.0.1:53192', '/plat/router/page', '{(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)u(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)r(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)l(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855):(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855)\"(e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855),\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:26:30', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/*******page(ba3c2168d1583ee6531ebd3caca23c0d05b1ba4afca96e13efb1d60a2f132965)\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"url\",\"resLogPrint\":\"0\",\"resLogSecure\":\"name,remark,url\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/do*****mple(2bac71247800d08d526fc863e8472714d97e2752ce9f28c292eed08a25cf8c73)\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:26:30');
INSERT INTO `router_log` VALUES (11, 'Smart', 'APP01', 'NSGPJEUJTO', '127.0.0.1:53218', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:27:04', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/*******page(ba3c2168d1583ee6531ebd3caca23c0d05b1ba4afca96e13efb1d60a2f132965)\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"url\",\"resLogPrint\":\"0\",\"resLogSecure\":\"name,remark,url\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/do*****mple(2bac71247800d08d526fc863e8472714d97e2752ce9f28c292eed08a25cf8c73)\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:27:04');
INSERT INTO `router_log` VALUES (12, 'Smart', 'APP01', 'LRVYPXKLRY', '127.0.0.1:53229', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:28:25', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/*******page(ba3c2168d1583ee6531ebd3caca23c0d05b1ba4afca96e13efb1d60a2f132965)\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"url\",\"resLogPrint\":\"0\",\"resLogSecure\":\"name,remark,url\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/do*****mple(2bac71247800d08d526fc863e8472714d97e2752ce9f28c292eed08a25cf8c73)\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:28:25');
INSERT INTO `router_log` VALUES (13, 'Smart', 'APP01', 'BYWWGSGDCJ', '127.0.0.1:53246', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:28:41', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路由分页\",\"url\":\"/plat/*******page(ba3c2168d1583ee6531ebd3caca23c0d05b1ba4afca96e13efb1d60a2f132965)\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"url\",\"resLogPrint\":\"0\",\"resLogSecure\":\"name,remark,url\",\"remark\":\"路由接口：分页\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用API示例\",\"url\":\"/do*****mple(2bac71247800d08d526fc863e8472714d97e2752ce9f28c292eed08a25cf8c73)\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:29:01');
INSERT INTO `router_log` VALUES (14, 'Smart', 'APP01', 'RVNSSGXYLP', '127.0.0.1:53450', '/plat/router/page', '{\"url\":\"\",\"serviceCode\":\"\",\"type\":\"\",\"current\":1,\"pageSize\":10,\"total\":2,\"showTotal\":true}', '2024-04-24 08:38:27', 200, '{\"code\":\"S000\",\"msg\":\"处理成功\",\"data\":{\"list\":[{\"id\":2,\"type\":\"0\",\"name\":\"路***(8f7eae6a8128384f217fc63380bc4bc6e7dfdd606928b4158cf07d8bff4cae42)\",\"url\":\"/plat/*******page(ba3c2168d1583ee6531ebd3caca23c0d05b1ba4afca96e13efb1d60a2f132965)\",\"serviceCode\":\"3\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"url\",\"resLogPrint\":\"0\",\"resLogSecure\":\"name,remark,url\",\"remark\":\"路由****页(0c293be091a3ac2642ce148f731fdee1a36cfeecb85bf5cfdaeefbe4fdab0798)\",\"mark\":\"0\"},{\"id\":1,\"type\":\"1\",\"name\":\"通用****例(9d3bee121d5ed182a704e4f59245b2179f04c7e36c60432ad7b05b166bbf4c98)\",\"url\":\"/do*****mple(2bac71247800d08d526fc863e8472714d97e2752ce9f28c292eed08a25cf8c73)\",\"serviceCode\":\"0\",\"logInDb\":\"0\",\"reqLogPrint\":\"0\",\"reqLogSecure\":\"\",\"resLogPrint\":\"0\",\"resLogSecure\":\"\",\"remark\":\"\",\"mark\":\"0\"}],\"total\":2},\"unPop\":true}', '2024-04-24 08:38:27');

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `login_switch` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '并发限制开关，0限制 1不限制',
  `login_num` int(0) NULL DEFAULT NULL COMMENT '最大登陆并发量，最小为1',
  `login_fail_switch` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆失败限制开关，0限制 1不限制',
  `login_fail_unit` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆失败限制 1秒 2分 3时 4天 ',
  `login_fail_num` int(0) NULL DEFAULT NULL COMMENT '登陆失败最大尝试次数',
  `login_fail_lock_unit` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆失败锁定 1秒 2分 3时 4天 ',
  `login_fail_lock_num` int(0) NULL DEFAULT NULL COMMENT '登陆失败锁定时长',
  `login_fail_try_num` int(0) NULL DEFAULT NULL COMMENT '登陆失败尝试次数',
  `logout_switch` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆过期开关，0限制 1不限制',
  `logout_unit` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆过期单位，0永不过期 1秒 2分 3时 4天 ',
  `logout_num` int(0) NULL DEFAULT NULL COMMENT '登陆过期长度数量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '0', 1, '0', '2', 5, '3', 12, 5, '0', '2', 15);

SET FOREIGN_KEY_CHECKS = 1;
