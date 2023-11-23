

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
) ENGINE=InnoDB AUTO_INCREMENT=111180 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of workflow_definition
-- ----------------------------
BEGIN;

INSERT INTO `workflow_definition` VALUES ('请假审批工作', 20, '请假流程审批', '2023-11-23 10:25:21');
COMMIT;

-- ----------------------------
-- Table structure for workflow_instance
-- ----------------------------
DROP TABLE IF EXISTS `workflow_instance`;
CREATE TABLE `workflow_instance` (
                                     `id` int(11) NOT NULL AUTO_INCREMENT,
                                     `workflow_definition_id` int(11) DEFAULT NULL,
                                     `current_step_id` int(11) DEFAULT NULL,
                                     `business_id` int(11) DEFAULT NULL,
                                     `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '1:已提交 2:审核通过 3:审核拒绝 5:已完成',
                                     `creation_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                     `assignee` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '冗余一下审核人',
                                     `creator` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '创建人',
                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=399 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of workflow_instance
-- ----------------------------
BEGIN;

INSERT INTO `workflow_instance` VALUES (398, 20, 16, 10086, '5', '2023-11-23 10:51:42', '王五', '小日本');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of workflow_step_definition
-- ----------------------------
BEGIN;

INSERT INTO `workflow_step_definition` VALUES (14, 20, 1, '主管审批', '张三', 15, '2023-11-23 10:26:22', -1);
INSERT INTO `workflow_step_definition` VALUES (15, 20, 2, '经理审批', '李四', 16, '2023-11-23 10:26:48', 14);
INSERT INTO `workflow_step_definition` VALUES (16, 20, 3, '总监审批', '王五', -1, '2023-11-23 10:26:49', 15);
COMMIT;

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

-- ----------------------------
-- Records of workflow_task
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
