package initialize

import (
	"fmt"
	"nocake/api"
	"nocake/global"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	engine := gin.Default()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 微信小程序API
	app := engine.Group("/app")
	{
		// 用户登录
		app.POST("/login", api.GetAppUser().UserLogin)
		// 商品分类
		app.GET("/category/option", api.GetAppCategory().GetCategoryOption)
	}

	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%s", global.Config.Server.Post))
}
