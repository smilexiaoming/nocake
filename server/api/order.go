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

type WebOrder struct {
	service.WebOrderService
}

type AppOrder struct {
	service.AppOrderService
}

func GetWebOrder() *WebOrder {
	return &WebOrder{}
}

func GetAppOrder() *AppOrder {
	return &AppOrder{}
}

func (o *WebOrder) DeleteOrder(context *gin.Context) {
	var param web.OrderDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	if count := o.Delete(param); count > 0 {
		response.Success(constant.Deleted, count, context)
		return
	}
	response.Error(constant.NotDeleted, context)
}

func (o *WebOrder) UpdateOrder(context *gin.Context) {
	var param web.OrderUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	if count := o.Update(param); count > 0 {
		response.Success("更新成功", count, context)
		return
	}
	response.Error("更新失败", context)
}

func (o *WebOrder) GetOrderList(context *gin.Context) {
	var param web.OrderListParam
	if err := context.ShouldBind(&param); err != nil {
		fmt.Println(err)
		response.Error(constant.ParamInvalid, context)
		return
	}
	productList, rows := o.GetList(param)
	response.SuccessPage("查询成功", productList, rows, context)
}

func (o *WebOrder) GetOrderDetail(context *gin.Context) {
	var param web.OrderDetailParam
	if err := context.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, context)
		return
	}
	productDetail := o.GetDetail(param)
	response.Success("查询成功", productDetail, context)
}

// @Summary 提交订单
// @Description 从购物车中提交订单
// @Accept  multipart/form-data
// @Produce  json
// @Param open_id formData string true "open_id"
// @Param messge formData string true "messge 备注"
// @Param address_id formData string true "选择地址的id"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/order/submit [post]
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

// @Summary 修改订单状态
// @Description
// @Accept  multipart/form-data
// @Param open_id formData string true "open_id"
// @Param order_id formData int true "订单号"
// @Param status formData int true "订单状态"
// @Produce  json
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/order/update [post]
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

// @Summary 获取订单列表
// @Description
// @Accept  multipart/form-data
// @Param open_id formData string true "open_id"
// @Param status formData int true "订单状态"
// @Param page_size formData int true "分页参数"
// @Param page_num formData int true "分页参数"
// @Produce  json
// @Success 200 {array} app.Order "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/order/list [get]
func (o *AppOrder) GetOrderList(c *gin.Context) {
	param := app.OrderQueryListParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	fmt.Printf("param: %v\n", param)
	orderList := o.GetList(param)
	if len(orderList) > 0 {
		response.Success(constant.Selected, orderList, c)
		return
	}
	response.Error(constant.NotSelected, c)
}
