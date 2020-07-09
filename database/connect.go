package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client for communication with MongoDB
var Client *mongo.Client

//Connect method
func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://stevan:Stevan.1@tickitcluster-trhkx.mongodb.net/tick-it?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Client = client
}
