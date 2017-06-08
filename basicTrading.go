package main
// run this program with go run *.go
import (
  "fmt"
  "time"
)

func basicTrade(currency string, array [100]float32, client *bitfinex.Client, buyingPrices *[100]float32) {
  var orignialPercent float32 = .05                                               //5%
  var percentageChange float32 = .005                                           //.5%
  var numberOfPointsToLookAt int = 36                                           //36 because at 20min intervals, this is 12hrs
  var nowPoint float32 = array[0]                                               //most recent price in array

  var lastPoint float32
  var actualPercentDifference float32
  var neededPercent float32
  var tradeID int64
  var actualBuyPrice float32

  for i := 0; i < numberOfPointsToLookAt; i++ {
    lastPoint = array[numberOfPointsToLookAt-i]                                 //this should go from 36 -> 0
    actualPercentDifference = (lastPoint - nowPoint)/nowPoint                         //calculates percent difference between the points
    neededPercent = orignialPercent + (percentageChange * i)                    //the needed pecent change to buy or sell, increases with i

    /*
    real quick this calculates the percent difference and as it gets closer to the current time
    it increases the needed percent to acutally buy and sell, this is to prevent volitile times from triggering buys and sells
    */

    if actualPercentDifference >= neededPercent {                               //buy if positive percent meets requirements
      var priceToBuy float32 = priceToBuy(currency)
      var amountToBuy float32 = amountToBuy()
      var buyOrderSuccessful bool
      if shouldBuy(currency) { // write this <---
        if checkAmount(currency, amountToBuy) {                                   //if within limits of min and max amount that can be bought and should
          if getWalletAmount("usd", client) > amountToBuy {                       //check theres enough in wallet
            tradeID = buy(currency, amountToBuy, priceToBuy, client)                      //buy
            fmt.Println("Attempting to buy")
            for j := 0; j < 6; j++ {                                              //try 6 times at 10 second intervals to buy
              time.Sleep(10 * time.Second)
              if orderStatus(tradeID, client) {
                fmt.Println("Buy order went through successfully")
                actualBuyPrice = checkBuyPrice(tradeID, client)
                for k := 0; k < len(buyingPrices); k++ {
                  if buyingPrices[k] == 0 { //fill the next avaliable spot thats not already filled with the newwest buying price
                    buyingPrices[k] = actualBuyPrice
                    break
                  }
                }
                buyOrderSuccessful = true
                j = 10
                fmt.Println("You have bought",amountToBuy,currency,"at",priceToBuy,"for",amountToBuy*priceToBuy)
                break
              }
              else {
                fmt.Println("Buy order has not completed, retrying", 6-i,"more times")
                buyOrderSuccessful = false
              }
            }
            if !buyOrderSuccessful {
              fmt.Println("Order has failed to go through completely, sorry")
              cancelOrder(orderID, client)
            }
          }
          else {
            fmt.Println("Not enough in wallet to buy with")
          }
        }
        else {
          fmt.Println("Ordersize to big or to small to buy with")
        }
      }
    }

    else if actualPercentDifference <= (neededPercent*-1) {                     //sell if negative percent meets requirements
      var priceToSell float32 = priceToSell(buyingPrices)
      var amountToSell float32 = amountToSell(currency, client)
      var sellOrderSuccessful bool
      if shouldSell(priceToSell) {
        if checkAmount(currency, amountToSell) {                                  //check if its within the max and min amounts you can sell
          if getWalletAmount(currency, client) > amountToSell {                   //check you have enough that you can sell as much as you want
            tradeID = sell(currency, amountToSell, priceToSell, client)                   //sell
            fmt.Println("Attempting to sell")
            for j :=0; j < 6; j++ {                                               //try 6 times at 10 second intervals to sell
              time.Sleep(10 * time.Second)
              if orderStatus(tradeID, client) {
                fmt.Println("Sell order went through successfully")
                for k := 0; k < len(buyingPrices); j++ {                        //because it all sold clear out the array
                  buyingPrices[k] = 0
                }
                sellOrderSuccessful = true
                j = 10
                fmt.Println("You have sold", amountToSell,currency,"at",priceToSell,"for",amountToSell*priceToSell)
                break
              }
              else {
                fmt.Println("Sell order has not been completed, retrying", 6-i, "more times")
                sellOrderSuccessful = false
              }
            }
            if !sellOrderSuccessful {
              fmt.Println("Sell order has failed completely, sorry,doing a market order for the rest")
              cancelOrder(orderID, client)
              sellMarketOrder(currency, amountToSell(currency, client), priceToSell(buyingPrices), client) //this should always go through, just in case the order is unsuccessful or only partially successful.
              //Maybe turn this into a thing that sells at a lower price and tries again
            }

          }
          else {
            fmt.Println("Not enough in wallet to buy with")
          }
        }
        else {
          fmt.Println("Ordersize to big or to small to sell with")
        }
      }
    }
  }
}

func amountToBuy() float32 {
  var amountInWallet float32 = getWalletAmount("usd", client)
  var amountToBuyWith float32

  amountToBuyWith = amountInWallet - amountInWallet *.1 //buy with only 90% of money in wallet
  return amountToBuyWith
}

func priceToBuy(currency) float32 {
  var bid float32 = getBid(currency)
  var buyPrice float32

  buyPrice = bid - bid*.01 //buying price is 1% lower than the current bid, to make it more likely to go through
  return buyPrice
}

func shouldBuy() bool {
  return true //you should always buy, maybe write something in here later
}

func amountToSell(currency, client) float32 {
  return getWalletAmount(currency, client) //this is redundant, but whoooo thhheeee heecccckkk carrreeessss!!!!
}

func priceToSell(buyingPrices [100]float32) float32 {
  var sum float32 = 0
  var average float32
  var lengthOfArrayToEnd uint8

  for i := 0; i < len(buyingPrices); i++ {  //get average buying price
    sum += buyingPrices[i]
    if buyingPrices[i] == 0 {
      lengthOfArrayToEnd = i
      break;
    }
  }

  average = sum/lengthOfArrayToEnd
  return average * 1.05 //five percent abouve the average buying price

}

func shouldSell(price float32) bool {
  if getLastPrice(currency) > price && getAsk(currency) > price {
    return true
  }
  else {
    return false
  }
}
