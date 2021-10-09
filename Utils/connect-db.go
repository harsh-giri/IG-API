package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func ConnecttoDB() *mongo.Collection {

	// Set client options
	//change the URI according to your database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.insta(), clientOptions)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//DB collection address which we are going to use
	//available to functions of all scope
	collection := client.Database("Appointy").Collection("users")

	return collection
}