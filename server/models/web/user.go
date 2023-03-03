package web

import "time"

// 用户数据映射模型
type User struct {
	Id          int       `gorm:"id"`
	Username    string    `gorm:"username"`     // 用户名称
	Password    string    `gorm:"password"`     // 用户密码
	Avatar      string    `gorm:"avatar"`       // 用户头像图片
	Deleted     int       `gorm:"deleted"`      // 逻辑删除
	CreatedTime time.Time `gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time"` // 删除时间
}

// 用户登录参数模型
type LoginParam struct {
	Username     string `form:"username" binding:"required"`
	Password     string `form:"password" binding:"required"`
	CaptchaId    string `form:"captchaId" binding:"required"`
	CaptchaValue string `form:"captchaValue" binding:"required"`
}

// 用户信息传输模型
type UserInfo struct {
	Sid   uint64 `form:"sid"`
	Token string `form:"token"`
}
