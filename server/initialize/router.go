package initialize

import (
	"fmt"
	"net/http"
	"nocake/api"
	_ "nocake/docs"
	"nocake/global"
	"nocake/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /app

// @securityDefinitions.basic  BasicAuth
func Router() {
	engine := gin.Default()
	// 开启跨域
	engine.Use(middleware.Cors())

	// ip来源校验
	engine.SetTrustedProxies([]string{"127.0.0.1"})

	// 404
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)

	// swagger文档
	engine.GET("/app/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	web := engine.Group("/web")
	{
		// 登录
		web.GET("/captcha", api.GetWebUser().GetCaptcha)
		web.POST("/login", api.GetWebUser().UserLogin)

		// JWT认证
		web.Use(middleware.JwtAuth())

		// 文件上传
		web.POST("/upload", api.GetWebFileUpload().FileUpload)

		// 商品管理
		web.POST("/goods/create", api.GetWebGoods().CreateGoods)
		web.DELETE("/goods/delete", api.GetWebGoods().DeleteGoods)
		web.PUT("/goods/update", api.GetWebGoods().UpdateGoods)
		web.PUT("/goods/status/update", api.GetWebGoods().UpdateGoodsStatus)
		web.GET("/goods/list", api.GetWebGoods().GetGoodsList)

		// 订单管理
		web.DELETE("/order/delete", api.GetWebOrder().DeleteOrder)
		web.PUT("/order/update", api.GetWebOrder().UpdateOrder)
		web.GET("/order/list", api.GetWebOrder().GetOrderList)
		web.GET("/order/detail", api.GetWebOrder().GetOrderDetail)

		// 类目管理
		web.POST("/category/create", api.GetWebCategory().CreateCategory)
		web.DELETE("/category/delete", api.GetWebCategory().DeleteCategory)
		web.PUT("/category/update", api.GetWebCategory().UpdateCategory)
		web.GET("/category/list", api.GetWebCategory().GetCategoryList)
		web.GET("/category/option", api.GetWebCategory().GetCategoryOption)

		// 数据统计
		web.GET("/today/data", api.GetWebStatistics().GetTodayData)
		web.GET("/order/data", api.GetWebStatistics().GetOrderData)
		web.GET("/shop/data", api.GetWebStatistics().GetShopData)
	}
	// 微信小程序API
	app := engine.Group("/app")
	{
		// 用户登录
		app.POST("/login", api.GetAppUser().UserLogin)

		app.Use(middleware.AppJwtAuth())

		// 商品分类
		app.GET("/category/option", api.GetAppCategory().GetCategoryOption)

		// 商品
		app.GET("/goods/list", api.GetAppGoods().GetGoodList)
		app.GET("/goods/detail", api.GetAppGoods().GetGoodsDetail)
		app.GET("/goods/search", api.GetAppGoods().SearchGoods)
		app.GET("/goods/hot", api.GetAppGoods().GetGoodsHot)

		// 购物车
		app.POST("/cart/set", api.GetAppCart().SetCart)
		app.POST("/cart/update", api.GetAppCart().UpdateCart)
		app.DELETE("/cart/delete", api.GetAppCart().DeleteCart)
		app.DELETE("/cart/clear", api.GetAppCart().ClearCart)
		app.GET("/cart/query", api.GetAppCart().GetCartInfo)

		// 订单
		app.POST("/order/submit", api.GetAppOrder().SubmitOrder)
		app.POST("/order/update", api.GetAppOrder().UpdateOrder)
		app.GET("/order/list", api.GetAppOrder().GetOrderList)

		// 地址管理
		app.POST("/address/submit", api.GetAppAddress().AddAddress)
		app.POST("/address/update", api.GetAppAddress().UpdateAddress)
		app.DELETE("/address/delete", api.GetAppAddress().DeleteAddress)
		app.GET("/address/list", api.GetAppAddress().GetAddressList)
		app.GET("/address/detail", api.GetAppAddress().GetAddressDetail)

	}

	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%s", global.Config.Server.Post))
}
