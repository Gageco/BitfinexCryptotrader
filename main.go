package main
// run this program with go run *.go
import (
  "fmt"
  "github.com/bitfinexcom/bitfinex-api-go/v1"
  "bufio"
  "os"
)


/*
TODO
- Check if order went through or not                                       --Done
- Get a price to sell at
- Check if this price is a good price to sell at (the lowest is above the buying price)
- make sure its above the lowest amount that can be bought and sold        --Done
- account for trading fees
- write trading algo
*/


func updatePriceArray(currency string, waitTime uint8, priceArray *[100]float32) {
  var lastPrice float32
  for i := len(priceArray)-1; i > 0; i-- {
    priceArray[i] = priceArray[i-1]                                             //moves all the values back one place
  }
  lastPrice = getHigh(currency)
  priceArray[0] = lastPrice

}

func main() {
  var priceArray[100]float32                                                      //to store the prices of the crypto
  var currency string = "eth"
  var timeToWait uint8 =  1                                                       //time to wait between trades in in seconds
  var buyingPrices[100]float32

  //bitfinex authenication
  inFile, _ := os.Open("./config")
  scanner := bufio.NewsScanner(inFile)
  scanner.Split(bufio.ScanLines)
  scanner.Scan()
  scanner.Scan()                                                                //skip first line
  key := scanner.Text()                                                         //second line holds key
  scanner.Scan()                                                                //skip thirdline
  scanner.Scan()
  secret := scanner.Text()                                                      //fourthline holds secret
  var client = bitfinex.NewClient().Auth(key, secret)
  inFile.Close()
  //initilization of junk
  initializeArrayValues(currency, timeToWait, &priceArray)                       //add 25 values to the array so that they can then be analyzed

  for i := 0; i < 10; i += 0 {
    updatePriceArray(currency, timeToWait, &priceArray)
    //basicTrading(currency, priceArray, client, &buyingPrices)
    //time.Sleep(time.Duration(timeToWait) * time.Minute)                       //sleep for minute after trade
  }
}
