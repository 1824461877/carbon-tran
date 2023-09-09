CREATE TABLE `trade_order` (
   `id` bigint NOT NULL AUTO_INCREMENT,
   `trade_order_id` varchar(255) NOT NULL COMMENT '交易订单id',
   `pay_order_id` varchar(255) NOT NULL COMMENT '支付订单id',
   `exchange_asset_id` varchar(255) NOT NULL COMMENT '交易所订单 id',
   `carbon_asset_id` varchar(255)  NOT NULL COMMENT '用户密码',
   `collection_id` varchar(255)  NOT NULL COMMENT '收款账号',
   `initiator` varchar(255)  NOT NULL COMMENT '交易的发起者',
   `recipient` varchar(255)  NOT NULL COMMENT '交易接受者',
   `trade_status` int NOT NULL COMMENT '交易状态',
   `number` int NOT NULL COMMENT '数量',
   `initiator_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   `finish_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`),
   UNIQUE KEY `trade_order_id_unique` (`trade_order_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;