package controller

import (
	"net/http"
	route "nocake/http_server/routers"

	"github.com/gin-gonic/gin"
)

//这里每个controller执行init方法都要注册自动路由
func init() {
	route.Register(&User{})
}

type User struct {
	
}

//控制器的方法 分页查询
func (api *User) Pages_get(c *gin.Context) { //,httpMethod string
	users := []int{1, 2, 3}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": users,
	})

}
func (api *User) AddUser(c *gin.Context) { //,httpMethod string
	users := []int{2, 3, 6}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": users,
	})

}
