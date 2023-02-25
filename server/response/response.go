package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageResult struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func Error(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    400,
		Message: message,
		Data:    "",
	})
}