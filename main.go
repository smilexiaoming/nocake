package main

import (
	_ "nocake/http_server/controllers"
	route "nocake/http_server/routers"
	"nocake/models"
	"nocake/pkg/logging"
	"nocake/pkg/setting"

	"nocake/pkg/gredis"
	"nocake/pkg/util"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {
	r := route.InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
