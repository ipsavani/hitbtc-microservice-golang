package services

import (
	"encoding/json"
	"fmt"
	"go_micro/config"
	"go_micro/models"
	"io/ioutil"
	"log"
	"net/http"
)

func getSymbolData(symbolID string) models.SymbolData {
	resp, err := http.Get(fmt.Sprintf(config.GetSymbolURL()+"%s", symbolID))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var symbol models.SymbolData
	err = json.Unmarshal(body, &symbol)
	if err != nil {
		log.Fatalln(err)
	}

	return symbol
}

func getCurrencyData(currencyID string) models.CurrencyData {
	resp, err := http.Get(fmt.Sprintf(config.GetCurrencyURL()+"%s", currencyID))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var currency models.CurrencyData
	err = json.Unmarshal(body, &currency)
	if err != nil {
		log.Fatalln(err)
	}

	return currency
}

func GetTradeData(symbol string) models.TradeData {
	symbolData := getSymbolData(symbol)
	currencyData := getCurrencyData(symbolData.BaseCurrency)
	tradeData := models.TradeData{
		Symbol:      symbol,
		SymbolID:    symbolData.BaseCurrency,
		FeeCurrency: symbolData.FeeCurrency,
		FullName:    currencyData.FullName,
	}
	return tradeData
}
