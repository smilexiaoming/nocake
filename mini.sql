DROP TABLE IF EXISTS t_user;
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `wx_id` varchar(45) NOT NULL DEFAULT '未知' COMMENT '微信基础号',
  `user_name` varchar(45) NOT NULL DEFAULT '未知' COMMENT '微信号',
  `nick_name` varchar(45) NOT NULL DEFAULT '未知' COMMENT '微信昵称',
  `mini_openid` varchar(45) NOT NULL DEFAULT '' COMMENT 'minio_penid',
  `vip_level` varchar(45) NOT NULL DEFAULT '0' COMMENT '会员等级',
  `head_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_on` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO t_user SET wx_id = 'test1';
INSERT INTO t_user SET wx_id = 'test2';
INSERT INTO t_user SET wx_id = 'test3';
INSERT INTO t_user SET wx_id = 'test4';