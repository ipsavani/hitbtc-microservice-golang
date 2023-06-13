package services

import (
	"encoding/json"
	"log"

	"go_micro/models"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// var currencies map[string]models.Currency
// var mutex = &sync.Mutex{}

// Initialize WebSocket connection
func InitWebSocketConnection() (*websocket.Conn, error) {
	u := url.URL{Scheme: "wss", Host: "api.hitbtc.com", Path: "/api/2/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Subscribe to a ticker for a symbol
func SubscribeToTicker(c *websocket.Conn, symbol string) error {
	subscriptionMessage := map[string]interface{}{
		"method": "subscribeTicker",
		"params": map[string]string{
			"symbol": symbol,
		},
		"id": time.Now().Unix(),
	}
	return c.WriteJSON(subscriptionMessage)
}

// Process incoming WebSocket messages
func ProcessMessage(message []byte, symData models.TradeData) models.Currency {
	var result map[string]interface{}
	json.Unmarshal(message, &result)
	var currency models.Currency

	params, ok := result["params"].(map[string]interface{})
	if !ok {
		if result["result"] == true {
			log.Printf("Ticker Subscribed Succesfully")
		} else {
			log.Println("Failed to parse params")
		}
		return models.Currency{}
	}
	if params != nil {
		currency = models.Currency{
			ID:          symData.SymbolID,
			FullName:    symData.FullName,
			FeeCurrency: symData.FeeCurrency,
			Ask:         params["ask"].(string),
			Bid:         params["bid"].(string),
			Last:        params["last"].(string),
			Open:        params["open"].(string),
			Low:         params["low"].(string),
			High:        params["high"].(string),
		}
	}
	return currency
}
