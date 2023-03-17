package api

import (
	"nocake/constant"
	"nocake/models/app"
	"nocake/response"
	"nocake/service"

	"github.com/gin-gonic/gin"
)

type AppCart struct {
	service.AppCartService
}

func GetAppCart() *AppCart {
	return &AppCart{}
}

// @Summary 更改购物车商品数量
// @Description 传入code进行鉴权
// @Accept  multipart/form-data
// @Produce  json
// @Param open_id formData string true "open_id"
// @Param goods_id formData int true "商品id"
// @Param cart_number formData int true "设置的数量"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/cart/set [post]
func (a *AppCart) SetCart(c *gin.Context) {
	param := app.CartSetParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}

	if added := a.Set(param); added {
		response.Success(constant.Created, nil, c)
		return
	}
	response.Error(constant.NotCreated, c)
}


// @Summary 删除购物车商品
// @Description
// @Accept  multipart/form-data
// @Produce  json
// @Param open_id formData string true "open_id"
// @Param goods_id formData int true "商品id"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/cart/delete [delete]
func (a *AppCart) DeleteCart(c *gin.Context) {
	param := app.CartDeleteParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if deleted := a.Delete(param); deleted > 0 {
		response.Success(constant.Deleted, nil, c)
		return
	}
	response.Error(constant.NotDeleted, c)
}

// @Summary 清空购物车
// @Description
// @Accept  multipart/form-data
// @Produce  json
// @Param open_id formData string true "open_id"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/cart/clear [delete]
func (a *AppCart) ClearCart(c *gin.Context) {
	param := app.CartClearParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}

	if deleted := a.Clear(param); deleted > 0 {
		response.Success(constant.Cleared, nil, c)
		return
	}
	response.Error(constant.NotCleared, c)
}

// @Summary 查看购物车
// @Description
// @Accept  multipart/form-data
// @Produce  json
// @Param open_id formData string true "open_id"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/cart/query [get]
func (a *AppCart) GetCartInfo(c *gin.Context) {
	param := app.CartQueryParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	info := a.GetInfo(param)
	response.Success(constant.Selected, info, c)
}
