package service

import (
	"fmt"
	"nocake/common"
	"nocake/global"
	"nocake/models/app"
	"nocake/models/web"
	"time"
)

type WebGoodsService struct {
}
type AppGoodsService struct {
}

// 创建商品
func (g *WebGoodsService) Create(param web.GoodsCreateParam) int64 {
	goods := web.Goods{
		Name:        param.Name,
		Brief:       param.Brief,
		CategoryId:  param.CategoryId,
		Keywords:    param.Keywords,
		Status:      param.Status,
		Weight:      param.Weight,
		PicUrl:      param.PicUrl,
		PicUrls:     param.PicUrls,
		Unit:        param.Unit,
		Price:       param.Price,
		CreatedTime: time.Now(),
	}
	return global.Db.Table("t_goods").Create(&goods).RowsAffected
}

// 删除商品
func (g *WebGoodsService) Delete(param web.GoodsDeleteParam) int64 {
	goods := app.Goods{
		Deleted:     1,
		DeletedTime: time.Now(),
	}
	return global.Db.Debug().Table("t_goods").Where("id = ?", param.Id).Updates(goods).RowsAffected
}

// 更新商品
func (g *WebGoodsService) Update(param web.GoodsUpdateParam) int64 {
	goods := web.Goods{
		Id:          param.Id,
		Name:        param.Name,
		Brief:       param.Brief,
		CategoryId:  param.CategoryId,
		Keywords:    param.Keywords,
		Status:      param.Status,
		Weight:      param.Weight,
		PicUrl:      param.PicUrl,
		PicUrls:     param.PicUrls,
		Unit:        param.Unit,
		Quantity:    param.Quantity,
		Price:       param.Price,
		UpdatedTime: time.Now(),
	}
	fmt.Printf("goods: %#v\n", goods)
	return global.Db.Debug().Table("t_goods").Model(&goods).Updates(goods).RowsAffected
}

// 更新商品状态
func (g *WebGoodsService) UpdateStatus(param web.GoodsStatusUpdateParam) int64 {
	goods := web.Goods{
		Id:     param.Id,
		Status: param.Status,
	}
	return global.Db.Debug().Table("t_goods").Model(&goods).Update("status", goods.Status).RowsAffected
}

// 获取商品列表
func (g *WebGoodsService) GetList(param web.GoodsListParam) ([]web.GoodsList, int64) {
	query := &web.Goods{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Name:       param.Name,
		Status:     param.Status,
		Deleted:    "0",
	}
	goodsList := make([]web.GoodsList, 0)
	rows := common.RestPage(param.Page, "t_goods", query, &goodsList, &[]web.Goods{})
	return goodsList, rows
}

func (g *AppGoodsService) GetList(param app.GoodsListQueryParam) []app.GoodsList {
	goodsList := make([]app.GoodsList, 0)
	global.Db.Debug().Table("t_goods").Where("category_id = ?", param.CategoryId).Limit(param.PageSize).Offset((param.PageNum - 1) * param.PageSize).Find(&goodsList)
	return goodsList
}

func (g *AppGoodsService) GetDetail(param app.GoodsDetailQueryParam) app.GoodsDetail {
	goodsDetail := app.GoodsDetail{}
	global.Db.Debug().Table("t_goods").Where("status = 0").First(&goodsDetail, param.GoodsId)
	return goodsDetail
}

func (g *AppGoodsService) Search(param app.GoodsSearchQueryParam) []app.GoodsList {
	goodsList := make([]app.GoodsList, 0)
	Db := global.Db.Table("t_goods").Where("status = 0")
	if param.Name != "" {
		Db = Db.Where("name like ?", "%"+param.Name+"%")
	}
	if param.Keywords != "" {
		Db = Db.Where("keywords like ?", "%"+param.Keywords+"%")
	}
	Db.Debug().Limit(param.PageSize).Offset((param.PageNum - 1) * param.PageSize).Find(&goodsList)
	return goodsList
}
