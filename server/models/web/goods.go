package web

import (
	"nocake/models"
	"time"
)

type Goods struct {
	Id          int       `gorm:"id" json:"id"`
	Name        string    `gorm:"name" json:"name"`               // 商品名称
	Brief       string    `gorm:"brief" json:"brief"`             // 商品简介
	CategoryId  int       `gorm:"category_id" json:"category_id"` // 商品所属类目ID
	Keywords    string    `gorm:"keywords" json:"keywords"`       // 商品关键字，采用逗号间隔
	Status      int       `gorm:"status" json:"status"`           // 状态
	Weight      int       `gorm:"weight" json:"weight"`
	PicUrl      string    `gorm:"pic_url" json:"pic_url"`           // 商品主图
	PicUrls     string    `gorm:"pic_urls" json:"pic_urls"`         // 商品附图JSON数组格式
	Unit        string    `gorm:"unit" json:"unit"`                 // 商品单位，例如件、盒
	Quantity    int       `gorm:"quantity" json:"quantity"`         // 商品库存
	Price       float64   `gorm:"price" json:"price"`               // 价格
	Deleted     int       `gorm:"deleted" json:"deleted"`           // 逻辑删除
	CreatedTime time.Time `gorm:"created_time" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time" json:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time" json:"deleted_time"` // 删除时间
}

// 商品创建参数模型
type GoodsCreateParam struct {
	Name       string  `form:"name"`        // 商品名称
	Brief      string  `form:"brief"`       // 商品简介
	Detail     string  `form:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId int     `form:"category_id"` // 商品所属类目ID
	Keywords   string  `form:"keywords"`    // 商品关键字，采用逗号间隔
	Status     int     `form:"status"`      // 状态
	Weight     int     `form:"weight"`
	PicUrl     string  `form:"pic_url"`  // 商品主图
	PicUrls    string  `form:"pic_urls"` // 商品附图JSON数组格式
	Unit       string  `form:"unit"`     // 商品单位，例如件、盒
	Price      float64 `form:"price"`    // 价格
}

// 商品删除参数模型
type GoodsDeleteParam struct {
	Id int `form:"id" binding:"required,gt=0"`
}

// 商品更新参数模型
type GoodsUpdateParam struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`        // 商品名称
	Brief      string  `json:"brief"`       // 商品简介
	Detail     string  `json:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId int     `json:"category_id"` // 商品所属类目ID
	Keywords   string  `json:"keywords"`    // 商品关键字，采用逗号间隔
	Status     int     `json:"status"`      // 状态
	Weight     int     `json:"weight"`
	PicUrl     string  `json:"pic_url"`  // 商品主图
	PicUrls    string  `json:"pic_urls"` // 商品附图JSON数组格式
	Unit       string  `json:"unit"`     // 商品单位，例如件、盒
	Quantity   int     `json:"quantity"` // 商品库存
	Price      float64 `json:"price"`    // 价格
}

// 商品状态更新参数模型
type GoodsStatusUpdateParam struct {
	Id     int `form:"id" binding:"required,gt=0"`
	Status int `form:"status" binding:"required,gt=0"`
}

// 商品列表查询参数模型
type GoodsListParam struct {
	Page       models.Page
	Id         int    `form:"id"`
	CategoryId int    `form:"category_id"`
	Name       string `form:"name"`
	Status     int    `form:"status"`
}

// 商品列表传输模型
type GoodsList struct {
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
	Quantity    int       `json:"quantity"`     // 商品库存
	Price       float64   `json:"price"`        // 价格
	CreatedTime time.Time `json:"created_time"` // 创建时间
}
