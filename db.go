package main

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var menuCollection *mongo.Collection

// InitMongoDB initializes the MongoDB connection
func InitMongoDB() {
    // Connect to MongoDB
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    // Set up the collection
    menuCollection = client.Database("cafeteria").Collection("menu")
}

// CreateMenuItem inserts a new menu item into MongoDB
func CreateMenuItem(item MenuItem) error {
    _, err := menuCollection.InsertOne(context.TODO(), item)
    return err
}

// GetMenuItems retrieves all menu items from MongoDB
func GetMenuItems() ([]MenuItem, error) {
    var items []MenuItem

    cursor, err := menuCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        return nil, err
    }

    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var item MenuItem
        err := cursor.Decode(&item)
        if err != nil {
            return nil, err
        }
        items = append(items, item)
    }

    return items, nil
}
