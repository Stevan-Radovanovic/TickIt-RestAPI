package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../database"
	"../models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetUserByID route
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	params := mux.Vars(r)
	fmt.Println(params)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println(id)
	collection := database.Client.Database("tick-it").Collection("users")

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found user")
	json.NewEncoder(w).Encode(user)
}

//GetUserByEmail route
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	params := mux.Vars(r)
	email := params["email"]

	collection := database.Client.Database("tick-it").Collection("users")

	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found user")
	json.NewEncoder(w).Encode(user)
}

//GetUsers route
func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := database.Client.Database("tick-it").Collection("users")
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

//DeleteUser route
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := database.Client.Database("tick-it").Collection("users")

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(deleteResult)
}
