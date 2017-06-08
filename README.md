# cryptotrader
a cryptocurrency trader for bitfinex

What this does is it trades different cryptocurrencies in an attempt to make money off of a volatile market. It does this by looking at the last hours worth of prices and looking at the percent change in price, buying if its down and selling if its up

### Setup
To set up you have to put in your public and private/secret API key into 'config.py' in the appropriate spots. This information can be created [here](https://www.bitfinex.com/api).

Example Setup
````
#API KEY PUBLIC
'your public api key'
#API KEY SECRET
'your secret api key'
````

This enables the main program to use this information to access your account and trade

#### NOTE:
Make sure you have all permissions enabled when creating the API key

### How It Works

What this program does is after you give it the amount you want to trade with, it access the amount of coins you have currently in the different currencies in Bitfinex. From here it gathers a price from Bitfinex and then continues to every 5 minutes.

It looks at the past hours worth of data (12, 5min segments) and looks at the change in price, if the change in price is greater than a certain amount it will buy or sell accordingly.

I know this is super basic and still has a ton of bugs but when as I am working them out this provides a nice basic structure to develop on in the future.

### Final Note
DO NOT USE THIS!
It is woefully incomplete and because of that has a good chance of actually losing you money, I am not responsible/liable if you lose money as a result of this program. 

Gage Coprivnicar
