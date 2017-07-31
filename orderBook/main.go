package main

type orderBooks []orderBook

type orderBook struct {
  Bids     []bid     `json:"bids"`
  Asks     []ask     `json:"asks"`
}

type bid struct {
  Price     string     `json:"price"`
  Amount    string     `json:"amount"`
  Time      string     `json:"timestamp"`
}

type ask struct {
  Price     string     `json:"price"`
  Amount    string     `json:"amount"`
  Time      string     `json:"timestamp"`
}

func main() {
  response, err = http.Get("https://api.bitfinex.com/v1/book/btcusd")
  if err != nil {
    fmt.Println(err)
  }
  defer response.Body.Close()

  // Read the data into a byte slice
  body, err = ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  }
  // Remove whitespace from response
  data := bytes.TrimSpace(body)

  // Remove leading slashes and blank space to get byte slice that can be unmarshaled from JSON
  data = bytes.TrimPrefix(data, []byte("// "))

  // Unmarshal the JSON byte slice to a predefined struct
  err = json.Unmarshal(data, &orderBook)
  if err != nil {
    fmt.Println(err)
    //no idea what most of this stuff does but its important
  }
  fmt.Println(orderBook.Bids[0])
}
