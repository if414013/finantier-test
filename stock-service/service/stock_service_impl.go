package service

import (
	"log"
	"stock-service/external"
	"stock-service/model"
	"stock-service/repository"
)

func NewStockService(stockRepository *repository.StockRepository) StockService {
	return &stockServiceImpl{
		StockRepository: *stockRepository,
	}
}

type stockServiceImpl struct {
	StockRepository repository.StockRepository
}

func (service *stockServiceImpl) GetStockInfoBySymbol(stockSymbol string) (response *model.GetStockResponse) {
	stock := service.StockRepository.GetStockInfoBySymbol(stockSymbol)
	if stock == nil {
		return nil
	}
	response = &model.GetStockResponse{
		Symbol:    stock.Symbol,
		Timestamp: stock.Timestamp,
		Open:      stock.Open,
		High:      stock.High,
		Low:       stock.Low,
		Close:     stock.Close,
		Volume:    stock.Volume,
	}
	log.Printf("Stock Info fetched successfully: %+v", response)
	return external.EncryptStockInfo(*response)
}
