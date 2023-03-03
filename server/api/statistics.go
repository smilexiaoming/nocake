package api

import (
	"nocake/constant"
	"nocake/response"
	"nocake/service"

	"github.com/gin-gonic/gin"
)

type WebStatistics struct {
	service.WebStatisticsService
}

func GetWebStatistics() *WebStatistics {
	return &WebStatistics{}
}

func (s *WebStatistics) GetTodayData(c *gin.Context) {
	todayData := s.TodayData()
	response.Success(constant.Selected, todayData, c)
}

func (s *WebStatistics) GetOrderData(c *gin.Context) {
	orderData := s.OrderData()
	response.Success(constant.Selected, orderData, c)
}

func (s *WebStatistics) GetShopData(c *gin.Context) {
	shopData := s.ShopData()
	response.Success(constant.Selected, shopData, c)
}
