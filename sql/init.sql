create
database go_start;

CREATE TABLE `article`
(
    `id`              bigint unsigned auto_increment comment '自增主键',
    `title`           varchar(3000) NOT NULL DEFAULT '' COMMENT '标题',
    `content`         text COMMENT '内容',
    `updated_by`      varchar(64)   NOT NULL DEFAULT '' COMMENT '更新人ID',
    `created_by`      varchar(64)   NOT NULL DEFAULT '' COMMENT '创建人ID',
    `updated_by_name` varchar(128)  NOT NULL DEFAULT '' COMMENT '更新人姓名',
    `created_by_name` varchar(128)  NOT NULL DEFAULT '' COMMENT '创建人姓名',
    `is_delete`       tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除标记 0 未删除 1 已删除',
    `updated_at`      datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at`      datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY               `ix_created_at` (`created_at`) USING BTREE,
    KEY               `ix_updated_at` (`updated_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='文章';

CREATE TABLE `user`
(
    `id`              bigint unsigned auto_increment comment '自增主键',
    `user_id`         bigint       NOT NULL DEFAULT 0 COMMENT '用户ID',
    `user_name`       varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
    `user_phone`      varchar(50)  NOT NULL DEFAULT '' COMMENT '用户手机号码',
    `password`        varchar(200)  NOT NULL DEFAULT '' COMMENT '密码',
    `updated_by`      varchar(64)  NOT NULL DEFAULT '' COMMENT '更新人ID',
    `created_by`      varchar(64)  NOT NULL DEFAULT '' COMMENT '创建人ID',
    `updated_by_name` varchar(128) NOT NULL DEFAULT '' COMMENT '更新人姓名',
    `created_by_name` varchar(128) NOT NULL DEFAULT '' COMMENT '创建人姓名',
    `is_delete`       tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除标记 0 未删除 1 已删除',
    `updated_at`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `ix_user_id` (`user_id`) USING BTREE,
    UNIQUE KEY `ix_user_name` (`user_name`) USING BTREE,
    KEY               `ix_user_name_phone` (`user_name`, `user_phone`) USING BTREE,
    KEY               `ix_created_at` (`created_at`) USING BTREE,
    KEY               `ix_updated_at` (`updated_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户';

CREATE TABLE `user_id_create`
(
    `id`  bigint unsigned auto_increment comment '自增主键',
    `tub` tinyint(1) NOT NULL DEFAULT 0 COMMENT '数据',
    PRIMARY KEY (`id`),
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户id生成表';