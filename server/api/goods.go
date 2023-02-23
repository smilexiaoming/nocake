package api

import (
	"nocake/constant"
	"nocake/models/app"
	"nocake/response"
	"nocake/service"

	"github.com/gin-gonic/gin"
)

type AppGoods struct {
	service.AppGoodsService
}

func GetAppGoods() *AppGoods {
	return &AppGoods{}
}

// @Summary 获取商品列表
// @Description 传入 categoryid | kewords | name
// @Accept  multipart/form-data
// @Param categoryid query int true "category_id"
// @Success 200 {object} app.GoodsList "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/goods/list [get]
func (g *AppGoods) GetGoodList(c *gin.Context) {
	param := app.GoodsListQueryParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	goodList := g.GetList(param)
	response.Success(constant.Selected, goodList, c)
}

// @Summary 获取商品详情
// @Description 传入 id
// @Accept  multipart/form-data
// @Param id query int true "id"
// @Success 200 {object} app.GoodsDetail "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/goods/detail [get]
func (g *AppGoods) GetGoodsDetail(c *gin.Context) {
	param := app.GoodsDetailQueryParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
	}
	goodDetail := g.GetDetail(param)
	response.Success(constant.Selected, goodDetail, c)
}

// @Summary 搜索商品
// @Description 传入   name | kewords
// @Accept  multipart/form-data
// @Param name query string false "name"
// @Param keywords query string false "keywords"
// @Success 200 {object} app.GoodsList "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/goods/search [get]
func (g *AppGoods) SearchGoods(c *gin.Context) {
	param := app.GoodsSearchQueryParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	goodList := g.Search(param)
	response.Success(constant.Selected, goodList, c)
}
