package repository

import (
	"log"
	"github.com/SimpleApplicationsOrg/stock/alphavantage"
	"stock-service/entity"
)

func NewStockRepository() StockRepository {
	return &stockRepositoryImpl{}
}

type stockRepositoryImpl struct {
}

func (repository *stockRepositoryImpl) GetStockInfoBySymbol(stockSymbol string) (response *entity.Stock) {
	// instantiate client
	avClient, err := alphavantage.NewAVClient()

	if err != nil {
		log.Printf("error whwn getting alphavantege client: %+v", err.Error())
		return nil
	}

	avResponse, err := avClient.TimeSeriesIntraday(stockSymbol, "1min")
	if err != nil {
		log.Printf("error calling alphavantege api: %+v", err.Error())
		return nil
	}

	timeSeries := *avResponse.TimeSeries
	for _, timeStamp := range timeSeries.TimeStamps() {
		value := (timeSeries)[timeStamp]
		response = &entity.Stock{
			Symbol:    stockSymbol,
			Timestamp: timeStamp,
			Open:      value.Open(),
			High:      value.High(),
			Low:       value.Low(),
			Close:     value.Close(),
			Volume:    value.Close(),
		}
		// just return 1 of the record (the latest info)
		return response
	}

	return nil
}
