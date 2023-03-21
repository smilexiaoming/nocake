package service

import (
	"encoding/json"
	"nocake/global"
	"nocake/models/app"
	"strconv"
	"strings"
)

type AppCartService struct {
}

func (c *AppCartService) Set(param app.CartSetParam) bool {
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	goodsId := strconv.Itoa(int(param.GoodsId))
	global.Rdb.HSet(ctx, key, goodsId, param.Options).Val()
	return true
}

func (c *AppCartService) Delete(param app.CartDeleteParam) int64 {
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	return global.Rdb.HDel(ctx, key, param.GoodsId).Val()
}

func (c *AppCartService) Clear(param app.CartClearParam) int64 {
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	return global.Rdb.Del(ctx, key).Val()
}

func (c *AppCartService) GetInfo(param app.CartQueryParam) app.CartInfo {
	cartInfo := app.CartInfo{}
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	goodsInfo := global.Rdb.HGetAll(ctx, key).Val()
	goodsIds := make([]string, 0)
	idsAndCounts := make(map[uint64]int, 0)
	optionPriceMap := make(map[int]float64)
	for goodsId, infos := range goodsInfo {
		Options := app.Options{}
		json.Unmarshal([]byte(infos), &Options)
		id, _ := strconv.Atoi(goodsId)

		// 存id列表
		goodsIds = append(goodsIds, goodsId)

		// 存数量
		count := Options.Count
		idsAndCounts[uint64(id)] = count

		// 存配件的价格
		var tempOptionPirce float64
		tempOptionPirce = 0
		for _, item := range Options.Option {
			for _, v := range item.Item {
				if v.Active {
					tempOptionPirce += float64(v.Price)
				}
			}
		}
		optionPriceMap[id] = tempOptionPirce
	}
	// 计算价格、总价、数量
	if len(goodsInfo) > 0 {
		global.Db.Debug().Table("t_goods").Find(&cartInfo.CartItem, goodsIds)
		for i, item := range cartInfo.CartItem {
			cartInfo.CartItem[i].Options = goodsInfo[strconv.Itoa(int(item.Id))]
			cartInfo.CartItem[i].Price = optionPriceMap[int(item.Id)] + item.Price*float64(idsAndCounts[item.Id])
			cartInfo.TotalPrice += cartInfo.CartItem[i].Price
			cartInfo.TotalCart += idsAndCounts[item.Id]
		}
	}
	return cartInfo
}
