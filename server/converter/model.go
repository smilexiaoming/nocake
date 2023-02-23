package model

type T_comment struct {
	Id          int    `:"id"`
	ValueId     int    `:"value_id"`     // 如果type=0，则是商品评论；如果是type=1，则是专题评论。
	CommentType int    `:"comment_type"` // 评论类型，如果type=0，则是商品评论；如果是type=1，则是专题评论；如果type=3，则是订单商品评论。
	Content     string `:"content"`      // 评论内容
	UserId      int    `:"user_id"`      // 用户表的用户ID
	HasPicture  int    `:"has_picture"`  // 是否含有图片
	PicUrls     string `:"pic_urls"`     // 图片地址列表，采用JSON数组格式
	Star        int    `:"star"`         // 评分， 1-5
	Deleted     int    `:"deleted"`      // 逻辑删除
}

type T_goods struct {
	Id           int     `:"id"`
	Name         string  `:"name"`        // 商品名称
	Brief        string  `:"brief"`       // 商品简介
	Detail       string  `:"detail"`      // 商品详细介绍，是富文本格式
	CategoryId   int     `:"category_id"` // 商品所属类目ID
	Gallery      string  `:"gallery"`     // 商品宣传图片列表，采用JSON数组格式
	Keywords     string  `:"keywords"`    // 商品关键字，采用逗号间隔
	IsOnSale     int     `:"is_on_sale"`  // 是否上架
	SortOrder    int     `:"sort_order"`
	PicUrl       string  `:"pic_url"`       // 商品页面商品图片
	ShareUrl     string  `:"share_url"`     // 商品分享朋友圈图片
	IsNew        int     `:"is_new"`        // 是否新品首发，如果设置则可以在新品首发页面展示
	IsHot        int     `:"is_hot"`        // 是否人气推荐，如果设置则可以在人气推荐页面展示
	Unit         string  `:"unit"`          // 商品单位，例如件、盒
	CounterPrice float64 `:"counter_price"` // 专柜价格
	RetailPrice  float64 `:"retail_price"`  // 零售价格
	Deleted      int     `:"deleted"`       // 逻辑删除
}

type T_order struct {
	Id            int     `:"id"`
	UserId        int     `:"user_id"`        // 用户表的用户ID
	OrderSn       string  `:"order_sn"`       // 订单编号
	OrderStatus   int     `:"order_status"`   // 订单状态
	Consignee     string  `:"consignee"`      // 收货人名称
	Mobile        string  `:"mobile"`         // 收货人手机号
	Address       string  `:"address"`        // 收货具体地址
	Message       string  `:"message"`        // 用户订单留言
	GoodsPrice    float64 `:"goods_price"`    // 商品总费用
	FreightPrice  float64 `:"freight_price"`  // 配送费用
	CouponPrice   float64 `:"coupon_price"`   // 优惠券减免
	IntegralPrice float64 `:"integral_price"` // 用户积分减免
	GrouponPrice  float64 `:"groupon_price"`  // 团购优惠价减免
	OrderPrice    float64 `:"order_price"`    // 订单费用， = goods_price + freight_price - coupon_price
	ActualPrice   float64 `:"actual_price"`   // 实付费用， = order_price - integral_price
	PayId         string  `:"pay_id"`         // 微信付款编号
	PayTime       string  `:"pay_time"`       // 微信付款时间
	ShipSn        string  `:"ship_sn"`        // 发货编号
	ShipChannel   string  `:"ship_channel"`   // 发货快递公司
	ShipTime      string  `:"ship_time"`      // 发货开始时间
	ConfirmTime   string  `:"confirm_time"`   // 用户确认收货时间
	Comments      int     `:"comments"`       // 待评价订单商品数量
	EndTime       string  `:"end_time"`       // 订单关闭时间
	Deleted       int     `:"deleted"`        // 逻辑删除
}

type T_user struct {
	Id            int    `:"id"`
	OpenId        string `:"open_id"`         // 微信登录openid
	WxId          string `:"wx_id"`           // 用户原始微信id
	Username      string `:"username"`        // 用户名称
	Nickname      string `:"nickname"`        // 用户昵称
	Password      string `:"password"`        // 用户密码
	Status        int    `:"status"`          // 0 可用, 1 禁用, 2 注销
	Gender        int    `:"gender"`          // 性别：0 未知， 1男， 1 女
	Birthday      string `:"birthday"`        // 生日
	LastLoginTime string `:"last_login_time"` // 最近一次登录时间
	LastLoginIp   string `:"last_login_ip"`   // 最近一次登录IP地址
	UserLevel     int    `:"user_level"`      // 0 普通用户，1 VIP用户，2 高级VIP用户
	Mobile        string `:"mobile"`          // 用户手机号码
	Avatar        string `:"avatar"`          // 用户头像图片
	SessionKey    string `:"session_key"`     // 微信登录会话KEY
	Deleted       int    `:"deleted"`         // 逻辑删除
	CreatedOn     string `:"created_on"`      // 创建时间
	UpdatedOn     string `:"updated_on"`      // 更新时间
	DeletedOn     string `:"deleted_on"`      // 删除时间
}

type T_address struct {
	Id            int    `:"id"`
	AddressName   string `:"address_name"`   // 收货人名称
	UserId        int    `:"user_id"`        // 用户表的用户ID
	Province      string `:"province"`       // 行政区域表的省ID
	City          string `:"city"`           // 行政区域表的市ID
	County        string `:"county"`         // 行政区域表的区县ID
	AddressDetail string `:"address_detail"` // 详细收货地址
	AreaCode      string `:"area_code"`      // 地区编码
	PostalCode    string `:"postal_code"`    // 邮政编码
	Tel           string `:"tel"`            // 手机号码
	IsDefault     int    `:"is_default"`     // 是否默认地址
	Deleted       int    `:"deleted"`        // 逻辑删除
}

type T_cart struct {
	Id             int     `:"id"`
	UserId         int     `:"user_id"`        // 用户表的用户ID
	GoodsId        int     `:"goods_id"`       // 商品表的商品ID
	GoodsSn        string  `:"goods_sn"`       // 商品编号
	GoodsName      string  `:"goods_name"`     // 商品名称
	ProductId      int     `:"product_id"`     // 商品货品表的货品ID
	Price          float64 `:"price"`          // 商品货品的价格
	CartNumber     int     `:"cart_number"`    // 商品货品的数量
	Specifications string  `:"specifications"` // 商品规格值列表，采用JSON数组格式
	Checked        int     `:"checked"`        // 购物车中商品是否选择状态
	PicUrl         string  `:"pic_url"`        // 商品图片或者商品货品图片
	Deleted        int     `:"deleted"`        // 逻辑删除
}

type T_category struct {
	Id        int    `:"id"`
	Name      string `:"name"`       // 类目名称
	Keywords  string `:"keywords"`   // 类目关键字，以JSON数组格式
	Desc      string `:"desc"`       // 类目广告语介绍
	Pid       int    `:"pid"`        // 父类目ID
	IconUrl   string `:"icon_url"`   // 类目图标
	PicUrl    string `:"pic_url"`    // 类目图片
	Level     int    `:"level"`      // 类目层级
	SortOrder int    `:"sort_order"` // 排序
	Deleted   int    `:"deleted"`    // 逻辑删除
}
