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
// @Accept  json
// @Produce  json
// @Param categoryid query int true "categoryid"
// @Param kewords query string true "kewords"
// @Param name query string true "name"
// @Success 200 {object} app.GoodsList "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /category/option [get]
func (g *AppGoods) GetGoodList(c *gin.Context) {
	param := app.GoodsQueryParam{}
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}
	goodsList := g.GetList(param)
	response.Success(constant.Selected, goodsList, c)
}
