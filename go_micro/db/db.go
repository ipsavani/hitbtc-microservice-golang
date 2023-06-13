package db

import (
	"context"
	"fmt"
	"go_micro/config"
	"go_micro/models"
	"go_micro/services"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = config.GetDBConnection()
var dbName string = config.GetDBName()
var collName string = config.GetCollectionName()

// Connect to the MongoDB
func connectDB() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

// Update a document in the database
func updateDocument(client *mongo.Client, ctx context.Context, update models.TradeData) {
	collection := client.Database(dbName).Collection(collName)
	filter := bson.M{"symbol": update.Symbol}
	updateBson := bson.M{"$set": update}
	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, updateBson, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %s data. \n", update.Symbol)
}

// Fetch a document from the database
func fetchDocument(client *mongo.Client, ctx context.Context, symbol string) models.TradeData {
	collection := client.Database(dbName).Collection(collName)
	filter := bson.M{"symbol": symbol}
	var result models.TradeData
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func StoreSymbolsToDB(symbols []string) {
	var wg sync.WaitGroup
	client, ctx := connectDB()
	defer client.Disconnect(ctx)

	for _, symbol := range symbols {
		var update models.TradeData
		update = services.GetTradeData(symbol)

		wg.Add(1)
		go func(update models.TradeData) {
			defer wg.Done()
			updateDocument(client, ctx, update)
		}(update)
	}
	wg.Wait()
}

func FetchSymbolFromDB(symbol string) models.TradeData {
	client, ctx := connectDB()
	defer client.Disconnect(ctx)

	doc := fetchDocument(client, ctx, symbol)
	return doc
}
