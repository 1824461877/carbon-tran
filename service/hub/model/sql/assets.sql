CREATE TABLE `assets` (
       `id` bigint NOT NULL AUTO_INCREMENT,
       `ass_id` varchar(255) NOT NULL COMMENT '用于api操作',
       `hid` varchar(255) NOT NULL COMMENT '资产唯一标识',
       `user_id` varchar(255)  NOT NULL COMMENT '资产所属用户',
       `source` varchar(255)  NOT NULL COMMENT '来源资产',
       `number` int NOT NULL COMMENT '数量',
       `retire_number` int NOT NULL COMMENT '注销数量',
       `project`  varchar(255) NOT NULL COMMENT '项目名称',
       `gs_id` varchar(20) NOT NULL COMMENT 'gs_id',
       `status` varchar(20) NOT NULL COMMENT 'status',
       `country` varchar(20) NOT NULL COMMENT 'country',
       `project_type` varchar(20) NOT NULL COMMENT 'country',
       `product` varchar(20) NOT NULL COMMENT 'product',
       `vers_head` int NOT NULL COMMENT 'vers 序号区间 （head）',
       `vers_tail` int NOT NULL COMMENT 'vers 序号区间 （tail）',
       `serial_number` varchar(255)  NOT NULL COMMENT '资产编号',
       `day` int(4)  NOT NULL COMMENT '日期时间',
       `listing` bool NOT NULL COMMENT '挂牌',
       `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '资产创建时间',
       PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


# 1,"ab082c70-644a-3711-ac28-d4755b79d26a','H102771824-85975-85975",H102771824-85975-85975,6bad7442-b629-380d-a6ea-ffa16e3b1aa8,gsf,5,0,"Houji太阳
# 能炊具项目",GS7604,issued,中国,太阳热能,VER,85975,85979,GS1-1-CN-GS7604-3-2020-21289,2020,0,2023-06-10 15:53:13
