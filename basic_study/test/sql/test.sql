# 签约模板
CREATE TABLE `pap_template` (
                                `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
                                `plan_id_wxmp` varchar(28) NOT NULL DEFAULT '' COMMENT '微信公众号 模板id',
                                `plan_id_wxapp` varchar(28) NOT NULL DEFAULT '' COMMENT '微信app 模板id',
                                `plan_id_alipay` varchar(28) NOT NULL DEFAULT '' COMMENT '支付宝 模板id',
                                `price` int NOT NULL DEFAULT 0 COMMENT '价格| 人民币分',
                                `cycle_period` int NOT NULL DEFAULT 1 COMMENT '周期-值',
                                `cycle_type` tinyint(3) NOT NULL DEFAULT 0 COMMENT '周期-类型：1: 日|2: 周|3: 月|4: 季|5: 年| 64: 秒',
                                `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '状态 1：有效 | 2：无效',
                                `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                                `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='签约模板';

# 签约记录
CREATE TABLE `contract` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `contract_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '签约id | 用雪花生成',
                            `pap_template_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '签约模板id',
                            `plan_id` varchar(28) NOT NULL DEFAULT '' COMMENT '模板id',
                            `pay_type` TINYINT(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型 [0:微信;2:QQ钱包;3:支付宝]',
                            `pay_platform` TINYINT(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付平台 [0:无效值;1:AndroiApp;2:iOSApp;3:微信公众号;4:H5商城;5:微信小程序;6:微信扫码;7:趣头条App;8:趣头条app(使用萌推微信H5支付)]',
                            `payment_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT 'Payment Id',
                            `buyer_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
                            `price` int NOT NULL DEFAULT 0 COMMENT '价格| 人民币分| 冗余字段：pap_template.price',
                            `cycle_period` int NOT NULL DEFAULT 1 COMMENT '周期-值',
                            `cycle_type` tinyint(3) NOT NULL DEFAULT 0 COMMENT '周期-类型：1: 日|2: 周|3: 月|4: 季|5: 年| 64: 秒',
                            `contract_time` datetime COMMENT '签约时间',
                            `delete_contract_time` datetime COMMENT '解约时间',
                            `contract_code`varchar(32) NOT NULL DEFAULT '' COMMENT '签约协议号',
                            `request_serial` bigint(20) NOT NULL DEFAULT 0 COMMENT '签约请求序列号',
                            `payment_contract_id` varchar(32) NOT NULL DEFAULT '' COMMENT '委托代扣协议id（对应结果通知中：[微信app -> contract_id, ]）',
                            `contract_expired_time` datetime COMMENT  '协议到期时间（指与微信协议）',
                            `contract_termination_mode` tinyint(3) NOT NULL DEFAULT 0 COMMENT '协议解约方式：0-未解约 1-有效期过启动解约 2-用户主动解约 3-商户API解约 4-商户平台解约 5-用户账号注销',
                            `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT ' 1 签约中，未支付| 2 签约成功| 3 签约失败|4 解约成功',
                            `contract_body` TEXT COMMENT '签约结果通知原文',
                            `delete_contract_body` TEXT COMMENT '解约通知原文',
                            `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                            `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                            PRIMARY KEY (`id`),
                            KEY (`buyer_id`),
                            UNIQUE KEY (`contract_id`),
                            UNIQUE KEY (`payment_id`),
                            UNIQUE KEY (`contract_code`),
                            UNIQUE KEY (`request_serial`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='签约记录';

# 委托扣款表
# 委托扣款时：发起扣款：insert payment -> 成功扣款：update payment
CREATE TABLE `pap_log` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `contract_id` bigint(20) unsigned NOT NULL COMMENT '签约 id',
                           `contract_code`varchar(32) NOT NULL DEFAULT '' COMMENT '签约协议号',
                           `plan_id` varchar(28) NOT NULL DEFAULT '' COMMENT '模板id',
                           `pay_type` TINYINT(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型 [0:微信;2:QQ钱包;3:支付宝]',
                           `pay_platform` TINYINT(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付平台 [0:无效值;1:AndroiApp;2:iOSApp;3:微信公众号;4:H5商城;5:微信小程序;6:微信扫码;7:趣头条App;8:趣头条app(使用萌推微信H5支付)]',
                           `buyer_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
                           `payment_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '续约每次扣款 Payment Id',
                           `price` int NOT NULL DEFAULT 0 COMMENT '价格| 人民币分| 冗余字段：pap.price',
                           `pap_expected_time` datetime COMMENT '预计扣款时间',
                           `pap_success_time`  datetime COMMENT '扣款成功时间',
                           `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '1 扣款中，未支付| 2 扣款成功| 3 扣款失败',
                           `notify_body` TEXT COMMENT '扣款通知原文',
                           `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
                           `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                           PRIMARY KEY (`id`),
                           KEY (`buyer_id`),
                           KEY (`contract_code`)
) ENGINE=InnoDB AUTO_INCREMENT=8000000 DEFAULT CHARSET=utf8mb4 COMMENT='委托扣款表';

INSERT INTO `pap_template` (`name`, `plan_id_wxmp`, `plan_id_wxapp` , `plan_id_alipay`, `price` , `cycle_period` , `cycle_type`)
VALUES
('连续包月', '130788', '130784', '', 3, 1, 3),
('连续包季', '130788', '130784', '', 4, 1, 4),
('连续包年', '130788', '130784', '', 5, 1, 5);


INSERT INTO `pap_template` ('id', 'name', 'plan_id_wxmp', 'plan_id_wxapp', 'plan_id_alipay', 'price', 'cycle_period', 'cycle_type')
VALUES (7, '连续包月', '130788', '130784', '', 3, 1, 3);


