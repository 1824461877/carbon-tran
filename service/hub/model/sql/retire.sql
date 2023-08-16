CREATE TABLE `retire` (
      `id` bigint NOT NULL AUTO_INCREMENT,
      `r_id` varchar(255)  NOT NULL COMMENT '注销单号ID',
      `ass_id` varchar(255)  NOT NULL COMMENT '注销的资产id',
      `user_id` varchar(255) NOT NULL COMMENT '资产注销用户id',
      `number` int  NOT NULL COMMENT '数量',
      `status` int(1) NOT NULL COMMENT '注销状态',
      `certificate_link` varchar(255) NOT NULL COMMENT '证书访问link',
      `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注销时间',
      PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;