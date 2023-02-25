package app

import "time"

type Comment struct {
	Id          int       `gorm:"id"`
	ValueId     int       `gorm:"value_id"`     // 如果type=0，则是商品评论；如果是type=1，则是专题评论。
	Commentype  int       `gorm:"comment_type"` // 评论类型，如果type=0，则是商品评论；如果是type=1，则是专题评论；如果type=3，则是订单商品评论。
	Content     string    `gorm:"content"`      // 评论内容
	OpenId      string    `gorm:"open_id"`      // 用户表的用户ID
	HasPicture  int       `gorm:"has_picture"`  // 是否含有图片
	PicUrls     string    `gorm:"pic_urls"`     // 图片地址列表，采用JSON数组格式
	Star        int       `gorm:"star"`         // 评分， 1-5
	Deleted     int       `gorm:"deleted"`      // 逻辑删除
	CreatedTime time.Time `gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time"` // 删除时间
}
