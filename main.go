package main

import (
	"cryptoTelegramBot/telegram"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	OpenTime           int
	CloseTime          int
}

func convertToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println(err)
	}
	return f
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
	}
	return i
}

func main() {
	telegram.SendMessage("508510759", "Hello World")
	getData("BTCUSDT", "1d")
}

func getData(symbol string, interval string) {
	response, err := http.Get("https://www.binance.com/api/v3/uiKlines?limit=1&symbol=" + symbol + "&interval=" + interval)
	if err != nil {
		log.Panicln(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	data := binanceData{}

	err = json.Unmarshal(body, &data)

	d := json.NewDecoder(strings.NewReader(string(body)))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(data)

	if err != nil {
		log.Println(err)
	}

	for _, v := range data {
		var candle candle
		candle.OpenTime = convertToInt(fmt.Sprint(v[0]))
		candle.Open = convertToFloat64(fmt.Sprint(v[1]))
		candle.High = convertToFloat64(fmt.Sprint(v[2]))
		candle.Low = convertToFloat64(fmt.Sprint(v[3]))
		candle.Close = convertToFloat64(fmt.Sprint(v[4]))
		candle.Volume = convertToFloat64(fmt.Sprint(v[5]))
		candle.CloseTime = convertToInt(fmt.Sprint(v[6]))
		fmt.Println(candle)
	}
}
