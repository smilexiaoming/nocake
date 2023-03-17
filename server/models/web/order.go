package web

import (
	"nocake/models"
	"time"
)

type Order struct {
	Id            int       `gorm:"id" json:"id"`
	OpenId        string    `gorm:"open_id" json:"open_id"`               // 用户表的用户ID
	GoodsInfo     string    `gorm:"goods_info" json:"goods_info"`         // 商品信息json
	Status        int       `gorm:"status" json:"status"`                 // 订单状态 1已提交 2已完成 3撤销 4已接单不可撤销
	SubStatus     int       `gorm:"sub_status" json:"sub_status"`         // 订单子状态 1已提交 2已完成 3撤销 4已接单不可撤销
	Address       string    `gorm:"address" json:"address"`               // 收货具体地址
	Message       string    `gorm:"message" json:"message"`               // 用户订单留言
	GoodsPrice    float64   `gorm:"goods_price" json:"goods_price"`       // 商品总费用
	GoodsCount    int       `gorm:"goods_count" json:"goods_count"`       // 商品总数量
	CouponPrice   float64   `gorm:"coupon_price" json:"coupon_price"`     // 优惠券减免
	DisPrice      float64   `gorm:"dis_price" json:"dis_price"`           // 配送费用
	IntegralPrice float64   `gorm:"integral_price" json:"integral_price"` // 用户积分减免
	GrouponPrice  float64   `gorm:"groupon_price" json:"groupon_price"`   // 团购优惠价减免
	OrderPrice    float64   `gorm:"order_price" json:"order_price"`       // 订单费用， = goods_price + dis_price - coupon_price
	ActualPrice   float64   `gorm:"actual_price" json:"actual_price"`     // 实付费用， = order_price - integral_price
	PayId         string    `gorm:"pay_id" json:"pay_id"`                 // 微信付款编号
	PayTime       time.Time `gorm:"pay_time" json:"pay_time"`             // 微信付款时间
	ShipSn        string    `gorm:"ship_sn" json:"ship_sn"`               // 外卖订单
	ShipChannel   int       `gorm:"ship_channel" json:"ship_channel"`     // 外卖平台
	ShipTime      time.Time `gorm:"ship_time" json:"ship_time"`           // 发货开始时间
	ConfirmTime   time.Time `gorm:"confirm_time" json:"confirm_time"`     // 用户确认收货时间
	EndTime       time.Time `gorm:"end_time" json:"end_time"`             // 订单关闭时间
	Deleted       int       `gorm:"deleted" json:"deleted"`               // 逻辑删除
	CreatedTime   time.Time `gorm:"created_time" json:"created_time"`     // 创建时间
	UpdatedTime   time.Time `gorm:"updated_time" json:"updated_time"`     // 更新时间
	DeletedTime   time.Time `gorm:"deleted_time" json:"deleted_time"`     // 删除时间
}

// 商品删除参数模型
type OrderDeleteParam struct {
	Id int `form:"id" binding:"required,gt=0"`
}

// 商品更新参数模型
type OrderUpdateParam struct {
	Id     int `form:"id"`
	Status int `form:"status"` // 状态
}

// 商品列表查询参数模型
type OrderListParam struct {
	Page   models.Page
	Status int `form:"status"`
}

// 商品列表查询参数模型
type OrderDetailParam struct {
	Id int `form:"id"`
}

type OrderItem struct {
	Id            int       `json:"id"`
	OpenId        string    `json:"open_id"`        // 用户表的用户ID
	Avatar        string    `json:"avatar"`         // 用户头像图片
	Username      string    `json:"username"`       // 用户名称
	Nickname      string    `json:"nickname"`       // 用户昵称
	GoodsInfo     string    `json:"goods_info"`     // 商品信息json
	Status        int       `json:"status"`         // 订单状态 1已提交 2已完成 3撤销 4已接单不可撤销
	SubStatus     int       `json:"sub_status"`     // 订单子状态 1已提交 2已完成 3撤销 4已接单不可撤销
	Address       string    `json:"address"`        // 收货具体地址
	Message       string    `json:"message"`        // 用户订单留言
	GoodsPrice    float64   `json:"goods_price"`    // 商品总费用
	GoodsCount    int       `json:"goods_count"`    // 商品总数量
	CouponPrice   float64   `json:"coupon_price"`   // 优惠券减免
	DisPrice      float64   `json:"dis_price"`      // 配送费用
	IntegralPrice float64   `json:"integral_price"` // 用户积分减免
	GrouponPrice  float64   `json:"groupon_price"`  // 团购优惠价减免
	OrderPrice    float64   `json:"order_price"`    // 订单费用， = goods_price + dis_price - coupon_price
	ActualPrice   float64   `json:"actual_price"`   // 实付费用， = order_price - integral_price
	PayId         string    `json:"pay_id"`         // 微信付款编号
	PayTime       time.Time `json:"pay_time"`       // 微信付款时间
	ShipSn        string    `json:"ship_sn"`        // 外卖订单
	ShipChannel   int       `json:"ship_channel"`   // 外卖平台
	ShipTime      time.Time `json:"ship_time"`      // 发货开始时间
	ConfirmTime   time.Time `json:"confirm_time"`   // 用户确认收货时间
	EndTime       time.Time `json:"end_time"`       // 订单关闭时间
	Deleted       int       `json:"deleted"`        // 逻辑删除
	CreatedTime   time.Time `json:"created_time"`   // 创建时间
	UpdatedTime   time.Time `json:"updated_time"`   // 更新时间
	DeletedTime   time.Time `json:"deleted_time"`   // 删除时间
}

// 订单详情传输模型
type OrderDetail struct {
	Order
	GoodsItem []GoodsItem `json:"goodsItem"`
}

// 订单商品项传输模型
type GoodsItem struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Options string  `json:"options"`
	Price  float64 `json:"price"`
	PicUrl string  `json:"pic_url"`
	Count  int     `json:"count"`
}
