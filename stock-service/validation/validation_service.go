package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"stock-service/exception"
)

func ValidateStockSymbol(stockSymbol string) {
	err := validation.Validate(stockSymbol,
		// not empty
		validation.Required.Error("Please append Stock Symbol Name as Path Param, fex /TSLA"),
		//https://www.investopedia.com/terms/s/stocksymbol.asp
		validation.Length(2, 5).Error("Stock Symbol length must between 2 and 5 inclusively"),
		//can't contain NUMBER
		is.UpperCase.Error("Stock Symbol can only contain upper case alphabet"),
		is.UTFLetter.Error("Stock Symbol can only contain upper case alphabet"),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
