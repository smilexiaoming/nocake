package app

type GoodsProduct struct {
	Id             int     `gorm:"id"`
	GoodsId        int     `gorm:"goods_id"`       // 商品表的商品ID
	Specifications string  `gorm:"specifications"` // 商品规格值列表，采用JSON数组格式
	Price          float64 `gorm:"price"`          // 商品货品价格
	Producnumber   int     `gorm:"product_number"` // 商品货品数量
	Url            string  `gorm:"url"`            // 商品货品图片
	Deleted        int     `gorm:"deleted"`        // 逻辑删除
	CreatedTime    string  `gorm:"created_time"`   // 创建时间
	UpdatedTime    string  `gorm:"updated_time"`   // 更新时间
	DeletedTime    string  `gorm:"deleted_time"`   // 删除时间
}
