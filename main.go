package main

import (
	"context"
	"fmt"
	"log"
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
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("error sstart", err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("work").Collection("work")

	// doc := bson.D{{"name", "Alice"}, {"age", 25}}
	// insertResult, err := collection.InsertOne(context.TODO(), doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted document with ID:", insertResult.InsertedID)

	// find single

	filter := bson.D{{"name", "sam"}}

	var result bson.M

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error handling", err)
		return
	}
	fmt.Println(result)

	update := bson.D{
		{"$set", bson.D{
			{"name", "yash"},
		}},
	}
	updateRes, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Error updating", err)
	}
	fmt.Println(updateRes.MatchedCount, updateRes.ModifiedCount)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Disconnect(ctx)
	if err != nil {
		fmt.Println("Error disconnect", err)
	}
	fmt.Println("Disconnected from MongoDB!")
}
