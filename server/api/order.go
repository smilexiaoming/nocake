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
