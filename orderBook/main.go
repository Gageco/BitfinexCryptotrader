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

func main() {
  var orders orderBooks
  var upTrend bool
  var downTrend bool
  var bidPrice float64
  var askPrice float64
  var err error
  correct := 0
  incorrect := 0

  // var order orderBook
  for i := 0; i < 10; i++ {                                                     //get orderbook info 10 times in 6 sec intervals
    fmt.Printf("Getting Orderbook Data. %d more seconds\n", (10-i)*6)
    getMostRecentOrders(&orders[i])

    time.Sleep(6 * time.Second)
  }
  fmt.Print("Bids:    Amount:", orders[0].Bids[0].Amount)
  fmt.Println(";    Price:", orders[0].Bids[0].Price)

  fmt.Print("Asks:    Amount:", orders[0].Asks[0].Amount)
  fmt.Println(";    Price:", orders[0].Asks[0].Price)

  for i := 0; i < 10; i++ {
    bidPrice, err = strconv.ParseFloat(orders[i].Bids[0].Amount, 32)
    checkErr(err)
    askPrice, err = strconv.ParseFloat(orders[i].Asks[0].Amount, 32)
    checkErr(err)

    if bidPrice > askPrice {         //more bids than asks, more buys than sells
      upTrend = true
      downTrend = false
      fmt.Println("Uptrend Predicted")
    } else if bidPrice < askPrice {  //more asks than bids, more sells than buys
      upTrend = false
      downTrend = true
      fmt.Println("Downtrend Predicted")
    } else {                                                                                           //if they are equal (dont think this is possible)
      upTrend = false
      downTrend = false
      fmt.Println("Neither Up Nor Down")
    }


    lastPrice, _ := strconv.ParseFloat(orders[i+1].Asks[0].Amount, 32)

    if upTrend && !downTrend {
      if lastPrice > bidPrice {
        fmt.Println("Uptrend True")
        correct++
      } else {
        fmt.Println("Uptrend False")
        incorrect++
      }
    } else if !upTrend && downTrend {
      if lastPrice < bidPrice {
        fmt.Println("Downtrend True")
        correct++
      } else {
        fmt.Println("Downtrend False")
        incorrect++
      }
    } else {
      fmt.Println("Neither Up Nor Down Was Found")
    }
    time.Sleep(10 * time.Second)
  }
  totalCorrect := incorrect+correct
  fmt.Println("Percent Correct:", correct/totalCorrect)
  time.Sleep(40 * time.Second)
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
