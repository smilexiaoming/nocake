package app

import "time"

type WebUser struct {
	Id          int       `gorm:"id"`
	Username    string    `gorm:"username"`     // 用户名称
	Password    string    `gorm:"password"`     // 用户密码
	Avatar      string    `gorm:"avatar"`       // 用户头像图片
	Deleted     int       `gorm:"deleted"`      // 逻辑删除
	CreatedTime time.Time `gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time"` // 删除时间
}
