CREATE TABLE `gsf_asset` (
       `id` bigint NOT NULL AUTO_INCREMENT,
       `account` varchar(255)  NOT NULL COMMENT '账户',
       `details` varchar(255)  NOT NULL COMMENT 'Project Details',
       `gsf_id` varchar(255)  NOT NULL COMMENT 'GS ID',
       `number` int  NOT NULL COMMENT '数量',
       `day` int(4)  NOT NULL COMMENT '日期',
       PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;