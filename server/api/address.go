package api

import (
	"nocake/constant"
	"nocake/models/app"
	"nocake/response"
	"nocake/service"

	"github.com/gin-gonic/gin"
)

type AppAddress struct {
	service.AppAddressService
}

func GetAppAddress() *AppAddress {
	return &AppAddress{}
}

// @Summary 新增地址
// @Description 
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData string true "收货人名称"
// @Param open_id formData string true "用户表的用户ID"
// @Param province formData string true "行政区域表的省ID"
// @Param city formData string true "行政区域表的市ID"
// @Param county formData string true "行政区域表的区县ID"
// @Param detail formData string true "详细收货地址"
// @Param area_code formData string true "地区编码"
// @Param postal_code formData string true "邮政编码"
// @Param tel formData string true "手机号码"
// @Param is_default formData int true "是否默认地址"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/address/add [post]
func (a *AppAddress) AddAddress(c *gin.Context) {
	param := app.AddressAddParam{}
	if err := c.Bind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if rowsAffect := a.Add(param); rowsAffect > 0 {
		response.Success(constant.Created, nil, c)
		return
	}
	response.Error(constant.Created, c)
}

// @Summary 删除地址
// @Description 
// @Accept  multipart/form-data
// @Produce  json
// @Param address_id formData int true "地址id"
// @Param open_id formData string true "open_id"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/address/delete [delete]
func (a *AppAddress) DeleteAddress(c *gin.Context) {
	param := app.AddressDeleteParam{}
	if err := c.Bind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if rowsAffect := a.Delete(param); rowsAffect > 0 {
		response.Success(constant.Deleted, nil, c)
		return
	}
	response.Error(constant.NotDeleted, c)
}

// @Summary 更改地址
// @Description 
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData string true "收货人名称"
// @Param open_id formData string true "用户表的用户ID"
// @Param province formData string true "行政区域表的省ID"
// @Param city formData string true "行政区域表的市ID"
// @Param county formData string true "行政区域表的区县ID"
// @Param detail formData string true "详细收货地址"
// @Param area_code formData string true "地区编码"
// @Param postal_code formData string true "邮政编码"
// @Param tel formData string true "手机号码"
// @Param is_default formData int true "是否默认地址"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/address/update [post]
func (a *AppAddress) UpdateAddress(c *gin.Context) {
	param := app.AddressUpdateParam{}
	if err := c.Bind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if rowsAffect := a.Updata(param); rowsAffect > 0 {
		response.Success(constant.Updated, nil, c)
		return
	}
	response.Error(constant.NotUpdated, c)
}

// @Summary 获取地址列表
// @Description
// @Accept  multipart/form-data
// @Produce  json
// @Param open_id formData string true "open_id"
// @Success 200 {object} app.Address "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/address/list [get]
func (a *AppAddress) GetAddressList(c *gin.Context) {
	param := app.AddressListQueryParam{}
	if err := c.Bind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	if addressInfo := a.GetList(param); len(addressInfo) > 0 {
		response.Success(constant.Selected, addressInfo, c)
		return
	}
	response.Error(constant.NotSelected, c)
}
