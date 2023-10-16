CREATE TABLE assets_sell (
      `id` bigint NOT NULL AUTO_INCREMENT,
      `ex_id` varchar(255) NOT NULL COMMENT '交易单号',
      `ass_id` varchar(255) NOT NULL COMMENT '资产编号',
      `country` varchar(5) NOT NULL COMMENT '资产所属国家',
      `collection_wallet_id` varchar(255) NOT NULL COMMENT '收款的钱包 id',
      `user_id` varchar(255)  NOT NULL COMMENT '资产所属用户',
      `source` varchar(255)  NOT NULL COMMENT '来源资产',
      `amount` DOUBLE NOT NULL COMMENT 'number(数量)/amount(金额)',
      `number` int NOT NULL COMMENT '数量',
      `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易创建时间',
      `end_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易结束时间',
      PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;