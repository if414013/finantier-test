package service

import "stock-service/model"

type StockService interface {
	GetStockInfoBySymbol(stockSymbol string) (response *model.GetStockResponse)
}
