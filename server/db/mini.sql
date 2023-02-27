-- 主表start
DROP TABLE IF EXISTS t_user;
CREATE TABLE `t_user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `open_id` varchar(63) NOT NULL DEFAULT "" COMMENT '微信登录openid',
    `wx_id` varchar(63) NOT NULL DEFAULT "" NULL COMMENT '用户原始微信id',
    `username` varchar(63) NOT NULL DEFAULT "" COMMENT '用户名称',
    `nickname` varchar(63) NOT NULL DEFAULT "" COMMENT '用户昵称',
    `password` varchar(63) NOT NULL DEFAULT "" COMMENT '用户密码',
    `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 可用, 1 禁用, 2 注销',
    `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别：0 未知， 1男， 1 女',
    `birthday` varchar(20) NOT NULL DEFAULT '0' COMMENT '生日',
    `last_login_time` timestamp DEFAULT NULL COMMENT '最近一次登录时间',
    `last_login_ip` varchar(63) NOT NULL DEFAULT "" COMMENT '最近一次登录IP地址',
    `user_level` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 普通用户，1 VIP用户，2 高级VIP用户',
    `mobile` varchar(20) NOT NULL COMMENT '用户手机号码',
    `avatar` varchar(255) NOT NULL COMMENT '用户头像图片',
    `session_key` varchar(100) NOT NULL DEFAULT "" COMMENT '微信登录会话KEY',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `username` (`username`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT '微信商城小程序-用户表';

DROP TABLE IF EXISTS t_category;
CREATE TABLE `t_category` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(63) NOT NULL DEFAULT "" COMMENT '类目名称',
    `keywords` varchar(1023) NOT NULL DEFAULT "" COMMENT '类目关键字，以JSON数组格式',
    `desc` varchar(255) NOT NULL DEFAULT "" COMMENT '类目广告语介绍',
    `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父类目ID',
    `icon_url` varchar(255) NOT NULL DEFAULT "" COMMENT '类目图标',
    `pic_url` varchar(255) NOT NULL DEFAULT "" COMMENT '类目图片',
    `level` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类目层级',
    `sort_order` tinyint(3) NOT NULL DEFAULT '50' COMMENT '排序',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    KEY `pid` (`pid`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT '微信商城小程序-类目表';

INSERT INTO `t_category` (`id`, `name`, `keywords`, `desc`, `pid`, `icon_url`, `pic_url`, `level`) VALUES (1, '蛋糕', '{"受众":"大众"}', "蛋糕，大家都喜欢", 0, 'http://tubiao.png', 'http://tupian.png', "1");
INSERT INTO `t_category` (`id`, `name`, `keywords`, `desc`, `pid`, `icon_url`, `pic_url`, `level`) VALUES (2, '甜品', '{"受众":"小众"}', "甜品，少数人喜欢", 0, 'http://tubiao.png', 'http://tupian.png', "1");

DROP TABLE IF EXISTS t_goods;
CREATE TABLE `t_goods` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(127) NOT NULL DEFAULT "" COMMENT '商品名称',
    `brief` varchar(255) NOT NULL DEFAULT "" COMMENT '商品简介',
    `detail` text NULL COMMENT '商品详细介绍，是富文本格式',
    `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品所属类目ID',
    `gallery` varchar(1023) NOT NULL DEFAULT "" COMMENT '商品宣传图片列表，采用JSON数组格式',
    `keywords` varchar(255) NOT NULL DEFAULT "" COMMENT '商品关键字，采用逗号间隔',
    `is_on_sale` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否上架',
    `sort_order` smallint(4) NOT NULL DEFAULT '100' COMMENT '',
    `pic_url` varchar(255) NOT NULL DEFAULT "" COMMENT '商品页面商品图片',
    `share_url` varchar(255) NOT NULL DEFAULT "" COMMENT '商品分享朋友圈图片',
    `is_new` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否新品首发，如果设置则可以在新品首发页面展示',
    `is_hot` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否人气推荐，如果设置则可以在人气推荐页面展示',
    `unit` varchar(31) NOT NULL DEFAULT '个' COMMENT '商品单位，例如件、盒',
    `counter_price` decimal(10, 2) NOT NULL DEFAULT '100000.00' COMMENT '专柜价格',
    `retail_price` decimal(10, 2) NOT NULL DEFAULT '100000.00' COMMENT '零售价格',
    `price` decimal(10, 2) NOT NULL DEFAULT '100000.00' COMMENT '线上价格',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    KEY `category_id` (`category_id`),
    KEY `sort_order` (`sort_order`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT '微信商城小程序-商品基本信息表';

INSERT INTO `t_goods` (`id`, `name`, `brief`, `detail`, `category_id`, `gallery`, `keywords`, `price`) VALUES (1, '水果蛋糕', '很多水果', "丰富原材料", 1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `detail`, `category_id`, `gallery`, `keywords`, `price`) VALUES (2, '水果蛋糕2', '很多水果', "丰富原材料", 1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128");
INSERT INTO `t_goods` (`id`, `name`, `brief`, `detail`, `category_id`, `gallery`, `keywords`, `price`) VALUES (3, '水果蛋糕2', '很多水果', "丰富原材料", 1, '["http://tubiao.png","http://tubiao1.png"]', "大众,水果,蛋糕", "128");

-- DROP TABLE IF EXISTS t_cart;
-- CREATE TABLE `t_cart` (
--     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
--     `open_id` varchar(63) NOT NULL DEFAULT "" COMMENT '用户表的用户ID',
--     `goods_id` int(11) NULL COMMENT '商品表的商品ID',
--     `goods_name` varchar(127) NULL COMMENT '商品名称',
--     `price` decimal(10, 2) NULL DEFAULT '0.00' COMMENT '商品货品的价格',
--     `cart_number` smallint(5) NULL DEFAULT '0' COMMENT '商品货品的数量',
--     `specifications` varchar(1023) NULL COMMENT '商品规格值列表，采用JSON数组格式',
--     `checked` tinyint(1) NULL DEFAULT '1' COMMENT '购物车中商品是否选择状态',
--     `pic_url` varchar(255) NULL COMMENT '商品图片或者商品货品图片',
--     `deleted` tinyint(1) NULL DEFAULT '0' COMMENT '逻辑删除',
--     `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--     `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
--     PRIMARY KEY (`id`)
-- ) ENGINE = InnoDB COMMENT '微信商城小程序-购物车商品表';
DROP TABLE IF EXISTS t_order;
CREATE TABLE `t_order` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `open_id` varchar(63) NOT NULL DEFAULT "" COMMENT '用户表的用户ID',
    `goods_ids_count`varchar(1023) NOT NULL DEFAULT "" COMMENT '货品id数量json',
    `status` smallint(6) NOT NULL DEFAULT '1' COMMENT '订单状态 1已提交 2已完成 3撤销 4已接单不可撤销 ',
    `consignee` varchar(63) NOT NULL DEFAULT "" COMMENT '收货人名称',
    `mobile` varchar(63) NOT NULL DEFAULT "" COMMENT '收货人手机号',
    `address` varchar(1023) NOT NULL  DEFAULT "" COMMENT '收货具体地址',
    `message` varchar(512) NOT NULL  DEFAULT "" COMMENT '用户订单留言',
    `goods_price` decimal(10, 2) NOT NULL DEFAULT "100000.00" COMMENT '商品总费用',
    `coupon_price` decimal(10, 2) NOT NULL DEFAULT "0.00" COMMENT '优惠券减免',
    `dis_price` decimal(10, 2) NOT NULL DEFAULT "100000.00" COMMENT '配送费用',
    `integral_price` decimal(10, 2) NOT NULL DEFAULT "0.00" COMMENT '用户积分减免',
    `groupon_price` decimal(10, 2) NOT NULL DEFAULT "0.00" COMMENT '团购优惠价减免',
    `order_price` decimal(10, 2) NOT NULL DEFAULT "100000.00" COMMENT '订单费用， = goods_price + dis_price - coupon_price',
    `actual_price` decimal(10, 2) NOT NULL DEFAULT "100000.00" COMMENT '实付费用， = order_price - integral_price',
    `pay_id` varchar(63) NOT NULL DEFAULT "" COMMENT '微信付款编号',
    `pay_time` timestamp DEFAULT NULL COMMENT '微信付款时间',
    `ship_sn` varchar(63) NOT NULL DEFAULT "" COMMENT '发货编号',
    `ship_channel` varchar(63) NOT NULL DEFAULT "" COMMENT '发货快递公司',
    `ship_time` timestamp NULL COMMENT '发货开始时间',
    `confirm_time` timestamp NULL COMMENT '用户确认收货时间',
    `comments` smallint(6) NOT NULL DEFAULT '0' COMMENT '待评价订单商品数量',
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
    `open_id` varchar(63) NOT NULL DEFAULT "" COMMENT '用户表的用户ID',
    `province` varchar(63) NOT NULL DEFAULT "" COMMENT '行政区域表的省ID',
    `city` varchar(63) NOT NULL DEFAULT "" COMMENT '行政区域表的市ID',
    `county` varchar(63) NOT NULL DEFAULT "" COMMENT '行政区域表的区县ID',
    `detail` varchar(127) NOT NULL DEFAULT "" COMMENT '详细收货地址',
    `area_code` char(6) NOT NULL DEFAULT "" COMMENT '地区编码',
    `postal_code` char(6) NOT NULL DEFAULT "" COMMENT '邮政编码',
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
DROP TABLE IF EXISTS t_comment;
CREATE TABLE `t_comment` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `value_id` int(11) NOT NULL DEFAULT '0' COMMENT '如果type=0，则是商品评论；如果是type=1，则是专题评论。',
    `comment_type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '评论类型，如果type=0，则是商品评论；如果是type=1，则是专题评论；如果type=3，则是订单商品评论。',
    `content` varchar(1023) NOT NULL COMMENT '评论内容',
    `open_id` varchar(63) NOT NULL DEFAULT "" COMMENT '用户表的用户ID',
    `has_picture` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否含有图片',
    `pic_urls` varchar(1023) NOT NULL DEFAULT "" COMMENT '图片地址列表，采用JSON数组格式',
    `star` smallint(6) NOT NULL DEFAULT '1' COMMENT '评分， 1-5',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` timestamp DEFAULT NULL COMMENT '删除时间',
    KEY `value_id` (`value_id`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT '微信商城小程序-评论表';
-- 主表end