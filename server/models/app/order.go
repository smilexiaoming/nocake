package app

import (
	"time"
)

type Order struct {
	Id            int       `gorm:"id" json:"id"`
	OpenId        string    `gorm:"open_id" json:"open_id"`               // 用户表的用户ID
	Options       string    `gorm:"goods_info" json:"goods_info"`         // 商品信息json
	Status        int       `gorm:"status" json:"status"`                 // 订单状态 1已提交 2已完成 3撤销 4已接单不可撤销
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

// 订单更新状态模型
type OrderUpdateParam struct {
	OrderId int    `form:"order_id"`
	OpenId  string `form:"open_id"`
	Status  int    `form:"status"`
}

// 提交订单模型
type OrderSubmitParam struct {
	OpenId        string `form:"open_id"`
	Message       string `form:"message"` // 用户订单留言
	AddressId     int    `form:"address_id"`
	ChoiceAddress string `form:"choice_address"`
}

// 查询订单列表模型
type OrderQueryListParam struct {
	OpenId   string `form:"open_id"`
	Status   int    `form:"status"`
	PageNum  int    `form:"page_num required,gt=0" gorm:"required,gt=0"`
	PageSize int    `form:"page_size required,gt=0" gorm:"required,gt=0"`
}

// 查询订单模型
type OrderQueryDetailParam struct {
	OpenId  string `form:"open_id"`
	OrderId int    `form:"order_id"`
}

// 订单列表模型
type GoodsItem struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	PicUrl  string  `json:"pic_url"`
	Options string  `json:"options"`
}

type OrderInfo struct {
	OrderId    int         `json:"order_id"`
	Status     int         `json:"stauts"`
	TotalPrice float64     `json:"total_price"`
	GoodsCount int         `json:"goods_count"`
	GoodsItem  []GoodsItem `json:"goods_item"`
}
