package app

type Category struct {
	Id          int    `gorm:"id"`
	Name        string `gorm:"name"`         // 类目名称
	Keywords    string `gorm:"keywords"`     // 类目关键字，以JSON数组格式
	Desc        string `gorm:"desc"`         // 类目广告语介绍
	Pid         int    `gorm:"pid"`          // 父类目ID
	IconUrl     string `gorm:"icon_url"`     // 类目图标
	PicUrl      string `gorm:"pic_url"`      // 类目图片
	Level       int    `gorm:"level"`        // 类目层级
	Sororder    int    `gorm:"sort_order"`   // 排序
	Deleted     int    `gorm:"deleted"`      // 逻辑删除
	CreatedTime string `gorm:"created_time"` // 创建时间
	UpdatedTime string `gorm:"updated_time"` // 更新时间
	DeletedTime string `gorm:"deleted_time"` // 删除时间
}
