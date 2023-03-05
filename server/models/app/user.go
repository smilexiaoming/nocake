package app

import (
	"time"
)

// 用户数据映射模型
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


// 用户登录凭证校验模型
type Code2Session struct {
	Code      string
	AppId     string
	AppSecret string
}

// 凭证校验后返回的JSON数据包模型
type Code2SessionResult struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    uint   `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// 用户信息,OpenID用户唯一标识
type UserInfo struct {
	OpenId string `json:"openId"`
}
