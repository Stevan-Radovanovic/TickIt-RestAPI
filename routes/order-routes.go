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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//GetOrderByID route
func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order models.Order
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := database.Client.Database("tick-it").Collection("orders")

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&order)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found order")
	json.NewEncoder(w).Encode(order)
}

//DeleteOrder route
func DeleteOrder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := database.Client.Database("tick-it").Collection("orders")

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(deleteResult)
}

//CreateOrder route
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order models.Order

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&order)

	// connect db
	collection := database.Client.Database("tick-it").Collection("orders")

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), order)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}
