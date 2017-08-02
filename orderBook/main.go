package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"
  // "github.com/wcharczuk/go-chart"
  "strconv"
)

const lengthOfOrderBooks = 60
const lengthOfBidsAsks = 10.0

func main() {
  var orders orderBooks
  var upTrendCounter int
  var downTrendCounter int
  var bidAmounts [lengthOfBidsAsks]float64
  var askAmounts [lengthOfBidsAsks]float64
  var err error
  var downTrend bool
  var upTrend bool
  var averageBidAmount float64
  var averageAskAmount float64
  var totalBidAmount float64
  var totalAskAmount float64
  correct := 0
  incorrect := 0

  // var order orderBook
  for {
    lastPrice := getLastPrice("btc")
    fmt.Println("Getting Orderbook Info")
  for i := 0; i < lengthOfOrderBooks; i++ {                                                   //get orderbook info
    getMostRecentOrders(&orders[i])
    time.Sleep(250 * time.Millisecond)
    // fmt.Println("Seconds Left:", 30-i)
  }
    /*
    This little section looks at the 30 orderbook flashes retrieved previously.
    It looks at the first four bids and asks of the flashes and then takes the average.
    If the averageBid > averageAsk upTrendCounter increments 1
    if averageAsk > averageBid downTrendCounter increments 1

    It then compares upTrendCounter and downTrendCounter to determine if an uptrend or downtrend is more likely.
    This is done based on if one is significantly higher than the other
    */
    downTrendCounter = 0
    upTrendCounter = 0
    for i := 0; i < lengthOfOrderBooks; i++ {
      for j := 0; j < lengthOfBidsAsks-1; j++ {                                                  //look at first four orders in order books for each second
        bidAmounts[j], err = strconv.ParseFloat(orders[i].Bids[j].Amount, 32)
        totalBidAmount += bidAmounts[j]
        checkErr(err)

        askAmounts[j], err = strconv.ParseFloat(orders[i].Asks[j].Amount, 32)
        totalAskAmount += askAmounts[j]
        checkErr(err)
      }

      //calculate average bid and ask price
      averageBidAmount = float64(totalBidAmount)/float64(lengthOfBidsAsks)
      averageAskAmount = float64(totalAskAmount)/float64(lengthOfBidsAsks)

      if averageBidAmount > averageAskAmount {         //more bids than asks, more buys than sells
        upTrendCounter++
      } else if averageAskAmount > averageBidAmount {  //more asks than bids, more sells than buys
        downTrendCounter++
      }
    }
    fmt.Println("Uptrend: ",upTrendCounter)
    fmt.Println("Downtrend: ",downTrendCounter)
    fmt.Println("r: ", (totalBidAmount - totalAskAmount)/(totalBidAmount + totalAskAmount))

    /*
    if uptrend or downtrend is 15% than the other it will predict either up or down
    */
    if float32(upTrendCounter)*1.20 > float32(downTrendCounter) {
      upTrend = true
      downTrend = false
      fmt.Println("Uptrend Predicted")
    } else if float32(downTrendCounter)*1.20 > float32(upTrendCounter) {
      downTrend = true
      upTrend = false
      fmt.Println("Downtrend Predicted")
    } else {
      fmt.Println("Not Significant Difference To Determine a Trend")
      downTrend = false
      upTrend = false
    }

      currentPrice := getLastPrice("btc")

      fmt.Println("Current Price:", currentPrice)
      fmt.Println("Last Price", lastPrice)
      if upTrend {
        if currentPrice > lastPrice {
          fmt.Println("Uptrend True")
          correct++
        } else if lastPrice > currentPrice {
          fmt.Println("Uptrend False")
          incorrect++
        } else {
          fmt.Println("They are equal")
        }

      } else if downTrend {
        if currentPrice < lastPrice {
          fmt.Println("Downtrend True")
          correct++
        } else if lastPrice < currentPrice {
          fmt.Println("Downtrend False")
          incorrect++
        } else {
          fmt.Println("THey are equal")
        }
      } else {
        fmt.Println("Neither Up Nor Down Was Found")
      }
      // time.Sleep(10 * time.Second)

    totalCorrect := incorrect+correct
    var percentCorrect float32
    percentCorrect = float32(correct)/float32(totalCorrect)
    fmt.Println("Percent Correct:", percentCorrect)
    // time.Sleep(40 * time.Second)
  }
}


func getMostRecentOrders(orders *orderBook) {
  response, err := http.Get("https://api.bitfinex.com/v1/book/btcusd")
  if err != nil {
    fmt.Println(err)
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  }
  data := bytes.TrimSpace(body)

  data = bytes.TrimPrefix(data, []byte("// "))

  err = json.Unmarshal(data, &orders)
  checkErr(err)

}

func checkErr(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
