package app

import (
	"time"
)

// 数据库地址信息
type Address struct {
	Id          int       `gorm:"id" json:"id"`
	Name        string    `gorm:"name" json:"name"`                 // 收货人名称
	OpenId      string    `gorm:"open_id" json:"open_id"`           // 用户表的用户ID
	Detail      string    `gorm:"detail" json:"detail"`             // 详细收货地址
	Tel         string    `gorm:"tel" json:"tel"`                   // 手机号码
	IsDefault   int       `gorm:"is_default" json:"is_default"`     // 是否默认地址
	Deleted     int       `gorm:"deleted" json:"deleted"`           // 逻辑删除
	CreatedTime time.Time `gorm:"created_time" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time" json:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time" json:"deleted_time"` // 删除时间
}

// 新增地址
type AddressAddParam struct {
	Name      string `form:"name" json:"name"`             // 收货人名称
	OpenId    string `form:"open_id" json:"open_id"`       // 用户表的用户ID
	Detail    string `form:"detail" json:"detail"`         // 详细收货地址
	Tel       string `form:"tel" json:"tel"`               // 手机号码
	IsDefault int    `form:"is_default" json:"is_default"` // 是否默认地址
}

// 删除地址
type AddressDeleteParam struct {
	AddressId int    `form:"address_id"` // 地址id
	OpenId    string `form:"open_id"`    // 用户表的用户ID
}

// 查看地址列表
type AddressListQueryParam struct {
	OpenId string `form:"open_id"` // 用户表的用户ID
}

// 地址列表
type AddressList struct {
	Address
}

// 更新地址
type AddressUpdateParam struct {
	AddressId int    `form:"address_id"`
	Name      string `form:"name"`       // 收货人名称
	OpenId    string `form:"open_id"`    // 用户表的用户ID
	Detail    string `form:"detail"`     // 详细收货地址
	Tel       string `form:"tel"`        // 手机号码
	IsDefault int    `form:"is_default"` // 是否默认地址
}

// 查看地址详情
type AddressQueryParam struct {
	AddressId int    `form:"address_id"`
	OpenId    string `form:"open_id"` // 用户表的用户ID
}

// 地址详情
type AddressDetail struct {
	Address
}
