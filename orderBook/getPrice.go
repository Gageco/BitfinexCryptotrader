package main

import (
  //"bytes"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	"time"
  "strconv"
)

type Ticker struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

func updateTicker(currency string, ticker *Ticker) {
  var urlCall string = "https://api.bitfinex.com/v1/pubticker/" + currency + "usd"

  client := &http.Client{Timeout: time.Second * 5}
  resp, _ := client.Get(urlCall)
  json.NewDecoder(resp.Body).Decode(&ticker)
}

func getMid(currency string) float32 {
  var ticker Ticker
  var float float32

  updateTicker(currency, &ticker)

  value, err := strconv.ParseFloat(ticker.Mid, 32)
  if err != nil {
    fmt.Println("error line 40")
  }
  float = float32(value)

  return float                                         //get mid price
}

func getHigh(currency string) float32 {                                         //get high price
  var ticker Ticker
  var float float32

  updateTicker(currency, &ticker)

  value, err := strconv.ParseFloat(ticker.High, 32)
  if err != nil {
    fmt.Println("error line 40")
  }
  float = float32(value)

  return float
}

func getBid(currency string) float32 {                                          //get bid price
  var ticker Ticker
  var float float32

  updateTicker(currency, &ticker)

  value, err := strconv.ParseFloat(ticker.Bid, 32)
  if err != nil {
    fmt.Println("error line 40")
  }
  float = float32(value)

  return float
}

func getAsk(currency string) float32 {                                          //get asking price
  var ticker Ticker
  var float float32

  updateTicker(currency, &ticker)

  value, err := strconv.ParseFloat(ticker.Ask, 32)
  if err != nil {
    fmt.Println("error line 40")
  }
  float = float32(value)

  return float
}

func getLastPrice(currency string) float32 {                                    //get LastPrice
  var ticker Ticker
  var float float32

  updateTicker(currency, &ticker)

  value, err := strconv.ParseFloat(ticker.LastPrice, 32)
  if err != nil {
    fmt.Println("error line 40")
  }
  float = float32(value)

  return float
}

func getLow(currency string) float32 {                                          //get low price
  var ticker Ticker
  var float float32

  updateTicker(currency, &ticker)

  value, err := strconv.ParseFloat(ticker.Low, 32)
  if err != nil {
    fmt.Println("error line 40")
  }
  float = float32(value)

  return float
}
