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

func (a *AppCart) AddCart(c *gin.Context) {
	param := app.CartAddParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}

	if added := a.Add(param); added > 0 {
		response.Success(constant.Created, nil, c)
		return
	}
	response.Error(constant.NotCreated, c)
}

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

func (a *AppCart) GetCartInfo(c *gin.Context) {
	param := app.CartQueryParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if info := a.GetInfo(param); len(info.CartItem) > 0 {
		response.Success(constant.Selected, info, c)
		return
	}
	response.Error(constant.NotSelected, c)
}
