package app

import "time"

type Category struct {
	Id          int       `gorm:"id"`
	Name        string    `gorm:"name"`         // 类目名称
	Brief       string    `gorm:"brief"`        // 类目简介
	Keywords    string    `gorm:"keywords"`     // 类目关键字，以JSON数组格式
	Pid         string    `gorm:"pid"`          // 父类目ID
	Level       int       `gorm:"level"`        // 类目层级
	IconUrl     string    `gorm:"icon_url"`     // 类目图标
	Weight      int       `gorm:"weight"`       // 排序
	Deleted     int       `gorm:"deleted"`      // 逻辑删除
	CreatedTime time.Time `gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time"` // 删除时间
}

// 类目选项查询参数模型
type CategoryQueryParam struct {
	Pid   int `form:"pid"`
	Level int `form:"level"`
}

// 商品类目选项传输模型
type CategoryOption struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	IconUrl string `json:"icon_url"`
}
