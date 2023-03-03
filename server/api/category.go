package api

import (
	"nocake/constant"
	"nocake/models/app"
	"nocake/models/web"
	"nocake/response"
	"nocake/service"

	"github.com/gin-gonic/gin"
)

type WebCategory struct {
	service.WebCategoryService
}

type AppCategory struct {
	service.AppCategoryService
}

func GetWebCategory() *WebCategory {
	return &WebCategory{}
}

func GetAppCategory() *AppCategory {
	return &AppCategory{}
}

func (c *WebCategory) CreateCategory(context *gin.Context) {
	var param web.CategoryCreateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	if count := c.Create(param); count > 0 {
		response.Success(constant.Created, count, context)
		return
	}
	response.Error(constant.NotCreated, context)
}

func (c *WebCategory) DeleteCategory(context *gin.Context) {
	var param web.CategoryDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	if count := c.Delete(param); count > 0 {
		response.Success(constant.Deleted, count, context)
		return
	}
	response.Error(constant.NotDeleted, context)
}

func (c *WebCategory) UpdateCategory(context *gin.Context) {
	var param web.CategoryUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	if count := c.Update(param); count > 0 {
		response.Success(constant.Updated, count, context)
		return
	}
	response.Error(constant.NotUpdated, context)
}

func (c *WebCategory) GetCategoryList(context *gin.Context) {
	var param web.CategoryQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	categoryList := c.GetList(param)
	response.Success(constant.Selected, categoryList, context)
}

func (c *WebCategory) GetCategoryOption(context *gin.Context) {
	option := c.GetOption()
	response.Success(constant.Selected, option, context)
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
