CREATE TABLE `ip_address` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
                              `ip` varchar(20) NOT NULL DEFAULT '' COMMENT 'ip',
                              `mask` varchar(20) NOT NULL DEFAULT '255.255.255.255' COMMENT '子网掩码 max 8, 32 即 255.0.0.0, 255.255.255.255（认为没有网络号，只代表一个主机号）',
                              `is_admin` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是 admin 级 ip。1 表示是 | 0 表示不是',
                              `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 | 0：无效',
                              `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                              `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_ak_id` (`ip`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='ip_address 表';

CREATE TABLE `role` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名称',
                        `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 0：无效',
                        `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                        `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

CREATE TABLE `object_role` (
                               `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `object_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '对象类型 0: unknown 1：Ip Whitelisting 2：Access Key',
    `object_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'object_id',
    `role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
    `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY (`id`),
    KEY `idx_object_type_object_id` (`object_type`, `object_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='object 角色表';

CREATE TABLE `permission` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `title` varchar(64) NOT NULL DEFAULT '' COMMENT '权限名称',
                              `entity_key` varchar(128) NOT NULL DEFAULT '' COMMENT 'entity_key',
                              `permission` tinyint(3) NOT NULL DEFAULT '' COMMENT 'permission: rwx',
                              `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 0：无效',
                              `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除标志 1：删除 0：未删除',
                              `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                              `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `unique_entity_key_permission` (`entity_key`, `permission`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限详情表';

CREATE TABLE `role_permission` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
                                   `permission_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '权限id',
                                   `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                                   PRIMARY KEY (`id`),
                                   KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限表';

CREATE TABLE `access_log` (
                                 `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                 `object_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '对象类型 0: unknown 1：Ip Whitelisting 2：Access Key',
                                 `object_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'object_id',
                                 `target_entity_keys` varchar(255) NOT NULL DEFAULT '' COMMENT '访问的 entity_keys',
                                 `query_params` longtext NOT NULL COMMENT 'get和post参数',
                                 `ua` varchar(255) NOT NULL DEFAULT '' COMMENT '访问 ua',
                                 `ip` varchar(32) NOT NULL DEFAULT '' COMMENT '访问 ip',
                                 `note` varchar(1000) NOT NULL DEFAULT '' COMMENT 'json 格式备注字段',
                                 `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 KEY `idx_object_type_object_id` (`object_type`, `object_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户操作记录表';



INSERT INTO `ip_address` (`id`, `name`, `ip`, `mask`, `is_admin`, `status`, `updated_time`, `created_time`)
VALUES(1, 'internal', '10.105.0.0', '255.255.0.0', 1, 1, '2019-08-28 16:36:30', '2019-08-28 16:36:30');


CREATE TABLE `ak_config` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `namespace` varchar(50) NOT NULL DEFAULT 'default' COMMENT '命名空间',
     `name` varchar(100) NOT NULL COMMENT 'name 唯一',
     `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
     `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
     `value` longtext COMMENT 'key 对应值，建议 json',
     `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ak 配置';