package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func init() {
	uri := "mongodb+srv://prakharjain496:golangdev@golangdb.v2pxtfw.mongodb.net/?retryWrites=true&w=majority"
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	dbClient = client

	log.Println("Connected to MongoDB")
}

func GetDBClient() *mongo.Client {
	return dbClient
}