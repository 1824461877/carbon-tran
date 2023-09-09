CREATE TABLE `pay_order` (
   `id` bigint NOT NULL AUTO_INCREMENT,
   `pay_order_id` varchar(255)  NOT NULL COMMENT '支付订单id',
   `pay_id` varchar(255) NOT NULL COMMENT '第三方支付的订单号',
   `initiator` varchar(255)  NOT NULL COMMENT '支付的发起者',
   `recipient` varchar(255)  NOT NULL COMMENT '支付接受者',
   `pay_status` int NOT NULL COMMENT '支付状态',
   `pay_amount` int NOT NULL COMMENT '支付金额',
   `initiator_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   `finish_time` timestamp DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`),
   UNIQUE KEY `pay_order_id_unique` (`pay_order_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;