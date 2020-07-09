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

//GetEvents route
func GetEvents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := database.Client.Database("tick-it").Collection("sportevents")
	findOptions := options.Find()
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var events []models.Event

	for cur.Next(context.TODO()) {

		var elem models.Event
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		events = append(events, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Println("Found events")
	json.NewEncoder(w).Encode(events)
}

//GetEventByID route
func GetEventByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event models.Event
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := database.Client.Database("tick-it").Collection("sportevents")

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&event)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found event")
	json.NewEncoder(w).Encode(event)
}
