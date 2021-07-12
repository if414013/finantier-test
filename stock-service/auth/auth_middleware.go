package auth

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"stock-service/exception"
)

func AuthenticateRequest(token string) (isAuthenticated bool) {
	resp, err := http.Get(os.Getenv("AUTH_SERVICE_ENDPOINT") + "?token=" + token)
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		exception.PanicIfNeeded(err)
	}
	stringBody := string(body)
	return gjson.Get(stringBody, "data.isAuthorized").Bool()
}
