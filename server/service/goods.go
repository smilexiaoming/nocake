package service

import (
	"fmt"
	"nocake/global"
	"nocake/models/app"
)

type AppGoodsService struct {
}

func (g *AppGoodsService) GetList(param app.GoodsQueryParam) []app.GoodsList {
	goodsList := make([]app.GoodsList, 0)
	Db := global.Db.Table("t_goods")
	fmt.Println("param :", param)
	if param.CategoryId != 0 {
		Db.Where("category_id = ?", param.CategoryId)
	}
	if param.Keywords != "" {
		Db.Where("keywords = ?", param.Keywords)
	}
	if param.Name != "" {
		Db.Where("name like ?", "%"+param.Name+"%")
	}
	Db.Find(&goodsList)
	return goodsList
}
