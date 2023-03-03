package service

import (
	"nocake/common"
	"nocake/global"
	"nocake/models/web"
	"strconv"
	"time"
)

type WebStatisticsService struct {
}

// 统计当日数据
func (s *WebStatisticsService) TodayData() web.TodayDate {
	today := common.NowTime()
	day := today[:10] + "%"
	pps := "status = 0 and created_time like ?"
	pds := "status = 1 and created_time like ?"
	prs := "status = 2 and created_time like ?"
	ras := "status = 3 and created_time like ?"
	fds := "status = 4 and created_time like ?"
	pas := "created_time like ?"
	orderDate := web.TodayDate{}
	global.Db.Table("t_order").Where(pps, day).Find(&web.Order{}).Count(&orderDate.PendPay)
	global.Db.Table("t_order").Where(prs, day).Find(&web.Order{}).Count(&orderDate.Payed)
	global.Db.Table("t_order").Where(ras, day).Find(&web.Order{}).Count(&orderDate.InDelivery)
	global.Db.Table("t_order").Where(pds, day).Find(&web.Order{}).Count(&orderDate.Canceled)
	global.Db.Table("t_order").Where(fds, day).Find(&web.Order{}).Count(&orderDate.Finished)
	global.Db.Table("t_order").Where(pas, day).Pluck("IFNULL(SUM(goods_price), 0) as pay_amount", &orderDate.PayAmount)
	return orderDate
}

// 统计订单数据
func (s *WebStatisticsService) OrderData() web.OrderData {
	var orders web.OrderData
	today := common.NowTime()
	for i, hour := 0, 6; i < 17; i++ {
		var hs string
		if hour < 10 {
			hs = "0" + strconv.Itoa(hour)
		} else {
			hs = strconv.Itoa(hour)
		}
		cond := today[:10] + " " + hs + "%"
		global.Db.Table("t_order").Where("created_time like ?", cond).Count(&orders.Orders[i])
		hour = hour + 1
	}
	return orders
}

// 统计店铺数据
func (s *WebStatisticsService) ShopData() web.ShopData {
	var weekInfo web.ShopData
	switch time.Now().Weekday() {
	case time.Monday:
		weekInfo = getWeekData(1)
	case time.Tuesday:
		weekInfo = getWeekData(2)
	case time.Wednesday:
		weekInfo = getWeekData(3)
	case time.Thursday:
		weekInfo = getWeekData(4)
	case time.Friday:
		weekInfo = getWeekData(5)
	case time.Saturday:
		weekInfo = getWeekData(6)
	case time.Sunday:
		weekInfo = getWeekData(7)
	default:
	}
	return weekInfo
}

func getWeekData(days int) web.ShopData {
	var wd web.ShopData
	for i, index := days-1, 0; i >= 0; i-- {
		var result []float64
		var amountSum float64
		nowTime := common.WeekTime(i) + "%"
		cLike := "created_time like ?"
		global.Db.Table("t_order").Where(cLike, nowTime).Pluck("goods_price", &result)
		for _, v := range result {
			amountSum += v
		}
		wd.Amount[index] = amountSum
		index++
	}
	return wd
}
