package main

type orderBooks [lengthOfOrderBooks]orderBook                                                   //you can take samples 6 seconds apart 10 times

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
