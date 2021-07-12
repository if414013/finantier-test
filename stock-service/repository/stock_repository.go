package repository

import "stock-service/entity"

type StockRepository interface {
	GetStockInfoBySymbol(stockSymbol string) (response *entity.Stock)
}
