package app

import "time"

type Address struct {
	Id          int       `gorm:"id"`
	Name        string    `gorm:"name"`         // 收货人名称
	OpenId      string    `gorm:"open_id"`      // 用户表的用户ID
	Province    string    `gorm:"province"`     // 行政区域表的省ID
	City        string    `gorm:"city"`         // 行政区域表的市ID
	County      string    `gorm:"county"`       // 行政区域表的区县ID
	Detail      string    `gorm:"detail"`       // 详细收货地址
	Tel         string    `gorm:"tel"`          // 手机号码
	IsDefault   int       `gorm:"is_default"`   // 是否默认地址
	Deleted     int       `gorm:"deleted"`      // 逻辑删除
	CreatedTime time.Time `gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time"` // 删除时间
}
