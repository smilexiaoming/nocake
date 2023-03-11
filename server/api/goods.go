package api

import (
	"fmt"
	"nocake/constant"
	"nocake/models/app"
	"nocake/models/web"
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

type WebGoods struct {
	service.WebGoodsService
}

func GetWebGoods() *WebGoods {
	return &WebGoods{}
}

func (g *WebGoods) CreateGoods(c *gin.Context) {
	var param web.GoodsCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if count := g.Create(param); count > 0 {
		response.Success(constant.Created, count, c)
		return
	}
	response.Error(constant.NotCreated, c)
}

func (g *WebGoods) DeleteGoods(c *gin.Context) {
	var param web.GoodsDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if count := g.Delete(param); count > 0 {
		response.Success(constant.Deleted, count, c)
		return
	}
	response.Error(constant.NotDeleted, c)
}

func (g *WebGoods) UpdateGoods(c *gin.Context) {
	var param web.GoodsUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	fmt.Printf("param: %#v\n", param)
	if count := g.Update(param); count > 0 {
		response.Success(constant.Updated, count, c)
		return
	}
	response.Error(constant.NotUpdated, c)
}

func (g *WebGoods) UpdateGoodsStatus(c *gin.Context) {
	var param web.GoodsStatusUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if count := g.UpdateStatus(param); count > 0 {
		response.Success(constant.Updated, count, c)
		return
	}
	response.Error(constant.NotUpdated, c)
}

func (g *WebGoods) GetGoodsList(c *gin.Context) {
	var param web.GoodsListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	goodsList, rows := g.GetList(param)
	response.SuccessPage(constant.Selected, goodsList, rows, c)
}

// @Summary 获取商品列表
// @Description 
// @Accept  multipart/form-data
// @Success 200 {object} app.GoodsList "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/goods/hot [get]
func (g *AppGoods) GetGoodsHot(c *gin.Context) {
	goodList := g.GetHot()
	if len(goodList) == 0 {
		response.Error(constant.NotSelected, c)
		return
	}
	response.Success(constant.Selected, goodList, c)
}

// @Summary 获取商品列表
// @Description 传入 category_id | kewords | name
// @Accept  multipart/form-data
// @Param category_id query int true "category_id"
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
	if len(goodList) == 0 {
		response.Error(constant.NotSelected, c)
		return
	}
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
		return
	}
	goodDetail := g.GetDetail(param)
	if goodDetail.Id == 0 {
		response.Error(constant.NotSelected, c)
		return
	}
	response.Success(constant.Selected, goodDetail, c)
}

// @Summary 搜索商品
// @Description 传入   name | kewords
// @Accept  multipart/form-data
// @Param name query string false "商品名称"
// @Param keywords query string false "关键字"
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
	if len(goodList) == 0 {
		response.Error(constant.NotSelected, c)
		return
	}
	response.Success(constant.Selected, goodList, c)
}
