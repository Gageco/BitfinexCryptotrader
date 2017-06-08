import time
import os
import requests
from random import randint
from datetime import datetime
import json
import base64
import hmac
import hashlib
from decimal import Decimal

#OPTIONS
artificialTrading = True #Trade on bitfinex or pretend trade
coinsToTrade = {'ethusd', 'etcusd', 'btcusd' , 'ltcusd', 'dshusd'} #coins on bitfinex to trade with , 'btcusd' , 'ltcusd', 'dshusd'
moneyInfo = {'moneyMax' : 0, 'usableMoney' : 0}
#END OPTIONS

###TO DO ###
# - make run as long as humanly possible
# - do some sort of trend thing
# - Make sure you buy abouve .1 of a coin
# - make artificialTrading  a thing again

###DONE ###
# - DONE I THINK: so moneymax isnt working how i want it to so look into fixing that
# - DONE: make acutally work with bitfinex
# - DONE (I THINK): SORT OUT TRADING BUG ~83
# - DONE: limit the buying of a single coin
# - DONE: make sure you sell above the average buying price
# - DONE: move all file stuff to a dict
# - DONE: generate the info in coinInfo programmatically using
# - DONE: remove the file writing junk
# - DONE: show total value usd

#####FUNCTIONS FOR SETTING UP BUYING AND SELLING#####
def updateBalances(): # see your balances.
    payload = {"request":"/v1/balances", "nonce":genNonce()}
    signed_payload = payloadPacker(payload)
    r = requests.post("https://api.bitfinex.com/v1/balances", headers=signed_payload, verify=True)
    rep = r.json()
    for item in rep:
        coin = item['currency']
        if coin != 'usd':
            coin = coin+'usd'
        try:
            coinInfo['funds'][coin] = item['amount']
        except KeyError:
            pass

def payloadPacker(payload): # packs and signs the payload of the request.

	j = json.dumps(payload)
	data = base64.standard_b64encode(j)

	h = hmac.new(API_SECRET, data, hashlib.sha384)
	signature = h.hexdigest()

	return {
		"X-BFX-APIKEY": API_KEY,
		"X-BFX-SIGNATURE": signature,
		"X-BFX-PAYLOAD": data
	}

def genNonce(): # generates a nonce, used for authentication.
	return str(long(time.time() * 1000000))


def placeOrder(amount, price, side, symbol, ord_type='exchange market', exchange='bitfinex'): # submit a new order.
	payload = {
		"request":"/v1/order/new",
		"nonce":genNonce(),
		"symbol":symbol,
		"amount":amount,
		"price":price,
		"exchange":exchange,
		"side":side,
		"type":ord_type
	}

	signed_payload = payloadPacker(payload)
	r = requests.post("https://api.bitfinex.com/v1/order/new", headers=signed_payload, verify=True)
	rep = r.json()

	try:
		print("Order ID: " + str(rep['order_id']))

	except:
		print str(rep['message']) + " AN ISSUE OCCURED"

	try:
		if rep['side'] == 'buy':
			sellingPrice = rep['price']
			if sellingPrice in coinInfo['costs'][symbol].keys():
				coinInfo['costs'][symbol][sellingPrice] = str(round(float(amount), 6)*2)[:-1] #coinInfo['costs']['etcusd'][sellingPrice] : amount sold
			else:
				coinInfo['costs'][symbol][sellingPrice] = str(round(float(amount), 6))[:-1]
	except:
		pass

#####PRICE GETTING#####
def generateSomePrices(): #generate some test prices because i am offline and cant ping bitfinex
    return randint(40, 50)

def getBitfinexPrices(coin): #pull the price of a coin from bitfinex
    url = 'https://api.bitfinex.com/v1/pubticker/'
    jsonURL = url + coin
    r = requests.get(jsonURL).json()
    return float(r['last_price'])

#####SELLING FUNCTIONS#####
def sell(coin, price):
    if artificialTrading == True: #Change to artificialTrading == True # to be able to have both artificialTrading and
    #realTrading in the same function for simple switching back and forth
        howMuchToSell(coin, price)

    if artificialTrading == False:
        toSell = str(howMuchToSell(coin, price))
        if round(float(toSell), 4) == round(float(0), 4):
            print "No Profit/No Money in " + coin
        else:
            price = str(price)
            placeOrder(toSell, price, 'sell', coin)
            moneyInfo['usableMoney'] += float(price)
        #the other code neccesary to sell on bitfinex

def howMuchToSell(coin, price):
    coinSell = 0
    tempdict = {}
    for item in coinInfo['costs'][coin]: #coinInfo( 'costs' : {'etcusd' : {'priceBoughtAt' : 'amount'}})
        if float(price) >= (float(item)*1.01): #make sure profit of at least .1%
            coinSell += round(float(coinInfo['costs'][coin][item]), 4)
        else:
            tempdict[item] = round(float(coinInfo['costs'][coin][item]), 4)

    print "Selling " + str(coinSell) + " " + str(coin) + " at " + str(price)
    #MAKE IT SO IT ROUNDS DOWN AND IS NOT ABOVE HOW MUCH I HAVE. OK SO THIS IS INCOMPREHENSIBLE
    #WHAT IT IS SAYING EFFECTIVELY IS TO GET THE AMOUNT THAT IS IN A COIN TO SELL INSTEAD OF JUST MAKING UP A NUMBER.
    # THIS WOULD LEAD BACK TO THE IDEA OF AN AVERAGE I THINK. SO JUST KEEP AWARES OF THIS
    #WHAT A PAIN
    #instead what i did was i rounded down significnatly
    coinInfo['costs'][coin] = tempdict
    return coinSell #whatever bitfinex needs to sell the coin

#####BUYING FUNCTIONS#####
def buy(coin, price):
    if artificialTrading == True: #change to artificialTrading == True when this thing is for real
        howMuchToBuy(coin, price)

    if artificialTrading == False:
        toBuy = str(howMuchToBuy(coin, price))
        if float(toBuy) != float(0):
            price = str(price)
            placeOrder(toBuy, price, 'buy', coin)
        #Then this is where the bitfinex stuff will be

def howMuchToBuy(coin, price):
    if moneyInfo['usableMoney'] >= moneyInfo['moneyMax']*.05 and limitAmountThatCanBeBought(coin, price): #dont trade when less than 5% of orignial investment
        moneyBuy = float(.20*moneyInfo['usableMoney']) #percent of money to trade in a given trade
        coinBuy = str(round(float(moneyBuy/price), 4))
        print "Buying " + str(coinBuy) + " " + str(coin) + " for $" + str(moneyBuy) + " at " + str(price)
        moneyInfo['usableMoney'] = moneyInfo['usableMoney']-moneyBuy
        return coinBuy #return whatever bitfinex needs to buy
    else:
        print "Short on funds/Too much of " + coin
        return 0

def limitAmountThatCanBeBought(coin, price):
    maxOfOneCoin = .30 * moneyInfo['moneyMax'] #buy only n percent of moneyInfo['moneyMax'] for a single coin
    amountOfCoin = round(float(coinInfo['funds'][coin]), 4)*round(float(price))
    if amountOfCoin >= maxOfOneCoin:
        return False
    else:
        return True

#####PRICE ANALYSIS FUNCTIONS#####
def tenPercentChange(coinName):
    listLength = len(coinPriceList[coinName])
    timeInterval = 12 #1 hr default 12, to get the last hours worth of info
    lowerEndX = listLength - timeInterval
    for x in xrange(lowerEndX, listLength):#instead of zero use lowerEndX
        nowPrice = coinPriceList[coinName][listLength-1]
        lastPrice = coinPriceList[coinName][x]

        coinInfo['currentPrice'][coinName] = nowPrice
        percentChangeInPrice = float(float(nowPrice) - float(lastPrice))/float(nowPrice)    #ok so just realized this isnt even the right equation for calculating percent change, its (last - now)/now
        if percentChangeInPrice >= .01: # 2percent +
            sell(coinName, nowPrice)
            break
        if percentChangeInPrice <= -.02: # 1 percent -
            buy(coinName, nowPrice)
            break

#####MISC FUNCTIONS#####
def generateCoinInfo():
    #working with coinInfo
    for coin in coinsToTrade:
        for option in coinInfo:
            coinInfo[option][coin] = 0
    coinInfo['funds']['usd'] = moneyInfo['moneyMax']

    #working with coinPriceDict and tickers
    for coin in coinsToTrade:
        coinPriceList[coin] = []
        coinInfo['costs'][coin] = {}


#####MAIN FUNCTIONS#####

def main():
    for coin in coinsToTrade:
        tenPercentChange(coin)

    print coinInfo['funds']

def cryptoprice():
    for coin in coinsToTrade:
        ###CHECK WHICH TYPE OF TRADING IS TO BE USED###
        if artificialTrading == True:
            price = generateSomePrices()
        if artificialTrading == False:
            price = getBitfinexPrices(coin)
        ###CHECK WHICH TYPE OF TRADING IS TO BE USED###
        coinPriceList[coin].append(price)

fp = open("./config.py")
lines = fp.readlines()
fp.close()
API_KEY = eval(lines[1]) # put your API public key here.
API_SECRET = eval(lines[3]) # put your API private key here.

#Dicts and variables
coinPriceList = {}
coinInfo = {'funds' : {}, 'averages' : {}, 'currentPrice' : {}, 'costs' : {}}
totalUSD = 0

artTrad = "t"

timeBetweenTrades = int(raw_input("How long would you like to wait between trades? (milliseconds 300-5min 900-15min): "))
while True:
    artTrad = raw_input("artificialTrading? (T/F): ").lower()
    if artTrad == "t":
        artificialTrading = True
        break
    if artTrad == "f":
        artificialTrading = False
        break
    else:
        print "Invalid"

moneyInfo['moneyMax'] = int(raw_input("How much would you like to trade with?: "))
moneyInfo['usableMoney'] = moneyInfo['moneyMax']
generateCoinInfo()
print "Initializing..."
for x in xrange(0,12): #get some intial values for the coins
    cryptoprice()
    time.sleep(timeBetweenTrades)
    print 12-x + " Times Left"


while True:
    cryptoprice()
    main()
    print coinInfo['costs']
    updateBalances()
    time.sleep(timeBetweenTrades) #300 secs of five min 900-15min
    print ""
