package app

import "time"

type Order struct {
	Id            int       `gorm:"id"`
	OpenId        string    `gorm:"open_id"`         // 用户表的用户ID
	GoodsIdsCount string    `gorm:"goods_ids_count"` // 货品id列表
	Status        int       `gorm:"status"`          // 订单状态
	Consignee     string    `gorm:"consignee"`       // 收货人名称
	Mobile        string    `gorm:"mobile"`          // 收货人手机号
	Address       string    `gorm:"address"`         // 收货具体地址
	Message       string    `gorm:"message"`         // 用户订单留言
	GoodsPrice    float64   `gorm:"goods_price"`     // 商品总费用
	CouponPrice   float64   `gorm:"coupon_price"`    // 优惠券减免
	DisPrice      float64   `gorm:"dis_price"`       // 配送费用
	IntegralPrice float64   `gorm:"integral_price"`  // 用户积分减免
	GrouponPrice  float64   `gorm:"groupon_price"`   // 团购优惠价减免
	OrderPrice    float64   `gorm:"order_price"`     // 订单费用， = goods_price + dis_price - coupon_price
	ActualPrice   float64   `gorm:"actual_price"`    // 实付费用， = order_price - integral_price
	PayId         string    `gorm:"pay_id"`          // 微信付款编号
	PayTime       time.Time `gorm:"pay_time"`        // 微信付款时间
	ShipSn        string    `gorm:"ship_sn"`         // 发货编号
	ShipChannel   string    `gorm:"ship_channel"`    // 发货快递公司
	ShipTime      time.Time `gorm:"ship_time"`       // 发货开始时间
	ConfirmTime   time.Time `gorm:"confirm_time"`    // 用户确认收货时间
	Comments      int       `gorm:"comments"`        // 待评价订单商品数量
	EndTime       time.Time `gorm:"end_time"`        // 订单关闭时间
	Deleted       int       `gorm:"deleted"`         // 逻辑删除
	CreatedTime   time.Time `gorm:"created_time"`    // 创建时间
	UpdatedTime   time.Time `gorm:"updated_time"`    // 更新时间
	DeletedTime   time.Time `gorm:"deleted_time"`    // 删除时间
}

// 订单更新状态模型
type OrderUpdateParam struct {
	Id     int    `form:"id"`
	OpenId string `form:"open_id"`
	Status int    `form:"status"`
}

// 提交订单模型
type OrderSubmitParam struct {
	OpenId string `form:"open_id"`
}

// 查询订单列表模型
type OrderQueryListParam struct {
	OpenId   string `form:"open_id"`
	Status   int    `form:"status"`
	PageNum  int    `form:"page_num" gorm:"default 1"`
	PageSize int    `form:"page_size" gorm:"default 10"`
}

// 查询订单模型
type OrderQueryDetailParam struct {
	OpenId  string `form:"open_id"`
	Status  int    `form:"status"`
	OrderId int    `form:"order_id"`
}

// 订单列表模型
type GoodsItem struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	PicUrl string  `json:"pic_url"`
	Count  int     `json:"count"`
}

type OrderDetail struct {
	Id         int         `json:"id"`
	Status     int         `json:"stauts"`
	TotalPrice float64     `json:"total_price"`
	GoodsCount uint        `json:"goods_count"`
	GoodsItem  []GoodsItem `json:"goods_item"`
}
