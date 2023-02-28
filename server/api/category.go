package api

import (
	"nocake/constant"
	"nocake/models/app"
	"nocake/response"
	"nocake/service"

	"github.com/gin-gonic/gin"
)

type AppCategory struct {
	service.AppCategoryService
}

func GetAppCategory() *AppCategory {
	return &AppCategory{}
}

// @Summary 获取品类列表
// @Description 传入level 和 pid
// @Accept  multipart/form-data
// @Produce  json
// @Param level query string false "level"
// @Param pid query string false "pid"
// @Success 200 {object} app.CategoryOption "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/category/option [get]
func (g *AppCategory) GetCategoryOption(c *gin.Context) {
	param := app.CategoryQueryParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	option, errMessage := g.GetOption(param)
	if errMessage != "" {
		response.Error(errMessage, c)
	}
	response.Success(constant.Selected, option, c)
}
