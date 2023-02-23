package service

import (
	"nocake/global"
	"nocake/models/app"
)

type AppGoodsService struct {
}

func (g *AppGoodsService) GetList(param app.GoodsListQueryParam) []app.GoodsList {
	goodsList := make([]app.GoodsList, 0)
	global.Db.Debug().Table("t_goods").Where("category_id = ?", param.CategoryId).Find(&goodsList)
	return goodsList
}

func (g *AppGoodsService) GetDetail(param app.GoodsDetailQueryParam) app.GoodsDetail {
	goodsDetail := *&app.GoodsDetail{}
	global.Db.Debug().Table("t_goods").Where("is_on_sale = 1").First(&goodsDetail, param.Id)
	return goodsDetail
}

func (g *AppGoodsService) Search(param app.GoodsSearchQueryParam) []app.GoodsList {
	goodsList := make([]app.GoodsList, 0)
	Db := global.Db.Table("t_goods").Where("is_on_sale = 1")
	if param.Name != "" {
		Db = Db.Where("name like ?", "%"+param.Name+"%")
	}
	if param.Keywords != "" {
		Db = Db.Where("keywords like ?", "%"+param.Keywords+"%")
	}
	Db.Debug().Find(&goodsList)
	return goodsList
}
