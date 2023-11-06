/*
 Author                : Pokeya

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Host           : 127.0.0.1:3306
 Source Schema         : mydb

 Target Server Type    : MySQL

 Date: 03/11/2023
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_rbac_role
-- ----------------------------
DROP TABLE IF EXISTS `t_rbac_role`;
CREATE TABLE `t_rbac_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'Primary key ID',
  `role_name` varchar(36) NOT NULL COMMENT 'Role name',
  `role_desc` text COMMENT 'Role description',
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Creation time',
  `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'Update time',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT 'Deletion time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_name` (`role_name`),
  KEY `idx_role_name` (`role_name`),
  KEY `idx_t_rbac_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of t_rbac_role
-- ----------------------------
BEGIN;
INSERT INTO `t_rbac_role` (`id`, `role_name`, `role_desc`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'root', 'Super administrator', '2023-11-03 16:16:00.786', '2023-11-03 16:16:00.786', NULL);
INSERT INTO `t_rbac_role` (`id`, `role_name`, `role_desc`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'admin', 'Administrator', '2023-11-03 16:16:00.793', '2023-11-03 16:16:00.793', NULL);
INSERT INTO `t_rbac_role` (`id`, `role_name`, `role_desc`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'user', 'Regular user', '2023-11-03 16:16:00.798', '2023-11-03 16:16:00.798', NULL);
INSERT INTO `t_rbac_role` (`id`, `role_name`, `role_desc`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'guest', 'Guest', '2023-11-03 16:16:00.801', '2023-11-03 16:16:00.801', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_rbac_user
-- ----------------------------
DROP TABLE IF EXISTS `t_rbac_user`;
CREATE TABLE `t_rbac_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL COMMENT 'Unique identifier',
  `name` varchar(50) NOT NULL COMMENT 'English name',
  `nickname` varchar(50) NOT NULL COMMENT 'Chinese name',
  `employee_id` bigint DEFAULT NULL COMMENT 'Employee ID',
  `employee_type` varchar(50) DEFAULT NULL COMMENT 'Employee type',
  `company` varchar(50) DEFAULT NULL COMMENT 'Company name',
  `work_country` varchar(50) DEFAULT NULL COMMENT 'Work country',
  `work_city` varchar(50) DEFAULT NULL COMMENT 'Work city',
  `department` varchar(50) DEFAULT NULL COMMENT 'Department',
  `email` varchar(50) NOT NULL COMMENT 'Email address',
  `avatar` text COMMENT 'User avatar URL',
  `username` varchar(50) DEFAULT NULL COMMENT 'System login username',
  `password` varchar(255) DEFAULT NULL COMMENT 'Local system login password',
  `is_active` tinyint(1) DEFAULT '1' COMMENT 'Whether the user is active (0 inactive, 1 active)',
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Creation time',
  `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'Update time',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT 'Deletion time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_email` (`email`),
  KEY `idx_username` (`username`),
  KEY `idx_password` (`password`),
  KEY `idx_is_active` (`is_active`),
  KEY `idx_t_rbac_user_deleted_at` (`deleted_at`),
  KEY `idx_uuid` (`uuid`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of t_rbac_user
-- ----------------------------
BEGIN;
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'efe0e38b-dde7-41ae-b3c9-3289bc3f5e2d', 'super', 'super-icon', 1001, 'Full-time', 'FutureCompany', 'China', 'Beijing', 'IT', 'super@example.com', 'https://example.com/user/avatar/1/avatar.png', 'super', 'super@0000', 1, '2023-11-03 16:16:01', '2023-11-03 16:16:01', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'abcdef01-2345-6789-abcd-ef0123456789', 'admin1', 'admin1', 1002, 'Full-time', 'FutureCompany', 'China', 'Beijing', 'IT', 'admin1@example.com', 'https://example.com/user/avatar/2/avatar.png', 'admin1', 'admin@0001', 1, '2023-11-03 16:16:02', '2023-11-03 16:16:02', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'abcdef01-5678-9abc-def0-123456789abc', 'admin2', 'admin2', 1003, 'Full-time', 'FutureCompany', 'China', 'Shanghai', 'HR', 'admin2@example.com', 'https://example.com/user/avatar/3/avatar.png', 'admin2', 'admin@0002', 1, '2023-11-03 16:16:03', '2023-11-03 16:16:03', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, '34567a89-bcde-0123-4567-89abcde01234', 'user1', 'user1', 1004, 'Full-time', 'FutureCompany', 'China', 'Shanghai', 'Sales', 'user1@example.com', 'https://example.com/user/avatar/4/avatar.png', 'user1', 'password1', 1, '2023-11-03 16:16:04', '2023-11-03 16:16:04', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, '45678b9a-cdef-1234-5678-9abcdef01234', 'user2', 'user2', 1005, 'Part-time', 'FutureCompany', 'China', 'Beijing', 'Finance', 'user2@example.com', 'https://example.com/user/avatar/5/avatar.png', 'user2', 'password2', 1, '2023-11-03 16:16:05', '2023-11-03 16:16:05', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, '56789cde-f012-3456-789a-bcdef0123456', 'user3', 'user3', 1006, 'Full-time', 'FutureCompany', 'China', 'Shanghai', 'Marketing', 'user3@example.com', 'https://example.com/user/avatar/6/avatar.png', 'user3', 'password3', 1, '2023-11-03 16:16:06', '2023-11-03 16:16:06', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, '6789def0-1234-5678-9abc-def012345678', 'user4', 'user4', 1007, 'Contractor', 'FutureCompany', 'China', 'Beijing', 'HR', 'user4@example.com', 'https://example.com/user/avatar/7/avatar.png', 'user4', 'password4', 1, '2023-11-03 16:16:07', '2023-11-03 16:16:07', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, '789abcde-2345-6789-abcd-ef0123456789', 'user5', 'user5', 1008, 'Full-time', 'FutureCompany', 'China', 'Shanghai', 'Operations', 'user5@example.com', 'https://example.com/user/avatar/8/avatar.png', 'user5', 'password5', 1, '2023-11-03 16:16:08', '2023-11-03 16:16:08', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, '89abcdef-3456-789a-bcde-f0123456789a', 'guest1', 'guest1', NULL, NULL, NULL, NULL, NULL, NULL, 'guest1@example.com', 'https://example.com/user/avatar/9/avatar.png', 'guest1', '000000', 1, '2023-11-03 16:16:09', '2023-11-03 16:16:09', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, '9abcdef0-4567-89ab-cdef-0123456789ab', 'guest2', 'guest2', NULL, NULL, NULL, NULL, NULL, NULL, 'guest2@example.com', 'https://example.com/user/avatar/10/avatar.png', 'guest2', '000000', 1, '2023-11-03 16:16:10', '2023-11-03 16:16:10', NULL);
INSERT INTO `t_rbac_user` (`id`, `uuid`, `name`, `nickname`, `employee_id`, `employee_type`, `company`, `work_country`, `work_city`, `department`, `email`, `avatar`, `username`, `password`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 'f703c47d-f533-4fa2-8ded-14dc3358a77e', 'guest3', 'guest3', NULL, NULL, NULL, NULL, NULL, NULL, 'guest3@example.com', 'https://example.com/user/avatar/11/avatar.png', 'guest3', '000000', 1, '2023-11-03 16:16:11', '2023-11-03 16:16:11', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_rbac_user_roles
-- ----------------------------
DROP TABLE IF EXISTS `t_rbac_user_roles`;
CREATE TABLE `t_rbac_user_roles` (
  `user_id` bigint unsigned NOT NULL COMMENT 'User ID',
  `role_id` bigint NOT NULL COMMENT 'Role ID',
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `fk_t_rbac_user_roles_role` (`role_id`),
  CONSTRAINT `fk_t_rbac_user_roles_role` FOREIGN KEY (`role_id`) REFERENCES `t_rbac_role` (`id`),
  CONSTRAINT `fk_t_rbac_user_roles_user` FOREIGN KEY (`user_id`) REFERENCES `t_rbac_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of t_rbac_user_roles
-- ----------------------------
BEGIN;
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (1, 1);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (2, 2);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (3, 2);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (4, 3);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (5, 3);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (6, 3);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (7, 3);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (8, 3);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (9, 4);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (10, 4);
INSERT INTO `t_rbac_user_roles` (`user_id`, `role_id`) VALUES (11, 4);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;