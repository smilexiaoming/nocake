package app

import "time"

type Cart struct {
	Id             int       `gorm:"id"`
	OpenId         string    `gorm:"open_id"`        // 用户表的用户ID
	GoodsId        int       `gorm:"goods_id"`       // 商品表的商品ID
	GoodsName      string    `gorm:"goods_name"`     // 商品名称
	Price          float64   `gorm:"price"`          // 商品货品的价格
	Carnumber      int       `gorm:"cart_number"`    // 商品货品的数量
	Specifications string    `gorm:"specifications"` // 商品规格值列表，采用JSON数组格式
	Checked        int       `gorm:"checked"`        // 购物车中商品是否选择状态
	PicUrl         string    `gorm:"pic_url"`        // 商品图片或者商品货品图片
	Deleted        int       `gorm:"deleted"`        // 逻辑删除
	CreatedTime    time.Time `gorm:"created_time"`   // 创建时间
	UpdatedTime    time.Time `gorm:"updated_time"`   // 更新时间
	DeletedTime    time.Time `gorm:"deleted_time"`   // 删除时间
}
