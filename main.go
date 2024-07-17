package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Ping failed", err)
	}

	fmt.Println("Connected successfully")

	collection := client.Database("work").Collection("work")

	doc := bson.D{{"name", "rahul"}, {"age", 25}}

	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println("Error inserting", err)

	}
	fmt.Println("insert succeeded", insertResult)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Disconnect(ctx)

	fmt.Println("Disconnected successfully")

}
