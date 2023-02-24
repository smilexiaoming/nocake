package app

type Goods struct {
	Id           int     `gorm:"id"`
	Name         string  `gorm:"name"`        // 商品名称
	Brief        string  `gorm:"brief"`       // 商品简介
	Detail       string  `gorm:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId   int     `gorm:"category_id"` // 商品所属类目ID
	Gallery      string  `gorm:"gallery"`     // 商品宣传图片列表，采用JSON数组格式
	Keywords     string  `gorm:"keywords"`    // 商品关键字，采用逗号间隔
	IsOnSale     int     `gorm:"is_on_sale"`  // 是否上架
	Sororder     int     `gorm:"sort_order"`
	PicUrl       string  `gorm:"pic_url"`       // 商品页面商品图片
	ShareUrl     string  `gorm:"share_url"`     // 商品分享朋友圈图片
	IsNew        int     `gorm:"is_new"`        // 是否新品首发，如果设置则可以在新品首发页面展示
	IsHot        int     `gorm:"is_hot"`        // 是否人气推荐，如果设置则可以在人气推荐页面展示
	Unit         string  `gorm:"unit"`          // 商品单位，例如件、盒
	CounterPrice float64 `gorm:"counter_price"` // 专柜价格
	RetailPrice  float64 `gorm:"retail_price"`  // 零售价格
	Price        float64 `gorm:"price"`         // 线上价格
	Deleted      int     `gorm:"deleted"`       // 逻辑删除
	CreatedTime  string  `gorm:"created_time"`  // 创建时间
	UpdatedTime  string  `gorm:"updated_time"`  // 更新时间
	DeletedTime  string  `gorm:"deleted_time"`  // 删除时间
}

// 商品列表参数
type GoodsListQueryParam struct {
	CategoryId int `form:"category_id" binding:"required,gt=0"`
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
	SortOrder    int     `json:"sort_order"`
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
	Id int `form:"id" binding:"required,gt=0"`
}

// 商品详情
type GoodsDetail struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`        // 商品名称
	Brief        string  `json:"brief"`       // 商品简介
	Detail       string  `json:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId   int     `json:"category_id"` // 商品所属类目ID
	Gallery      string  `json:"gallery"`     // 商品宣传图片列表，采用JSON数组格式
	Keywords     string  `json:"keywords"`    // 商品关键字，采用逗号间隔
	IsOnSale     int     `json:"is_on_sale"`  // 是否上架
	SortOrder    int     `json:"sort_order"`
	PicUrl       string  `json:"pic_url"`       // 商品页面商品图片
	IsNew        int     `json:"is_new"`        // 是否新品首发，如果设置则可以在新品首发页面展示
	IsHot        int     `json:"is_hot"`        // 是否人气推荐，如果设置则可以在人气推荐页面展示
	Unit         string  `json:"unit"`          // 商品单位，例如件、盒
	CounterPrice float64 `json:"counter_price"` // 专柜价格
	RetailPrice  float64 `json:"retail_price"`  // 零售价格
	Price        float64 `json:"price"`         // 线上价格
}

// 商品详情参数
type GoodsSearchQueryParam struct {
	Name     string `form:"name"`     // 商品名称
	Keywords string `form:"keywords"` // 商品关键字，采用逗号间隔
}
