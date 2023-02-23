package app

type Category struct {
	Id        int    `gorm:"id"`
	Name      string `gorm:"name"`       // 类目名称
	Keywords  string `gorm:"keywords"`   // 类目关键字，以JSON数组格式
	Desc      string `gorm:"desc"`       // 类目广告语介绍
	Pid       int    `gorm:"pid"`        // 父类目ID
	IconUrl   string `gorm:"icon_url"`   // 类目图标
	PicUrl    string `gorm:"pic_url"`    // 类目图片
	Level     int    `gorm:"level"`      // 类目层级
	SortOrder int    `gorm:"sort_order"` // 排序
	Deleted   int    `gorm:"deleted"`    // 逻辑删除
}

// 类目选项查询参数模型
type CategoryQueryParam struct {
	Pid   int `form:"pid"`
	Level int `form:"level"`
}

// 商品类目选项传输模型
type CategoryOption struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}
