package app

type Cart struct {
	Id             int     `gorm:"id"`
	OpenId         int     `gorm:"open_id"`        // 用户表的用户ID
	GoodsId        int     `gorm:"goods_id"`       // 商品表的商品ID
	GoodsName      string  `gorm:"goods_name"`     // 商品名称
	Price          float64 `gorm:"price"`          // 商品货品的价格
	Carnumber      int     `gorm:"cart_number"`    // 商品货品的数量
	Specifications string  `gorm:"specifications"` // 商品规格值列表，采用JSON数组格式
	Checked        int     `gorm:"checked"`        // 购物车中商品是否选择状态
	PicUrl         string  `gorm:"pic_url"`        // 商品图片或者商品货品图片
	Deleted        int     `gorm:"deleted"`        // 逻辑删除
	CreatedTime    string  `gorm:"created_time"`   // 创建时间
	UpdatedTime    string  `gorm:"updated_time"`   // 更新时间
	DeletedTime    string  `gorm:"deleted_time"`   // 删除时间
}

// 购物车添加参数模型
type CartAddParam struct {
	GoodsId   uint   `form:"goods_id" binding:"required"`
	OpenId    string `form:"open_id" binding:"required"`
	Carnumber uint   `form:"cart_number" binding:"required"`
}

// 购物车查询参数模型
type CartQueryParam struct {
	OpenId string `form:"open_id" binding:"required"`
}

// 购物车清除参数模型
type CartClearParam struct {
	OpenId string `form:"open_id" binding:"required"`
}

// 购物车删除或减少商品参数模型
type CartDeleteParam struct {
	OpenId    string `form:"open_id" binding:"required"`
	GoodsId   string `form:"goods_id" binding:"required"`
	Carnumber int    `form:"cart_number"`
}

// 购物车商品项传输模型
type CartItem struct {
	Id        uint64  `gorm:"primaryKey" json:"id"`
	Url       string  `gorm:"url"  json:"url"`
	GoodsName string  `gorm:"goods_name"      json:"goods_name"`
	Price     float64 `gorm:"price"      json:"price"`
	Carnumber int     `gorm:"cart_number"      json:"cart_number"`
}

// 购物车信息传输模型
type CartInfo struct {
	CartItem   []CartItem `json:"cart_item"`
	TotalCart  int        `json:"total_cart"`
	TotalPrice float64    `json:"total_price"`
}
