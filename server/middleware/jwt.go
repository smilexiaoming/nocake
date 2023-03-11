package middleware

import (
	"fmt"
	"nocake/common"
	"nocake/response"

	"github.com/gin-gonic/gin"
)

// JwtAuth JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Error("未登录或非法访问", c)
			c.Abort()
			return
		}
		if err := common.VerifyToken(token); err != nil {
			response.Error("登录已过期，请重新登录", c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func AppJwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		skey := c.Request.Header.Get("skey")
		fmt.Printf("skey: %v\n", skey)
		if skey == "" {
			response.Error("未登录或非法访问", c)
			c.Abort()
			return
		}
		if err := common.VerifySkey(skey); err != nil {
			response.Error("登录已过期，请重新登录", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
