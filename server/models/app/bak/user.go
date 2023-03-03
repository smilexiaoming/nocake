package app

import "time"

type User struct {
	Id           int       `gorm:"id"`
	OpenId       string    `gorm:"open_id"`         // 微信登录openid
	Nickname     string    `gorm:"nickname"`        // 用户昵称
	Avatar       string    `gorm:"avatar"`          // 用户头像图片
	Gender       int       `gorm:"gender"`          // 性别：0 未知， 1男， 1 女
	Birthday     string    `gorm:"birthday"`        // 生日
	LasloginTime time.Time `gorm:"last_login_time"` // 最近一次登录时间
	UserLevel    int       `gorm:"user_level"`      // 0 普通用户，1 VIP用户，2 高级VIP用户
	Tel          string    `gorm:"tel"`             // 用户手机号码
	Deleted      int       `gorm:"deleted"`         // 逻辑删除
	CreatedTime  time.Time `gorm:"created_time"`    // 创建时间
	UpdatedTime  time.Time `gorm:"updated_time"`    // 更新时间
	DeletedTime  time.Time `gorm:"deleted_time"`    // 删除时间
}
