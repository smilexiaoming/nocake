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
// @Description 传入level
// @Accept  json
// @Produce  json
// @Param level query string true "level"
// @Success 200 {object} app.CategoryOption "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /category/option [get]
func (c *AppCategory) GetCategoryOption(context *gin.Context) {
	var param app.CategoryQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	option, errMessage := c.GetOption(param)
	if errMessage != "" {
		response.Error(errMessage, context)
	}
	response.Success(constant.Selected, option, context)
}
