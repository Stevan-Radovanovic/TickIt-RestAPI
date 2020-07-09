package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"

	"../database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetOrdersByEmail route
func GetOrdersByEmail(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	email := params["email"]

	collection := database.Client.Database("tick-it").Collection("orders")
	findOptions := options.Find()
	filter := bson.M{"email": email}

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var orders []models.Order

	for cur.Next(context.TODO()) {

		var elem models.Order
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		orders = append(orders, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Println("Found orders")
	json.NewEncoder(w).Encode(orders)
}

//GetOrders route
func GetOrders(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := database.Client.Database("tick-it").Collection("orders")
	findOptions := options.Find()
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var orders []models.Order

	for cur.Next(context.TODO()) {

		var elem models.Order
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		orders = append(orders, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Println("Found orders")
	json.NewEncoder(w).Encode(orders)
}
