CREATE TABLE `users`
(
    `id`         bigint       NOT NULL AUTO_INCREMENT,
    `mobile`     varchar(11)  NOT NULL DEFAULT '' COMMENT '手机号',
    `password`   varchar(255) NOT NULL DEFAULT '' COMMENT '密码(MD5加密)',
    `nickname`   varchar(50)  NOT NULL DEFAULT '' COMMENT '用户昵称',
    `avatar`     varchar(255) NOT NULL DEFAULT '' COMMENT '头像URL',
    `status`     tinyint      NOT NULL DEFAULT '1' COMMENT '状态(1:正常 2:禁用)',
    `created_at` datetime              DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';