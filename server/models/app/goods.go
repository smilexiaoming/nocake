package app

import "time"

type Goods struct {
	Id          int       `gorm:"id"`
	Name        string    `gorm:"name"`        // 商品名称
	Brief       string    `gorm:"brief"`       // 商品简介
	Detail      string    `gorm:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId  int       `gorm:"category_id"` // 商品所属类目ID
	Keywords    string    `gorm:"keywords"`    // 商品关键字，采用逗号间隔
	Status      int       `gorm:"status"`      // 状态
	Weight      int       `gorm:"weight"`
	PicUrl      string    `gorm:"pic_url"`      // 商品主图
	PicUrls     string    `gorm:"pic_urls"`     // 商品附图JSON数组格式
	Unit        string    `gorm:"unit"`         // 商品单位，例如件、盒
	Price       float64   `gorm:"price"`        // 价格
	Deleted     int       `gorm:"deleted"`      // 逻辑删除
	CreatedTime time.Time `gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time"` // 删除时间
}

// 商品列表参数
type GoodsListQueryParam struct {
	CategoryId int `form:"category_id" binding:"required,gt=0"`
	PageNum    int `form:"page_num" gorm:"required,gt=0"`
	PageSize   int `form:"page_size" gorm:"required,gt=0"`
}

// 商品列表
type GoodsList struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`        // 商品名称
	Brief        string  `json:"brief"`       // 商品简介
	Detail       string  `json:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId   int     `json:"category_id"` // 商品所属类目ID
	Gallery      string  `json:"gallery"`     // 商品宣传图片列表，采用JSON数组格式
	Keywords     string  `json:"keywords"`    // 商品关键字，采用逗号间隔
	IsOnSale     int     `json:"is_on_sale"`  // 是否上架
	Weight       int     `json:"weight"`
	PicUrl       string  `json:"pic_url"`       // 商品页面商品图片
	IsNew        int     `json:"is_new"`        // 是否新品首发，如果设置则可以在新品首发页面展示
	IsHot        int     `json:"is_hot"`        // 是否人气推荐，如果设置则可以在人气推荐页面展示
	Unit         string  `json:"unit"`          // 商品单位，例如件、盒
	CounterPrice float64 `json:"counter_price"` // 专柜价格
	RetailPrice  float64 `json:"retail_price"`  // 零售价格
	Price        float64 `json:"price"`         // 线上价格
}

// 商品详情参数
type GoodsDetailQueryParam struct {
	GoodsId int `form:"goods_id" binding:"required,gt=0"`
}

// 商品详情
type GoodsDetail struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`        // 商品名称
	Brief       string    `json:"brief"`       // 商品简介
	Detail      string    `json:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId  int       `json:"category_id"` // 商品所属类目ID
	Keywords    string    `json:"keywords"`    // 商品关键字，采用逗号间隔
	Status      int       `json:"status"`      // 状态
	Weight      int       `json:"weight"`
	PicUrl      string    `json:"pic_url"`      // 商品主图
	PicUrls     string    `json:"pic_urls"`     // 商品附图JSON数组格式
	Unit        string    `json:"unit"`         // 商品单位，例如件、盒
	Price       float64   `json:"price"`        // 价格
}

// 商品详情参数
type GoodsSearchQueryParam struct {
	Name     string `form:"name"`     // 商品名称
	Keywords string `form:"keywords"` // 商品关键字，采用逗号间隔
	PageNum  int    `form:"page_num" gorm:"required,gt=0"`
	PageSize int    `form:"page_size" gorm:"required,gt=0"`
}
