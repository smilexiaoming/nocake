package app

import "time"

type OrderGoods struct {
	Id             int       `gorm:"id"`
	OrderId        int       `gorm:"order_id"`       // 订单表的订单ID
	GoodsId        int       `gorm:"goods_id"`       // 商品表的商品ID
	GoodsName      string    `gorm:"goods_name"`     // 商品名称
	GoodsNumber    int       `gorm:"goods_number"`   // 商品货品的购买数量
	Price          float64   `gorm:"price"`          // 商品货品的售价
	Specifications string    `gorm:"specifications"` // 商品货品的规格列表
	PicUrl         string    `gorm:"pic_url"`        // 商品货品图片或者商品图片
	GoodsComment   int       `gorm:"goods_comment"`  // 订单商品评论，如果是-1，则超期不能评价；如果是0，则可以评价；如果其他值，则是comment表里面的评论ID。
	Deleted        int       `gorm:"deleted"`        // 逻辑删除
	CreatedTime    time.Time `gorm:"created_time"`   // 创建时间
	UpdatedTime    time.Time `gorm:"updated_time"`   // 更新时间
	DeletedTime    time.Time `gorm:"deleted_time"`   // 删除时间
}
