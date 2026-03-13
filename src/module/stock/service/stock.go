package service

import "ecommerceBackend/src/module/stock/repository"

type Stock struct {
	stockRepository *repository.Stock
}

func NewStockService(repo *repository.Stock) Stock {
	return Stock{
		stockRepository: repo,
	}
}
