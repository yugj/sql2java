-- user
CREATE TABLE `t_user`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `name`        varchar(20) NOT NULL COMMENT 'name',
    `birth`       datetime    NOT NULL COMMENT 'birth day',
    `age`         int         NOT NULL COMMENT 'age',
    `is_deleted`  tinyint(1) NOT NULL COMMENT '是否删除',
    `create_by`   varchar(20) NOT NULL COMMENT '创建人',
    `modify_by`   varchar(20) NOT NULL COMMENT '修改人',
    `create_time` datetime    NOT NULL COMMENT '创建时间',
    `modify_time` datetime    NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='user';

-- product
CREATE TABLE `t_product`
(
    `id`               bigint(20) NOT NULL AUTO_INCREMENT,
    `product_id`       bigint(20) NOT NULL COMMENT '产品编号',
    `area_id`          bigint(20) NOT NULL COMMENT '区域编号',
    `area_alias_name`  varchar(20) NOT NULL COMMENT '区域面别名',
    `created_user_id`  bigint(20) NOT NULL COMMENT '创建人',
    `modified_user_id` bigint(20) NOT NULL COMMENT '修改人',
    `created`          datetime    NOT NULL COMMENT '创建时间',
    `modified`         datetime    NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_product_id` (`product_id`, `area_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='产品设计区域表';
