package validation

import (
	"encryption-service/exception"
	"encryption-service/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateStockSymbol(request model.StockInfo) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Symbol, validation.Required),
		validation.Field(&request.Timestamp, validation.Required),
		validation.Field(&request.Open, validation.Required),
		validation.Field(&request.High, validation.Required),
		validation.Field(&request.Low, validation.Required),
		validation.Field(&request.Close, validation.Required),
		validation.Field(&request.Volume, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
