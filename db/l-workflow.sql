
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for approval_history
-- ----------------------------
DROP TABLE IF EXISTS `approval_history`;
CREATE TABLE `approval_history` (
                                    `instance_id` int(11) DEFAULT NULL,
                                    `id` int(11) NOT NULL AUTO_INCREMENT,
                                    `step_id` int(11) DEFAULT NULL,
                                    `participant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                    `result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '1已经提交，2审核通过，3审核拒绝 ，4重新发起提交，5已完成',
                                    `res_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
                                    `business_id` int(11) DEFAULT NULL,
                                    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                    `reason` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '拒绝原因',
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=111152 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for workflow_definition
-- ----------------------------
DROP TABLE IF EXISTS `workflow_definition`;
CREATE TABLE `workflow_definition` (
                                       `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                       `id` int(11) NOT NULL AUTO_INCREMENT,
                                       `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                       `create_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for workflow_instance
-- ----------------------------
DROP TABLE IF EXISTS `workflow_instance`;
CREATE TABLE `workflow_instance` (
                                     `id` int(11) NOT NULL AUTO_INCREMENT,
                                     `workflow_definition_id` int(11) DEFAULT NULL,
                                     `current_step_id` int(11) DEFAULT NULL,
                                     `business_id` int(11) DEFAULT NULL,
                                     `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '1;//已提交\n2;//审核通过\n3;//审核拒绝\n5;//已完成',
                                     `creation_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                     `assignee` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '冗余一下审核人',
                                     `creator` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '创建人',
                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=390 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for workflow_step_definition
-- ----------------------------
DROP TABLE IF EXISTS `workflow_step_definition`;
CREATE TABLE `workflow_step_definition` (
                                            `id` int(11) NOT NULL AUTO_INCREMENT,
                                            `workflow_definition_id` int(11) DEFAULT NULL COMMENT '流程id',
                                            `step_number` int(11) DEFAULT NULL COMMENT '步骤编号',
                                            `step_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '步骤名称',
                                            `assignee` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '审核人或角色',
                                            `next_step_id` int(11) DEFAULT NULL COMMENT '下一步',
                                            `create_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                            `pre_step_id` int(11) DEFAULT NULL COMMENT '上一步',
                                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for workflow_task
-- ----------------------------
DROP TABLE IF EXISTS `workflow_task`;
CREATE TABLE `workflow_task` (
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `workflow_instance_id` int(11) DEFAULT NULL,
                                 `step_definition_id` int(11) DEFAULT NULL,
                                 `business_id` int(11) DEFAULT NULL,
                                 `assignee` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `creation_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                 `completion_time` datetime DEFAULT NULL,
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

SET FOREIGN_KEY_CHECKS = 1;
