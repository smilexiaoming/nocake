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

// SuccessPage 请求成功返回分页结果
func SuccessPage(message string, data interface{}, rows int64, c *gin.Context) {
	page := &PageResult{Total: rows, List: data}
	c.JSON(http.StatusOK, Response{200, message, page})
}
