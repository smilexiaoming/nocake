package app

import "time"

// 用户数据映射模型
type User struct {
	Id            int       `gorm:"id"`
	OpenId        string    `gorm:"open_id"`         // 微信登录openid
	WxId          string    `gorm:"wx_id"`           // 用户原始微信id
	Username      string    `gorm:"username"`        // 用户名称
	Nickname      string    `gorm:"nickname"`        // 用户昵称
	Password      string    `gorm:"password"`        // 用户密码
	Status        int       `gorm:"status"`          // 0 可用, 1 禁用, 2 注销
	Gender        int       `gorm:"gender"`          // 性别：0 未知， 1男， 1 女
	Birthday      string    `gorm:"birthday"`        // 生日
	LastLoginTime time.Time `gorm:"last_login_time"` // 最近一次登录时间
	LastLoginIp   string    `gorm:"last_login_ip"`   // 最近一次登录IP地址
	UserLevel     int       `gorm:"user_level"`      // 0 普通用户，1 VIP用户，2 高级VIP用户
	Mobile        string    `gorm:"mobile"`          // 用户手机号码
	Avatar        string    `gorm:"avatar"`          // 用户头像图片
	SessionKey    string    `gorm:"session_key"`     // 微信登录会话KEY
	Deleted       int       `gorm:"deleted"`         // 逻辑删除
	CreatedTime   time.Time `gorm:"created_time"`    // 创建时间
	UpdatedTime   time.Time `gorm:"updated_time"`    // 更新时间
	DeletedTime   time.Time `gorm:"deleted_time"`    // 删除时间
}

// 	Id       uint64 `gorm:"primaryKey"` // 用户编号
// 	OpenId   string `gorm:"open_id"`    // 微信用户唯一标识
// 	Username string `gorm:"username"`   // 用户名称
// 	Password string `gorm:"password"`   // 用户密码
// 	Status   uint   `gorm:"status"`     // 用户状态
// 	Created  string `gorm:"created"`    // 创建时间
// 	Updated  string `gorm:"updated"`    // 更新时间
// }

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
