package service

import (
	"encoding/json"
	"nocake/common"
	"nocake/global"
	"nocake/models/app"
	"nocake/models/web"
	"strconv"
	"strings"
	"time"
)

type AppOrderService struct {
}
type WebOrderService struct {
}

// 删除订单
func (o *WebOrderService) Delete(param web.OrderDeleteParam) int64 {
	order := app.Order{
		Deleted:     1,
		DeletedTime: time.Now(),
	}

	return global.Db.Debug().Table("t_order").Where("id = ?", param.Id).Updates(order).RowsAffected
}

// 更新订单
func (o *WebOrderService) Update(param web.OrderUpdateParam) int64 {
	order := web.Order{
		Id:          param.Id,
		Status:      param.Status,
		UpdatedTime: time.Now(),
	}
	return global.Db.Table("t_order").Model(&order).Updates(order).RowsAffected
}

// 获取订单列表
func (o *WebOrderService) GetList(param web.OrderListParam) ([]web.OrderItem, int64) {
	orders := make([]web.Order, 0)
	query := &web.Order{
		Status: param.Status,
	}
	rows := common.RestPage(param.Page, "t_order", query, &orders, &[]web.Order{})
	orderList := make([]web.OrderItem, 0)
	var user app.User
	for _, o := range orders {
		global.Db.Table("t_user").Where("open_id = ?", o.OpenId).First(&user)
		order := web.OrderItem{
			Id:            o.Id,
			Avatar:        user.Avatar,
			Nickname:      user.Nickname,
			GoodsPrice:    o.GoodsPrice,
			GoodsIdsCount: o.GoodsIdsCount,
			GoodsCount:    o.GoodsCount,
			Status:        o.Status,
			CreatedTime:   o.CreatedTime,
		}
		orderList = append(orderList, order)
	}
	return orderList, rows
}

// 获取订单详情
func (o *WebOrderService) GetDetail(param web.OrderDetailParam) (od web.OrderDetail) {
	var order web.Order
	// 查询订单信息
	global.Db.Table("t_order").First(&order, param.Id)

	goodIds := make([]string, 0)
	goodsIdCount := make(map[string]string, 0)
	err := json.Unmarshal([]byte(order.GoodsIdsCount), &goodsIdCount)
	if err != nil {
		return web.OrderDetail{}
	}
	goods := make([]app.Goods, 0)
	goodsCount := 0
	for k, v := range goodsIdCount {
		goodIds = append(goodIds, k)
		count, _ := strconv.Atoi(v)
		goodsCount += count
	}
	orderDetail := web.OrderDetail{
		Order: order,
	}
	global.Db.Debug().Table("t_goods").Find(&goods, goodIds)
	for _, v := range goods {
		idstr := strconv.Itoa(v.Id)
		coutInt, _ := strconv.Atoi(goodsIdCount[idstr])
		goodItem := web.GoodsItem{
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
	rows_affect := global.Db.Table("t_address").Where("open_id = ? and id = ?", param.OpenId, param.AddressId).Find(&address).RowsAffected
	adressStr := ""
	addressInfo := app.AddressAddParam{
		Name:     address.Name,
		OpenId:   address.OpenId,
		Province: address.Province,
		City:     address.City,
		County:   address.County,
		Detail:   address.Detail,
		Tel:      address.Tel,
	}
	if rows_affect > 0 {
		addressByte, _ := json.Marshal(addressInfo)
		adressStr = string(addressByte)
	}
	orderInfo := app.Order{
		Address:       adressStr,
		GoodsIdsCount: string(GoodsIdsCountInfo),
		GoodsPrice:    cartInfo.TotalPrice,
		GoodsCount:    cartInfo.TotalCart,
		OpenId:        param.OpenId,
		CreatedTime:   time.Now(),
		Message:       param.Message,
	}

	rowsAffected = global.Db.Table("t_order").Create(&orderInfo).RowsAffected
	if rowsAffected > 0 {
		go global.Rdb.Del(ctx, key)
	}
	return rowsAffected
}

func (o *AppOrderService) GetList(param app.OrderQueryListParam) []app.OrderInfo {
	orders := make([]app.Order, 0)
	global.Db.Debug().Table("t_order").Where("open_id = ? and status = ?", param.OpenId, param.Status).Limit(param.PageSize).Offset((param.PageNum - 1) * param.PageSize).Find(&orders)
	orderList := make([]app.OrderInfo, 0)
	for _, order := range orders {
		orderList = append(orderList, GetDetail(order))
	}
	return orderList
}

func GetDetail(param app.Order) app.OrderInfo {
	goodIds := make([]string, 0)
	goodsIdCount := make(map[string]string, 0)
	err := json.Unmarshal([]byte(param.GoodsIdsCount), &goodsIdCount)
	if err != nil {
		return app.OrderInfo{}
	}
	goods := make([]app.Goods, 0)
	goodsCount := 0
	for k, v := range goodsIdCount {
		goodIds = append(goodIds, k)
		count, _ := strconv.Atoi(v)
		goodsCount += count
	}
	orderDetail := app.OrderInfo{
		OrderId:    param.Id,
		Status:     param.Status,
		TotalPrice: param.GoodsPrice,
		GoodsCount: goodsCount,
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
