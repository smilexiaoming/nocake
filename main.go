package main

import (
	"fmt"
	route "nocake/http_server/routers"
	_ "nocake/http_server/routers/api"
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
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	r := route.InitRouter()
	r.Run(port) // listen and serve on 0.0.0.0:8080
}
