package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Zone Struct
type Zone struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Amount string             `json:"amount" bson:"amount,omitempty"`
}

//Event Model
type Event struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Date        string             `json:"date" bson:"date,omitempty"`
	Zones       []Zone             `json:"zones" bson:"zones,omitempty"`
}

//User Model
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

//Order Model
type Order struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email  string             `json:"email,omitempty" bson:"email,omitempty"`
	Ticket string             `json:"ticket" bson:"ticket,omitempty"`
	Date   string             `json:"date" bson:"date,omitempty"`
	Amount string             `json:"amount" bson:"amount,omitempty"`
}

var client *mongo.Client

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := client.Database("tick-it").Collection("users")
	findOptions := options.Find()
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var users []User

	for cur.Next(context.TODO()) {

		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(users, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Printf("Found users")
	json.NewEncoder(w).Encode(users)
}

func main() {
	fmt.Println("Tick-It Card Service Starting...")

	client = connect()

	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func connect() *mongo.Client {
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
	return client
}
