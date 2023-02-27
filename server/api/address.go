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
