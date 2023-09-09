CREATE TABLE `user_wallet` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `wallet_id` varchar(255) NOT NULL COMMENT '钱包ID',
    `user_id` varchar(255) NOT NULL COMMENT '钱包所属用户ID',
    `name` varchar(255) NOT NULL COMMENT '钱包名',
    `cid` varchar(255) NOT NULL COMMENT '第三方钱包id',
    `wallet_type` int(4) NOT NULL COMMENT '钱包类型',
    `default_collection` bool NOT NULL COMMENT '默认收款',
#     `password` varchar(255)  NOT NULL COMMENT '用户密码',
#     `amount` DOUBLE NOT NULL COMMENT '钱包余额',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `wallet_id_unique` (`wallet_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;