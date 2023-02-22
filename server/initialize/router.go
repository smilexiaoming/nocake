package initialize

import "github.com/gin-gonic/gin"

func Router() {
	engine := gin.Default()

	// 微信小程序API
	app := engine.Group("/app")
	{
		app.POST("/login")
	}
}
