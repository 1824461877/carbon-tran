CREATE TABLE `user_wallet` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `wallet_id` varchar(255) NOT NULL COMMENT '钱包ID',
    `user_id` varchar(255) NOT NULL COMMENT '钱包所属用户ID',
    `name` varchar(255) NOT NULL COMMENT '钱包名',
    `password` varchar(255)  NOT NULL COMMENT '用户密码',
    `amount` DOUBLE NOT NULL COMMENT '钱包余额',
    `salt` varchar(255)  NOT NULL COMMENT '密码salt',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `wallet_id_unique` (`wallet_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;