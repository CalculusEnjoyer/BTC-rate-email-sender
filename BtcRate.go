package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func getBtcUahFloat() (float64, *RequestError) {
	//https://docs.kuna.io/reference/getv3exchangeratescurrency
	//api documentation
	url := "https://api.kuna.io:443/v3/exchange-rates/btc"

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()

	if err == nil {
		body, _ := ioutil.ReadAll(response.Body)
		data := make(map[string]interface{})
		json.Unmarshal(body, &data)
		uahPrice := data["uah"].(float64)
		return uahPrice, nil
	} else {
		return 0, &RequestError{StatusCode: 404, Err: errors.New("Can not get BTC/UAH rate")}
	}
}
