package service

import (
	"encoding/json"
	"nocake/global"
	"nocake/models/app"
	"strconv"
	"strings"
	"time"
)

type AppOrderService struct {
}

func (o *AppOrderService) Update(param app.OrderUpdateParam) int64 {
	orderInfo := app.Order{
		Id:          param.OrderId,
		OpenId:      param.OpenId,
		Status:      param.Status,
		UpdatedTime: time.Now(),
	}
	return global.Db.Debug().Table("t_order").Model(&orderInfo).Updates(orderInfo).RowsAffected
}

func (o *AppOrderService) Submit(param app.OrderSubmitParam) int64 {
	var rowsAffected int64
	cartInfo := app.CartInfo{}
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	goodsIdNumberInfo := global.Rdb.HGetAll(ctx, key).Val()
	GoodsIdsCountInfo, _ := json.Marshal(goodsIdNumberInfo)
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
	if len(goodsIdNumberInfo) <= 0 {
		return rowsAffected
	}

	global.Db.Debug().Table("t_goods").Find(&cartInfo.CartItem, goodsIds)
	for i, item := range cartInfo.CartItem {
		cartInfo.CartItem[i].Carnumber = idsAndCounts[item.Id]
		cartInfo.TotalPrice = cartInfo.TotalPrice + item.Price*float64(idsAndCounts[item.Id])
		cartInfo.TotalCart = cartInfo.TotalCart + idsAndCounts[item.Id]
	}

	address := app.AddressAddParam{}
	rows_affect := global.Db.Table("t_address").Where("open_id = ? and is_default = 1", param.OpenId).Find(&address).RowsAffected
	adressStr := ""
	addressInfo := app.AddressAddParam{
		Name:       address.Name,
		OpenId:     address.OpenId,
		Province:   address.Province,
		City:       address.City,
		County:     address.County,
		Detail:     address.Detail,
		AreaCode:   address.AreaCode,
		PostalCode: address.PostalCode,
		Tel:        address.Tel,
	}
	if rows_affect > 0 {
		addressByte, _ := json.Marshal(addressInfo)
		adressStr = string(addressByte)
	}
	orderInfo := app.Order{
		Address:       adressStr,
		GoodsIdsCount: string(GoodsIdsCountInfo),
		GoodsPrice:    cartInfo.TotalPrice,
		OpenId:        param.OpenId,
		CreatedTime:   time.Now(),
	}

	rowsAffected = global.Db.Table("t_order").Create(&orderInfo).RowsAffected
	if rowsAffected > 0 {
		go global.Rdb.Del(ctx, key)
	}
	return rowsAffected
}

func (o *AppOrderService) GetList(param app.OrderQueryListParam) []app.Order {
	orderList := make([]app.Order, 0)
	global.Db.Debug().Table("t_order").Where("open_id = ? and status = ?", param.OpenId, param.Status).Limit(param.PageSize).Offset((param.PageNum - 1) * param.PageSize).Find(&orderList)
	return orderList
}

func (o *AppOrderService) GetDetail(param app.OrderQueryDetailParam) app.OrderDetail {
	orderDetail := app.OrderDetail{}
	order := app.Order{}
	global.Db.Debug().Table("t_order").Where("id = ? and open_id = ?", param.OrderId, param.OrderId).Find(&order)
	goodIds := make([]string, 0)
	goodsIdCount := make(map[string]string, 0)
	err := json.Unmarshal([]byte(order.GoodsIdsCount), &goodsIdCount)
	if err != nil {
		return orderDetail
	}
	goods := make([]app.Goods, 0)
	for k, _ := range goodsIdCount {
		goodIds = append(goodIds, k)
	}
	global.Db.Debug().Table("t_goods").Find(&goods, goodIds)
	for _, v := range goods {
		idstr := strconv.Itoa(v.Id)
		coutInt, _ := strconv.Atoi(goodsIdCount[idstr])
		goodItem := app.GoodsItem{
			Id:     v.Id,
			Name:   v.Name,
			Price:  v.Price,
			PicUrl: v.PicUrl,
			Count:  coutInt,
		}
		orderDetail.GoodsItem = append(orderDetail.GoodsItem, goodItem)
	}
	return orderDetail
}
