CREATE TABLE `access_key` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `name` varchar(20) NOT NULL DEFAULT '' COMMENT 'ak 名称',
                              `ak_id` varchar(20) NOT NULL DEFAULT '' COMMENT 'Access Key Id',
                              `ak_secret` varchar(30) NOT NULL DEFAULT '' COMMENT 'Access Key Secret',
                              `is_root` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是 root 级 ak。1 表示是 | 0 表示不是',
                              `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 | 0：无效',
                              `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                              `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_ak_id` (`ak_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='access_key 表';

CREATE TABLE `role` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
                        `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 0：无效',
                        `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                        `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色表';

CREATE TABLE `access_key_role` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `access_key_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'access_key_id',
                                   `role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
                                   `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                                   PRIMARY KEY (`id`),
                                   INDEX (`access_key_id`),
                                   INDEX (`role_id`),
                                   FOREIGN KEY (access_key_id) REFERENCES access_key(id),
                                   FOREIGN KEY (role_id) REFERENCES role(id),
                                   KEY `idx_access_key_id` (`access_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='access_key 角色表';

CREATE TABLE `permission` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `title` varchar(50) NOT NULL DEFAULT '' COMMENT '权限名称',
                              `keys` varchar(1000) NOT NULL DEFAULT '' COMMENT 'json 数组',
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
                                   INDEX (`permission_id`),
                                   INDEX (`role_id`),
                                   FOREIGN KEY (role_id) REFERENCES role(id),
                                   FOREIGN KEY (permission_id) REFERENCES permission(id),
                                   KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色权限表';

CREATE TABLE `operation_log` (
                                 `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                 `access_key_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'access_key_id',
                                 `target_keys` varchar(255) NOT NULL DEFAULT '' COMMENT '访问的 keys',
                                 `query_params` longtext NOT NULL COMMENT 'get和post参数',
                                 `ua` varchar(255) NOT NULL DEFAULT '' COMMENT '访问 ua',
                                 `ip` varchar(32) NOT NULL DEFAULT '' COMMENT '访问 ip',
                                 `note` varchar(1000) NOT NULL DEFAULT '' COMMENT 'json 格式备注字段',
                                 `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 CONSTRAINT fk_access_key_id FOREIGN KEY (access_key_id) REFERENCES access_key(id),
                                 KEY `idx_access_key_id` (`access_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户操作记录表';



INSERT INTO `access_key` (`id`, `ak_id`, `ak_secret`, `is_root`, `status`, `updated_time`, `created_time`)
VALUES(1, 'LTAICfL25lEQjVKf', 'gwH4r9ysGwrlHRlKGAjBUhlNqx4VUX', 1, 1, '2019-08-28 16:36:30', '2019-08-28 16:36:30');