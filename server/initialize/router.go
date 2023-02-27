package initialize

import (
	"fmt"
	"nocake/api"
	_ "nocake/docs"
	"nocake/global"

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
	engine.Use(gin.Recovery())

	engine.GET("/app/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// engine.SetTrustedProxies([]string{"127.0.0.1"})
	// 微信小程序API
	app := engine.Group("/app")
	{
		// 用户登录
		app.POST("/login", api.GetAppUser().UserLogin)

		// 商品分类
		app.GET("/category/option", api.GetAppCategory().GetCategoryOption)

		// 商品
		app.GET("/goods/list", api.GetAppGoods().GetGoodList)
		app.GET("/goods/detail", api.GetAppGoods().GetGoodsDetail)
		app.GET("/goods/search", api.GetAppGoods().SearchGoods)

		// 购物车
		app.POST("/cart/add", api.GetAppCart().AddCart)
		app.POST("/cart/set", api.GetAppCart().SetCart)
		app.DELETE("/cart/delete", api.GetAppCart().DeleteCart)
		app.DELETE("/cart/clear", api.GetAppCart().ClearCart)
		app.GET("/cart/query", api.GetAppCart().GetCartInfo)

		// 订单
		app.POST("/order/submit", api.GetAppOrder().SubmitOrder)
		app.POST("/order/update", api.GetAppOrder().UpdateOrder)
		app.GET("/order/list", api.GetAppOrder().GetOrderList)
		app.GET("/order/detail", api.GetAppOrder().GetOrderDetail)

		// 地址管理
		app.POST("/address/submit", api.GetAppAddress().AddAddress)
		app.POST("/address/update", api.GetAppAddress().UpdateAddress)
		app.DELETE("/address/delete", api.GetAppAddress().DeleteAddress)
		app.GET("/address/list", api.GetAppAddress().GetAddressList)

	}

	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%s", global.Config.Server.Post))
}
