package main

import (
  "time"
  //"fmt"
)

func initializeArrayValues(currency string, waitTime uint8, array *[100]float32) {
  var newPrice float32
  var lengthOfArrayToModify int = 25

  for i := 0; i < lengthOfArrayToModify; i++ {
    newPrice = getLastPrice(currency)
    array[i] = newPrice
    time.Sleep(time.Duration(waitTime) * time.Second)
  }

}
