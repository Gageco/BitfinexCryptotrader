package main

import (
  "fmt"
  "strings"
  "github.com/bitfinexcom/bitfinex-api-go/v1"
)

func buy(currency string, amount float32, price float32, client *bitfinex.Client) int64 {                      //buy function
  var typeOfCoinToSell string = currency + "usd"

  coinToSell := strings.ToUpper(typeOfCoinToSell)
  data, err := client.Orders.Create(coinToSell, float64(amount), float64(price), bitfinex.OrderTypeExchangeLimit)

  if err != nil {
    fmt.Println("Error:", err)
  }

  return data.ID


}

func sell(currency string, amount float32, price float32, client *bitfinex.Client) int64 {                     //sell function
  var typeOfCoinToSell string = currency + "usd"
  var negAmount float32 = amount * -1

  coinToSell := strings.ToUpper(typeOfCoinToSell)
  negAmount = -1*amount
  data, err := client.Orders.Create(coinToSell, float64(negAmount), float64(price), bitfinex.OrderTypeExchangeLimit)

  if err != nil {
    fmt.Println("Error:", err)
  }

  return data.ID

}

func sellMarketOrder(currency string, amount float32, price float32, client *bitfinex.Client) int64 {
  var typeOfCoinToSell string = currency + "usd"
  var negAmount float32 = amount * -1

  coinToSell := strings.ToUpper(typeOfCoinToSell)
  negAmount = -1*amount
  data, err := client.Orders.Create(coinToSell, float64(negAmount), float64(price), bitfinex.OrderTypeExchangeMarket)

  if err != nil {
    fmt.Println("Error:", err)
  }
  return data.ID
}

func funcAmountToBuy(client *bitfinex.Client) float32 {
  var amountInWallet float32 = getWalletAmount("usd", client)
  var amountToBuyWith float32

  amountToBuyWith = amountInWallet - amountInWallet *.1 //buy with only 90% of money in wallet
  return amountToBuyWith
}

func funcPriceToBuy(currency string) float32 {
  var bid float32 = getBid(currency)
  var buyPrice float32

  buyPrice = bid - bid*.01 //buying price is 1% lower than the current bid, to make it more likely to go through
  return buyPrice
}

func shouldBuy() bool {
  return true //you should always buy, maybe write something in here later
}

func funcAmountToSell(currency string, client *bitfinex.Client) float32 {
  return getWalletAmount(currency, client) //this is redundant, but whoooo thhheeee heecccckkk carrreeessss!!!!
}

func funcPriceToSell(buyingPrices [100]float32) float32 {
  var sum float32 = 0
  var average float32
  var lengthOfArrayToEnd int

  for i := 0; i < len(buyingPrices); i++ {  //get average buying price
    sum += buyingPrices[i]
    if buyingPrices[i] == 0 {
      lengthOfArrayToEnd = i
      break;
    }
  }

  average = float32(sum)/float32(lengthOfArrayToEnd)
  return average * 1.05 //five percent abouve the average buying price

}

func shouldSell(price float32, currency string) bool {
  if getLastPrice(currency) > price && getAsk(currency) > price {
    return true
  } else {
    return false
  }
}
