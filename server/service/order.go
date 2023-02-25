package service

import (
	"encoding/json"
	"fmt"
	"nocake/global"
	"nocake/models/app"
	"strconv"
	"strings"
	"time"
)

type AppOrderService struct {
}

func (o *AppOrderService) Update(param app.OrderUpdateParam) int64 {
	fmt.Printf("param: %v\n", param)
	orderInfo := app.Order{
		Id:          param.Id,
		OpenId:      param.OpenId,
		Status:      param.Status,
		UpdatedTime: time.Now(),
	}
	return global.Db.Debug().Table("t_order").Model(&orderInfo).Updates(orderInfo).RowsAffected
}

func (o *AppOrderService) Submit(param app.OrderSubmitParam) int64 {
	cartInfo := app.CartInfo{}
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	goodsIdNumberInfo := global.Rdb.HGetAll(ctx, key).Val()
	goodsIds := make([]string, 0)
	// 遍历，获取所有的goods_id
	idsAndCounts := make(map[uint64]int, 0)
	for goodsId, cartNumber := range goodsIdNumberInfo {
		id, _ := strconv.Atoi(goodsId)
		count, _ := strconv.Atoi(cartNumber)
		goodsIds = append(goodsIds, goodsId)
		idsAndCounts[uint64(id)] = count
	}
	// 计算价格、总价、数量
	if len(goodsIdNumberInfo) > 0 {
		global.Db.Debug().Table("t_goods").Find(&cartInfo.CartItem, goodsIds)
		for i, item := range cartInfo.CartItem {
			cartInfo.CartItem[i].Carnumber = idsAndCounts[item.Id]
			cartInfo.TotalPrice = cartInfo.TotalPrice + item.Price*float64(idsAndCounts[item.Id])
			cartInfo.TotalCart = cartInfo.TotalCart + idsAndCounts[item.Id]
		}
	}
	address := app.Address{}
	orderInfo := app.Order{
		GoodsIds:    strings.Join(goodsIds, ""),
		GoodsPrice:  cartInfo.TotalPrice,
		OpenId:      param.OpenId,
		CreatedTime: time.Now(),
	}
	global.Db.Table("t_address").Where("open_id = ? and is_default = 1", param.OpenId).Find(&address)

	addressJson, _ := json.Marshal(address)
	orderInfo.Address = string(addressJson)
	rowsAffected := global.Db.Table("t_order").Create(&orderInfo).RowsAffected
	if rowsAffected > 0 {
		go global.Rdb.Del(ctx, key)
	}
	return rowsAffected
}
