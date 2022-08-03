package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type binanceData [][]interface{}

type candle struct {
	Symbol             string
	Open               float64
	Close              float64
	High               float64
	Low                float64
	Interval           string
	Volume             float64
	UpperShadowPercent float64
	DownShadowPercent  float64
	OpenTime           int32
	CloseTime          int32
}

func main() {

	response, err := http.Get("https://www.binance.com/api/v3/uiKlines?limit=1&symbol=BTCUSDT&interval=4h")
	if err != nil {
		log.Panicln(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	data := binanceData{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Println(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}

}
