CREATE TABLE `ip_address` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
                              `ip` varchar(20) NOT NULL DEFAULT '' COMMENT 'ip',
                              `mask` tinyint(2) NOT NULL DEFAULT '0' COMMENT '子网掩码 max 0, 32 即 0.0.0.0, 255.255.255.255',
                              `is_admin` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是 admin 级 ip。1 表示是 | 0 表示不是',
                              `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 | 0：无效',
                              `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                              `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_ak_id` (`ip`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='ip_address 表';

CREATE TABLE `role` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名称',
                        `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 0：无效',
                        `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                        `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色表';

CREATE TABLE `object_role` (
                               `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                               `object_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '对象类型 0: unknown 1：Ip Whitelisting 2：Access Key',
                               `object_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'object_id',
                               `role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
                               `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                               PRIMARY KEY (`id`),
                               KEY `idx_object_type_object_id` (`object_type`, `object_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='object 角色表';

CREATE TABLE `permission` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `title` varchar(64) NOT NULL DEFAULT '' COMMENT '权限名称',
                              `key` varchar(128) NOT NULL DEFAULT '' COMMENT 'key',
                              `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 0：无效',
                              `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                              `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='权限详情表';

CREATE TABLE `role_permission` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
                                   `permission_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '权限id',
                                   `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                                   PRIMARY KEY (`id`),
                                   KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色权限表';

CREATE TABLE `access_log` (
                                 `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                 `object_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '对象类型 0: unknown 1：Ip Whitelisting 2：Access Key',
                                 `object_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'object_id',
                                 `target_keys` varchar(255) NOT NULL DEFAULT '' COMMENT '访问的 keys',
                                 `query_params` longtext NOT NULL COMMENT 'get和post参数',
                                 `ua` varchar(255) NOT NULL DEFAULT '' COMMENT '访问 ua',
                                 `ip` varchar(32) NOT NULL DEFAULT '' COMMENT '访问 ip',
                                 `note` varchar(1000) NOT NULL DEFAULT '' COMMENT 'json 格式备注字段',
                                 `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 KEY `idx_object_type_object_id` (`object_type`, `object_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户操作记录表';



INSERT INTO `ip_address` (`id`, `ip`, `mask`, `is_root`, `status`, `updated_time`, `created_time`)
VALUES(1, '10.105.0.0', 16, 1, 1, '2019-08-28 16:36:30', '2019-08-28 16:36:30');