package main

import (
  "fmt"
  "strings"
  "github.com/bitfinexcom/bitfinex-api-go/v1"
)

func buy(currency string, amount float32, price float32, client *bitfinex.Client) int64 {                      //buy function
  var typeOfCoinToSell string = currency + "usd"

  coinToSell = ToUpper(bitfinex.currency + "USD")
  data, err := client.Orders.Create(coinToSell, amount, price, bitfinex.OrderTypeExchangeLimit)

  if err != nil {
    fmt.Println("Error:", err)
  }

  return data.ID


}

func sell(currency string, amount float32, price float32, client *bitfinex.Client) int64 {                     //sell function
  var typeOfCoinToSell string = currency + "usd"
  var negAmount float32 = amount * -1

  coinToSell = ToUpper(bitfinex.currency + "USD")
  negAmount = -1*amount
  data, err := client.Orders.Create(coinToSell, negAmount, price, bitfinex.OrderTypeExchangeLimit)

  if err != nil {
    fmt.Println("Error:", err)
  }

  return data.ID

}

func sellMarketOrder(currency string, amount float32, client *bitfinex.Client) {
  var typeOfCoinToSell string = currency + "usd"
  var negAmount float32 = amount * -1

  coinToSell = ToUpper(bitfinex.currency + "USD")
  negAmount = -1*amount
  data, err := client.Orders.Create(coinToSell, negAmount, price, bitfinex.OrderTypeExchangeMarket)

  if err != nil {
    fmt.Println("Error:", err)
  }
}
