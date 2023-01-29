CREATE TABLE `user_info`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`           varchar(128)    NOT NULL DEFAULT '' COMMENT '姓名',
    `user_name`      varchar(128)    NOT NULL DEFAULT '' COMMENT '用户名',
    `password`       varchar(128)    NOT NULL DEFAULT '' COMMENT '用户密码',
    `follow_count`   int(8)          NOT NULL DEFAULT 0 COMMENT '关注数量',
    `follower_count` int(8)          NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    `created_at`     timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='用户信息表';

