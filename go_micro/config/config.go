package config

//db config
var dburi string = "mongodb+srv://psavani:1234567890@cluster0.t4m7vqz.mongodb.net/?retryWrites=true&w=majority"
var dbName string = "hitbtcDatabase"
var collName string = "tradeData"

//symbols
var symbols []string = []string{"BTCUSD", "ETHBTC"}

//net/http config
var symbolurl string = "https://api.hitbtc.com/api/2/public/symbol/"
var currencyurl string = "https://api.hitbtc.com/api/2/public/currency/"

func GetSymbolURL() string {
	return symbolurl
}

func GetCurrencyURL() string {
	return currencyurl
}

func GetSymbols() []string {
	return symbols
}

func GetDBConnection() string {
	return dburi
}

func GetDBName() string {
	return dbName
}

func GetCollectionName() string {
	return collName
}
