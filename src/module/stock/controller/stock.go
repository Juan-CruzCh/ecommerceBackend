package controller

import (
	"ecommerceBackend/src/module/stock/service"
)

type Stock struct {
	stockService *service.Stock
}

func NewStockController(service *service.Stock) Stock {
	return Stock{
		stockService: service,
	}
}
