package main

import (
	_ "nocake/http_server/controllers"
	route "nocake/http_server/routers"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r := route.InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
