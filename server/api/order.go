package api

import (
	"nocake/constant"
	"nocake/models/app"
	"nocake/response"
	"nocake/service"

	"github.com/gin-gonic/gin"
)

type AppOrder struct {
	service.AppOrderService
}

func GetAppOrder() *AppOrder {
	return &AppOrder{}
}

func (o *AppOrder) SubmitOrder(c *gin.Context) {
	param := app.OrderSubmitParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if rowsAffected := o.Submit(param); rowsAffected > 0 {
		response.Success(constant.Submited, nil, c)
		return
	}
	response.Error(constant.NotSubmited, c)
}

func (o *AppOrder) UpdateOrder(c *gin.Context) {
	param := app.OrderUpdateParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if rowsAffected := o.Update(param); rowsAffected > 0 {
		response.Success(constant.Submited, nil, c)
		return
	}
	response.Error(constant.NotSubmited, c)
}

func (o *AppOrder) GetOrderList(c *gin.Context) {
	param := app.OrderQueryListParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	orderList := o.GetList(param)
	if len(orderList) > 0 {
		response.Success(constant.Selected, orderList, c)
		return
	}
	response.Error(constant.NotSelected, c)
}

func (o *AppOrder) GetOrderDetail(c *gin.Context) {
	param := app.OrderQueryDetailParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	orderDetail := o.GetDetail(param)
	response.Success(constant.Selected, orderDetail, c)
	return
}
