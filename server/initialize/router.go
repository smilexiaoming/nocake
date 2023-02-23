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

	engine.GET("/app/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 微信小程序API
	app := engine.Group("/app")
	{
		// 用户登录
		app.POST("/login", api.GetAppUser().UserLogin)
		// 商品分类
		app.GET("/category/option", api.GetAppCategory().GetCategoryOption)
		app.GET("/goods/list", api.GetAppGoods().GetGoodList)
	}

	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%s", global.Config.Server.Post))
}
