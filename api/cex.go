package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"utiiz.dev/go/crypto/data"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*data.Rate, error) {
	currency = strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, currency))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json := string(bodyBytes)
	fmt.Println(json)

	rate := data.Rate{Currency: currency, Price: 20}
	return &rate, nil
}
