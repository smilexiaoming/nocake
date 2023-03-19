package app

import "time"

type Order struct {
	Id            int       `gorm:"id"`
	OpenId        string    `gorm:"open_id"`         // 用户表的用户ID
	GoodsIdsCount string    `gorm:"goods_ids_count"` // 货品id数量json
	Status        int       `gorm:"status"`          // 订单状态 1已提交 2已完成 3撤销 4已接单不可撤销
	Address       string    `gorm:"address"`         // 收货具体地址
	Message       string    `gorm:"message"`         // 用户订单留言
	GoodsPrice    float64   `gorm:"goods_price"`     // 商品总费用
	GoodsCount    int       `gorm:"goods_count"`     // 商品总数量
	CouponPrice   float64   `gorm:"coupon_price"`    // 优惠券减免
	DisPrice      float64   `gorm:"dis_price"`       // 配送费用
	IntegralPrice float64   `gorm:"integral_price"`  // 用户积分减免
	GrouponPrice  float64   `gorm:"groupon_price"`   // 团购优惠价减免
	OrderPrice    float64   `gorm:"order_price"`     // 订单费用， = goods_price + dis_price - coupon_price
	ActualPrice   float64   `gorm:"actual_price"`    // 实付费用， = order_price - integral_price
	PayId         string    `gorm:"pay_id"`          // 微信付款编号
	PayTime       time.Time `gorm:"pay_time"`        // 微信付款时间
	ShipSn        string    `gorm:"ship_sn"`         // 外卖订单
	ShipChannel   int       `gorm:"ship_channel"`    // 外卖平台
	ShipTime      time.Time `gorm:"ship_time"`       // 发货开始时间
	ConfirmTime   time.Time `gorm:"confirm_time"`    // 用户确认收货时间
	EndTime       time.Time `gorm:"end_time"`        // 订单关闭时间
	Deleted       int       `gorm:"deleted"`         // 逻辑删除
	CreatedTime   time.Time `gorm:"created_time"`    // 创建时间
	UpdatedTime   time.Time `gorm:"updated_time"`    // 更新时间
	DeletedTime   time.Time `gorm:"deleted_time"`    // 删除时间
}
