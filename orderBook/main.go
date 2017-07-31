package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"
  // "github.com/wcharczuk/go-chart"
  // "strconv"
)

func main() {
  var orders orderBooks
  // var order orderBook
  for i := 0; i < 10; i++ {                                                     //get orderbook info 10 times in 6 sec intervals
    fmt.Printf("Getting Orderbook Data. %d more seconds\n", (10-i)*6)
    getMostRecentOrders(&orders[i])

    time.Sleep(6 * time.Second)
  }
  fmt.Println(orders)
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
  if err != nil {
    fmt.Println(err)
  }
}
