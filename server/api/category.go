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

func (c *AppCategory) GetCategoryOption(context *gin.Context) {
	var param app.CategoryQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	option := c.GetOption(param)
	response.Success(constant.Selected, option, context)
}
