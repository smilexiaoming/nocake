package service

import (
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
	global.Rdb.HSet(ctx, key, goodsId, int64(param.Carnumber)).Val()
	return true
}

func (c *AppCartService) Update(param app.CartUpdateParam) int64 {
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	goodsId := strconv.Itoa(int(param.GoodsId))
	if param.Carnumber != 0 {
		value := global.Rdb.HIncrBy(ctx, key, goodsId, int64(param.Carnumber)).Val()
		if value < 0 {
			return global.Rdb.HDel(ctx, key, goodsId).Val()
		}
		return value
	}
	return global.Rdb.HIncrBy(ctx, key, goodsId, int64(param.Carnumber)).Val()
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
	return cartInfo
}
