package web

import "time"

type Category struct {
	Id          int       `gorm:"id"`
	Name        string    `gorm:"name"`         // 类目名称
	Brief       string    `gorm:"brief"`        // 类目简介
	Keywords    string    `gorm:"keywords"`     // 类目关键字，以JSON数组格式
	Pid         int       `gorm:"pid"`          // 父类目ID
	Level       int       `gorm:"level"`        // 类目层级
	IconUrl     string    `gorm:"icon_url"`     // 类目图标
	Weight      int       `gorm:"weight"`       // 排序
	Deleted     int       `gorm:"deleted"`      // 逻辑删除
	CreatedTime time.Time `gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time"` // 删除时间
}

// 商品类目创建参数模型
type CategoryCreateParam struct {
	Name     string `form:"name"`     // 类目名称
	Brief    string `form:"brief"`    // 类目简介
	Keywords string `form:"keywords"` // 类目关键字，以JSON数组格式
	Pid      int    `form:"pid"`      // 父类目ID
	Level    int    `form:"level"`    // 类目层级
	IconUrl  string `form:"icon_url"` // 类目图标
	Weight   int    `form:"weight"`   // 排序
}

// 商品类目删除参数模型
type CategoryDeleteParam struct {
	Id int `form:"id" binding:"required,gt=0"`
}

// 商品类目更新参数模型
type CategoryUpdateParam struct {
	Id       int    `form:"id" binding:"required,gt=0"`
	Name     string `form:"name"`     // 类目名称
	Brief    string `form:"brief"`    // 类目简介
	Keywords string `form:"keywords"` // 类目关键字，以JSON数组格式
	Pid      int    `form:"pid"`      // 父类目ID
	Level    int    `form:"level"`    // 类目层级
	IconUrl  string `form:"icon_url"` // 类目图标
	Weight   int    `form:"weight"`   // 排序
}

// 商品类目查询参数模型
type CategoryQueryParam struct {
	Name  string `form:"name"`  // 类目名称
	Pid   int    `form:"pid"`   // 父类目ID
	Level int    `form:"level"` // 类目层级
}

// 商品类目列表传输模型
type CategoryList struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Pid    int    `json:"pid"`
	Level  int    `json:"level"`
	Weight int    `json:"weight"`
}

// 商品类目选项传输模型
type CategoryOption struct {
	Value    int              `json:"value"`
	Label    string           `json:"label"`
	Children []CategoryOption `json:"children"`
}
