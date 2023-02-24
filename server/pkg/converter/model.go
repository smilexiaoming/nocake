package model

type Cart struct {
	Id             int     `gorm:"id"`
	UserId         int     `gorm:"user_id"`        // 用户表的用户ID
	GoodsId        int     `gorm:"goods_id"`       // 商品表的商品ID
	GoodsName      string  `gorm:"goods_name"`     // 商品名称
	Price          float64 `gorm:"price"`          // 商品货品的价格
	Carnumber      int     `gorm:"cart_number"`    // 商品货品的数量
	Specifications string  `gorm:"specifications"` // 商品规格值列表，采用JSON数组格式
	Checked        int     `gorm:"checked"`        // 购物车中商品是否选择状态
	PicUrl         string  `gorm:"pic_url"`        // 商品图片或者商品货品图片
	Deleted        int     `gorm:"deleted"`        // 逻辑删除
}

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

type Order struct {
	Id            int     `gorm:"id"`
	UserId        int     `gorm:"user_id"`        // 用户表的用户ID
	OrderSn       string  `gorm:"order_sn"`       // 订单编号
	OrderStatus   int     `gorm:"order_status"`   // 订单状态
	Consignee     string  `gorm:"consignee"`      // 收货人名称
	Mobile        string  `gorm:"mobile"`         // 收货人手机号
	Address       string  `gorm:"address"`        // 收货具体地址
	Message       string  `gorm:"message"`        // 用户订单留言
	GoodsPrice    float64 `gorm:"goods_price"`    // 商品总费用
	Freighprice   float64 `gorm:"freight_price"`  // 配送费用
	CouponPrice   float64 `gorm:"coupon_price"`   // 优惠券减免
	IntegralPrice float64 `gorm:"integral_price"` // 用户积分减免
	GrouponPrice  float64 `gorm:"groupon_price"`  // 团购优惠价减免
	OrderPrice    float64 `gorm:"order_price"`    // 订单费用， = goods_price + freight_price - coupon_price
	ActualPrice   float64 `gorm:"actual_price"`   // 实付费用， = order_price - integral_price
	PayId         string  `gorm:"pay_id"`         // 微信付款编号
	PayTime       string  `gorm:"pay_time"`       // 微信付款时间
	ShipSn        string  `gorm:"ship_sn"`        // 发货编号
	ShipChannel   string  `gorm:"ship_channel"`   // 发货快递公司
	ShipTime      string  `gorm:"ship_time"`      // 发货开始时间
	ConfirmTime   string  `gorm:"confirm_time"`   // 用户确认收货时间
	Comments      int     `gorm:"comments"`       // 待评价订单商品数量
	EndTime       string  `gorm:"end_time"`       // 订单关闭时间
	Deleted       int     `gorm:"deleted"`        // 逻辑删除
}

type Address struct {
	Id            int    `gorm:"id"`
	AddressName   string `gorm:"address_name"`   // 收货人名称
	UserId        int    `gorm:"user_id"`        // 用户表的用户ID
	Province      string `gorm:"province"`       // 行政区域表的省ID
	City          string `gorm:"city"`           // 行政区域表的市ID
	County        string `gorm:"county"`         // 行政区域表的区县ID
	AddressDetail string `gorm:"address_detail"` // 详细收货地址
	AreaCode      string `gorm:"area_code"`      // 地区编码
	PostalCode    string `gorm:"postal_code"`    // 邮政编码
	Tel           string `gorm:"tel"`            // 手机号码
	IsDefault     int    `gorm:"is_default"`     // 是否默认地址
	Deleted       int    `gorm:"deleted"`        // 逻辑删除
}

type Comment struct {
	Id         int    `gorm:"id"`
	ValueId    int    `gorm:"value_id"`     // 如果type=0，则是商品评论；如果是type=1，则是专题评论。
	Commentype int    `gorm:"comment_type"` // 评论类型，如果type=0，则是商品评论；如果是type=1，则是专题评论；如果type=3，则是订单商品评论。
	Content    string `gorm:"content"`      // 评论内容
	UserId     int    `gorm:"user_id"`      // 用户表的用户ID
	HasPicture int    `gorm:"has_picture"`  // 是否含有图片
	PicUrls    string `gorm:"pic_urls"`     // 图片地址列表，采用JSON数组格式
	Star       int    `gorm:"star"`         // 评分， 1-5
	Deleted    int    `gorm:"deleted"`      // 逻辑删除
}

type Goods struct {
	Id           int     `gorm:"id"`
	Name         string  `gorm:"name"`        // 商品名称
	Brief        string  `gorm:"brief"`       // 商品简介
	Detail       string  `gorm:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId   int     `gorm:"category_id"` // 商品所属类目ID
	Gallery      string  `gorm:"gallery"`     // 商品宣传图片列表，采用JSON数组格式
	Keywords     string  `gorm:"keywords"`    // 商品关键字，采用逗号间隔
	IsOnSale     int     `gorm:"is_on_sale"`  // 是否上架
	Sororder     int     `gorm:"sort_order"`
	PicUrl       string  `gorm:"pic_url"`       // 商品页面商品图片
	ShareUrl     string  `gorm:"share_url"`     // 商品分享朋友圈图片
	IsNew        int     `gorm:"is_new"`        // 是否新品首发，如果设置则可以在新品首发页面展示
	IsHot        int     `gorm:"is_hot"`        // 是否人气推荐，如果设置则可以在人气推荐页面展示
	Unit         string  `gorm:"unit"`          // 商品单位，例如件、盒
	CounterPrice float64 `gorm:"counter_price"` // 专柜价格
	RetailPrice  float64 `gorm:"retail_price"`  // 零售价格
	Deleted      int     `gorm:"deleted"`       // 逻辑删除
}

type GoodsProduct struct {
	Id             int     `gorm:"id"`
	GoodsId        int     `gorm:"goods_id"`       // 商品表的商品ID
	Specifications string  `gorm:"specifications"` // 商品规格值列表，采用JSON数组格式
	Price          float64 `gorm:"price"`          // 商品货品价格
	Producnumber   int     `gorm:"product_number"` // 商品货品数量
	Url            string  `gorm:"url"`            // 商品货品图片
	Deleted        int     `gorm:"deleted"`        // 逻辑删除
}

type OrderGoods struct {
	Id             int     `gorm:"id"`
	OrderId        int     `gorm:"order_id"`       // 订单表的订单ID
	GoodsId        int     `gorm:"goods_id"`       // 商品表的商品ID
	GoodsName      string  `gorm:"goods_name"`     // 商品名称
	Producid       int     `gorm:"product_id"`     // 商品货品表的货品ID
	GoodsNumber    int     `gorm:"goods_number"`   // 商品货品的购买数量
	Price          float64 `gorm:"price"`          // 商品货品的售价
	Specifications string  `gorm:"specifications"` // 商品货品的规格列表
	PicUrl         string  `gorm:"pic_url"`        // 商品货品图片或者商品图片
	GoodsComment   int     `gorm:"goods_comment"`  // 订单商品评论，如果是-1，则超期不能评价；如果是0，则可以评价；如果其他值，则是comment表里面的评论ID。
	Deleted        int     `gorm:"deleted"`        // 逻辑删除
}

type User struct {
	Id           int    `gorm:"id"`
	OpenId       string `gorm:"open_id"`         // 微信登录openid
	WxId         string `gorm:"wx_id"`           // 用户原始微信id
	Username     string `gorm:"username"`        // 用户名称
	Nickname     string `gorm:"nickname"`        // 用户昵称
	Password     string `gorm:"password"`        // 用户密码
	Status       int    `gorm:"status"`          // 0 可用, 1 禁用, 2 注销
	Gender       int    `gorm:"gender"`          // 性别：0 未知， 1男， 1 女
	Birthday     string `gorm:"birthday"`        // 生日
	LasloginTime string `gorm:"last_login_time"` // 最近一次登录时间
	LasloginIp   string `gorm:"last_login_ip"`   // 最近一次登录IP地址
	UserLevel    int    `gorm:"user_level"`      // 0 普通用户，1 VIP用户，2 高级VIP用户
	Mobile       string `gorm:"mobile"`          // 用户手机号码
	Avatar       string `gorm:"avatar"`          // 用户头像图片
	SessionKey   string `gorm:"session_key"`     // 微信登录会话KEY
	Deleted      int    `gorm:"deleted"`         // 逻辑删除
	CreatedTime  string `gorm:"created_time"`    // 创建时间
	UpdatedTime  string `gorm:"updated_time"`    // 更新时间
	DeletedTime  string `gorm:"deleted_time"`    // 删除时间
}
