package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./models"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func getUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	params := mux.Vars(r)
	email := params["email"]

	collection := client.Database("tick-it").Collection("users")

	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found user")
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := client.Database("tick-it").Collection("users")
	findOptions := options.Find()
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User

	for cur.Next(context.TODO()) {

		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Println("Found users")
	json.NewEncoder(w).Encode(users)
}

func main() {
	fmt.Println("Tick-It Card Service Starting...")

	client = connect()

	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{email}", getUserByEmail).Methods("GET")

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
