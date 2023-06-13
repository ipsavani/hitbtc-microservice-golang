package serviceController

import (
	"fmt"
	"go_micro/config"
	"go_micro/db"
	"go_micro/models"
	"go_micro/services"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

var currencies map[string]models.Currency
var mutex = &sync.Mutex{}

func startWebsocket(symbol string) {
	symData := db.FetchSymbolFromDB(symbol)

	c, err := services.InitWebSocketConnection()
	if err != nil {
		log.Fatal("Failed to dial:", err)
	}
	defer c.Close()

	err = services.SubscribeToTicker(c, symbol)
	if err != nil {
		log.Fatal("Failed to subscribe:", err)
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Failed to read:", err)
			return
		}
		mutex.Lock()
		currencies[symbol] = services.ProcessMessage(message, symData)
		mutex.Unlock()
	}
}

func Run() {
	currencies = make(map[string]models.Currency)
	symbols := config.GetSymbols()
	db.StoreSymbolsToDB(symbols)

	for _, symbol := range symbols {
		go startWebsocket(symbol)
	}
	router := gin.Default()
	router.GET("/currency/:symbol", getCurrency)
	router.GET("/currency/all", getAllCurrencies)

	router.Run(":8080")
}

func getCurrency(c *gin.Context) {
	symbol := c.Param("symbol")
	// mutex.Lock()
	currency, ok := currencies[symbol]
	// mutex.Unlock()
	if !ok {
		c.JSON(404, gin.H{"message": fmt.Sprintf("Currency with symbol %s not found", symbol)})
		return
	}

	c.JSON(200, currency)
}

func getAllCurrencies(c *gin.Context) {
	// mutex.Lock()
	allCurrencies := make([]models.Currency, 0, len(currencies))
	for _, currency := range currencies {
		allCurrencies = append(allCurrencies, currency)
	}
	// mutex.Unlock()

	c.JSON(200, gin.H{"currencies": allCurrencies})
}
