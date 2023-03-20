package app

// 购物车设置商品数量参数模型
type CartUpdateParam struct {
	GoodsId   uint   `form:"goods_id" binding:"required"`
	OpenId    string `form:"open_id" binding:"required"`
	Carnumber uint   `form:"cart_number" binding:"required"`
}

// 购物车修改商品数量参数模型
type CartSetParam struct {
	GoodsId uint   `form:"goods_id" binding:"required"`
	OpenId  string `form:"open_id" binding:"required"`
	Options string `form:"options" binding:"required"`
}

// 购物车删除商品参数模型
type CartDeleteParam struct {
	OpenId  string `form:"open_id" binding:"required"`
	GoodsId string `form:"goods_id" binding:"required"`
}

// 购物车清除参数模型
type CartClearParam struct {
	OpenId string `form:"open_id" binding:"required"`
}

// 购物车查询参数模型
type CartQueryParam struct {
	OpenId string `form:"open_id" binding:"required"`
}

// 购物车商品项传输模型
type CartItem struct {
	Id      uint64  `gorm:"primaryKey" json:"id"`
	PicUrl  string  `gorm:"pic_url"  json:"pic_url"`
	Name    string  `gorm:"name"      json:"name"`
	Price   float64 `gorm:"price"      json:"price"`
	Options string  `gorm:"options"      json:"options"`
}

// 购物车信息传输模型
type CartInfo struct {
	CartItem   []CartItem `json:"cart_item"`
	TotalCart  int        `json:"total_cart"`
	TotalPrice float64    `json:"total_price"`
}

type OptionItem struct {
	Name   string
	Price  int
	Active bool
}
type Option struct {
	Name string       `json:"name"`
	Item []OptionItem `json:"item"`
	Type int          `json:"type"`
}
type Options struct {
	Option []Option `json:"option"`
	Count  int      `json:"count"`
}
