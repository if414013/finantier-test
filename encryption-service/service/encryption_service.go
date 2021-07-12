package service

import "encryption-service/model"

type EncryprionService interface {
	EncryptStockInfo(stockInfo model.StockInfo) (response *model.StockInfo)
}
