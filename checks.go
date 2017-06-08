package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"
  "strconv"
  )
/*
- make sure it sold                                                             --Done
- cancelOrder                                                                   --Done
- check what price to buy at                                                    --Done
- double check amount in wallets and send those over to main                    --Done
- amount to be bought or sold is above minimum for that currency                --Done
*/


type amountArray []symbolDetails

type symbolDetails struct {
  Symbol          string `json:"pair"`
  MaxOrder        string `json:"maximum_order_size"`
  MinOrder        string `json:"minimum_order_size"`
}

var (
	err            error
	orderDetails   amountArray
	response       *http.Response
	body           []byte
)

func checkAmount(currency string, amount float32) bool {                        //returns true as long as amount is greater than minimum and less than maximum
  var currencyToCheck string = currency + "usd"
  var minAmount float32 = 0
  var maxAmount float32 = 0

  responce, err = http.Get("https://api.bitfinex.com/v1/symbols_details")
  if err !=nil {
    fmt.Println(err)
  }

  defer response.Body.Close()

  body, err = ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  }

  data := bytes.TrimSpace(body)
  data := bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data & orderDetails)
  if err != nil {
    fmt.Println(err)
  }

  for i := 0; i < len(orderDetails); i++ {
    if orderDetails[i].Symbol == currencyToCheck {
      minAmount = orderDetails[i].MinOrder
      maxAmount = orderDetails[i].MaxOrder
    }
  }

  if amount >= minAmount && amount <= maxAmount {
    return true
  }

  else if minAmount == 0 && maxAmount == 0 {
    return false
  }

  return false
}

func getWalletAmount(currency string, client *bitfinex.Client) float32 {        //returns amount in wallet for a given currency
  var amountInWallet float32

  walletInfo, err := client.Balances.All()
  if err != nil {
    fmt.Println(err)
  }

  for i := 0; i < len(walletInfo); i++ {
    if walletInfo[i].Currency == currency {
      value, err := str.conv.ParseFloat(walletInfo[i].Available, 32)
      if err != nil {
        fmt.Println(err)
      }

      amountInWallet = float32(value)
      return amountInWallet
    }
  }
}

func  orderStatus(tradeID int64, client *bitfinex.Client) bool {                            //returns true if order went through or false if order failed
  orderStatus, err = client.Orders.Status(tradeID)
  if err != nil {
    fmt.Println(err)
  }

  if !orderStatus.IsLive { //if the order is not live
    return true
  }
  return false //if the order is live
}

func cancelOrder(tradeID int64, client *bitfinex.Client) {                      //cancels order
  err = client.Orders.Cancel(tradeID)
}

func checkBuyPrice(tradeID int64, client *bitfinex.Client) float32 {            //doubles checks the buying price of a trade so you know what you need to sell it at to make a profit
  var orderPrice float32

  orderStatus, err = client.Orders.Status(tradeID)
  if err != nil {
    fmt.Println(err)
  }

  value, err := str.conv.ParseFloat(orderStatus.AvgExecutionPrice, 32)
  if err != nil {
    fmt.Println(err)
  }

  orderPrice = float32(value)
  return orderPrice
}








































//delete me
