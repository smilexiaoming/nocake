-- 主表start
DROP TABLE IF EXISTS t_user;
CREATE TABLE `t_user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `open_id` varchar(63) NOT NULL DEFAULT "" COMMENT '微信登录openid',
    `nickname` varchar(63) NOT NULL DEFAULT "" COMMENT '用户昵称',
    `avatar` varchar(255) NOT NULL DEFAULT "" COMMENT '用户头像图片',
    `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别：0 未知， 1男， 1 女',
    `birthday` varchar(20) NOT NULL DEFAULT '0' COMMENT '生日',
    `user_level` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 普通用户，1 VIP用户，2 高级VIP用户',
    `tel` varchar(20) NOT NULL DEFAULT "" COMMENT '用户手机号码',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    KEY `open_id` (`open_id`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT '微信商城小程序-用户表';

DROP TABLE IF EXISTS t_category;
CREATE TABLE `t_category` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(63) NOT NULL DEFAULT "" COMMENT '类目名称',
    `brief` varchar(255) NOT NULL DEFAULT "" COMMENT '类目简介',
    `keywords` varchar(1023) NOT NULL DEFAULT "" COMMENT '类目关键字，以JSON数组格式',
    `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父类目ID',
    `level` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类目层级',
    `icon_url` varchar(255) NOT NULL DEFAULT "" COMMENT '类目图标',
    `weight` tinyint NOT NULL DEFAULT '50' COMMENT '排序',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    KEY `uk_pid_level` (`pid`,`level`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT '微信商城小程序-类目表';

INSERT INTO `t_category` (`id`, `name`, `keywords`, `brief`, `pid`, `icon_url`, `level`) VALUES (1, '蛋糕', '{"受众":"大众"}', "蛋糕，大家都喜欢", 0, 'http://tubiao.png', "1");
INSERT INTO `t_category` (`id`, `name`, `keywords`, `brief`, `pid`, `icon_url`, `level`) VALUES (2, '甜品', '{"受众":"小众"}', "甜品，少数人喜欢", 0, 'http://tubiao.png', "1");
INSERT INTO `t_category` (`id`, `name`, `keywords`, `brief`, `pid`, `icon_url`, `level`) VALUES (3, '甜品1', '{"受众":"小众"}', "甜品，少数人喜欢", 0, 'http://tubiao.png', "1");
INSERT INTO `t_category` (`id`, `name`, `keywords`, `brief`, `pid`, `icon_url`, `level`) VALUES (4, '蛋糕2', '{"受众":"大众"}', "蛋糕，大家都喜欢", 0, 'http://tubiao.png', "1");
INSERT INTO `t_category` (`id`, `name`, `keywords`, `brief`, `pid`, `icon_url`, `level`) VALUES (5, '甜品3', '{"受众":"小众"}', "甜品，少数人喜欢", 0, 'http://tubiao.png', "1");
INSERT INTO `t_category` (`id`, `name`, `keywords`, `brief`, `pid`, `icon_url`, `level`) VALUES (6, '甜品4', '{"受众":"小众"}', "甜品，少数人喜欢", 0, 'http://tubiao.png', "1");

DROP TABLE IF EXISTS t_goods;
CREATE TABLE `t_goods` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(127) NOT NULL DEFAULT "" COMMENT '商品名称',
    `brief` varchar(1023) NOT NULL DEFAULT "" COMMENT '商品简介',
    `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品所属类目ID',
    `keywords` varchar(255) NOT NULL DEFAULT "" COMMENT '商品关键字，采用逗号间隔',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1未上架 2已上架',
    `is_hot` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1不推荐 2推荐',
    `weight` tinyint(1) NOT NULL DEFAULT '100' COMMENT '',
    `pic_url` varchar(255) NOT NULL DEFAULT "" COMMENT '商品主图',
    `pic_urls` varchar(1023) NOT NULL DEFAULT "" COMMENT '商品附图JSON数组格式',
    `unit` varchar(31) NOT NULL DEFAULT '个' COMMENT '商品单位，例如件、盒',
    `quantity` int(11) NOT NULL DEFAULT '1' COMMENT '商品库存',
    `price` decimal(10, 2) NOT NULL DEFAULT '100000.00' COMMENT '价格',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    KEY `category_id` (`category_id`),
    KEY `weight` (`weight`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT '微信商城小程序-商品基本信息表';

INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (1, '水果蛋糕1', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (2, '水果蛋糕2', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (3, '水果蛋糕3', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (4, '水果蛋糕4', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (5, '水果蛋糕5', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (6, '水果蛋糕6', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (7, '水果蛋糕7', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (8, '水果蛋糕8', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (9, '水果蛋糕9', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (10, '水果蛋糕10', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (11, '水果蛋糕11', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (12, '水果蛋糕12', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (13, '水果蛋糕13', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (14, '水果蛋糕14', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (15, '水果蛋糕15', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (16, '水果蛋糕16', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (17, '水果蛋糕17', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (18, '水果蛋糕18', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (19, '水果蛋糕19', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (20, '水果蛋糕20', '很多水果',  1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (21, '甜品1', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (22, '甜品2', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (23, '甜品3', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (24, '甜品4', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (25, '甜品5', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (26, '甜品6', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (27, '甜品7', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (28, '甜品8', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (29, '甜品9', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (30, '甜品10', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (31, '甜品11', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (32, '甜品12', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (33, '甜品13', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (34, '甜品14', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (35, '甜品15', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (36, '甜品16', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (37, '甜品17', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (38, '甜品18', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (39, '甜品19', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `category_id`, `pic_urls`, `keywords`, `price`, `pic_url`) VALUES (40, '甜品20', '很多水果',  2, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128","http://1.14.106.241/images/cake.png");
DROP TABLE IF EXISTS t_order;
CREATE TABLE `t_order` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `open_id` varchar(63) NOT NULL DEFAULT "" COMMENT '用户表的用户ID',
    `goods_ids_count`varchar(1023) NOT NULL DEFAULT "" COMMENT '货品id数量json',
    `status` smallint(6) NOT NULL DEFAULT '1' COMMENT '订单状态 1已提交 2已完成 3撤销 4已接单不可撤销 ',
    `sub_status` smallint(6) NOT NULL DEFAULT '1' COMMENT '订单子状态 1已提交 2已完成 3撤销 4已接单不可撤销 ',
    `address` varchar(1023) NOT NULL  DEFAULT "" COMMENT '收货具体地址',
    `message` varchar(512) NOT NULL  DEFAULT "" COMMENT '用户订单留言',
    `goods_price` decimal(10, 2) NOT NULL DEFAULT "100000.00" COMMENT '商品总费用',
    `goods_count` tinyint NOT NULL DEFAULT "0" COMMENT '商品总数量',
    `coupon_price` decimal(10, 2) NOT NULL DEFAULT "0.00" COMMENT '优惠券减免',
    `dis_price` decimal(10, 2) NOT NULL DEFAULT "100000.00" COMMENT '配送费用',
    `integral_price` decimal(10, 2) NOT NULL DEFAULT "0.00" COMMENT '用户积分减免',
    `groupon_price` decimal(10, 2) NOT NULL DEFAULT "0.00" COMMENT '团购优惠价减免',
    `order_price` decimal(10, 2) NOT NULL DEFAULT "100000.00" COMMENT '订单费用， = goods_price + dis_price - coupon_price',
    `actual_price` decimal(10, 2) NOT NULL DEFAULT "100000.00" COMMENT '实付费用， = order_price - integral_price',
    `pay_id` varchar(63) NOT NULL DEFAULT "" COMMENT '微信付款编号',
    `pay_time` timestamp DEFAULT NULL COMMENT '微信付款时间',
    `ship_sn` varchar(63) NOT NULL DEFAULT "" COMMENT '外卖订单',
    `ship_channel` tinyint NOT NULL DEFAULT "0" COMMENT '外卖平台',
    `ship_time` timestamp NULL COMMENT '发货开始时间',
    `confirm_time` timestamp NULL COMMENT '用户确认收货时间',
    `end_time` timestamp NULL COMMENT '订单关闭时间',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT '微信商城小程序-订单表';
DROP TABLE IF EXISTS t_address;
CREATE TABLE `t_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(63) NOT NULL DEFAULT "" COMMENT '收货人名称',
    `open_id` varchar(63) NOT NULL DEFAULT "" COMMENT '用户表的用户',
    `province` varchar(63) NOT NULL DEFAULT "" COMMENT '行政区域表的省',
    `city` varchar(63) NOT NULL DEFAULT "" COMMENT '行政区域表的市',
    `county` varchar(63) NOT NULL DEFAULT "" COMMENT '行政区域表的区县',
    `detail` varchar(127) NOT NULL DEFAULT "" COMMENT '详细收货地址',
    `tel` varchar(20) NOT NULL DEFAULT "" COMMENT '手机号码',
    `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认地址',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    KEY `open_id` (`open_id`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT '微信商城小程序-收货地址表';
INSERT INTO `t_address` (`id`, `name`, `open_id`, `province`, `is_default`) VALUES (1, '张印', 'lainzhang', "成都", 1);

DROP TABLE IF EXISTS t_web_user;
CREATE TABLE `t_web_user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(63) NOT NULL DEFAULT "" COMMENT '用户名称',
    `password` varchar(63) NOT NULL DEFAULT "" COMMENT '用户密码',
    `avatar` varchar(255) NOT NULL DEFAULT "" COMMENT '用户头像图片',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `username` (`username`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT '微信商城web后台-用户表';

insert into t_web_user set username='zy', password='123456';
insert into t_web_user set username='gly', password='123456';