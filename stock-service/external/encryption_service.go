package external

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"stock-service/exception"
	"stock-service/model"
)

func EncryptStockInfo(stockInfo model.GetStockResponse) (encryptedStockInfo *model.GetStockResponse) {
	postBody, _ := json.Marshal(stockInfo)
	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(os.Getenv("ENCRYPTION_SERVICE_ENDPOINT"), "application/json", requestBody)

	if err != nil {
		exception.PanicIfNeeded(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	stringBody := string(body)
	//stringData := gjson.Get(stringBody, "data")
	return &model.GetStockResponse{
		Symbol:    gjson.Get(stringBody, "data.symbol").String(),
		Timestamp: gjson.Get(stringBody, "data.timestamp").String(),
		Open:      gjson.Get(stringBody, "data.open").String(),
		High:      gjson.Get(stringBody, "data.high").String(),
		Low:       gjson.Get(stringBody, "data.low").String(),
		Close:     gjson.Get(stringBody, "data.close").String(),
		Volume:    gjson.Get(stringBody, "data.volume").String(),
	}
}
