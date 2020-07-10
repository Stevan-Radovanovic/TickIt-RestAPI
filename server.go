package main

import (
	"fmt"
	"log"
	"net/http"

	"./database"
	"./routes"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Tick-It Card Service Starting...")

	database.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/users", routes.GetUsers).Methods("GET")
	r.HandleFunc("/users/id/{id}", routes.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{email}", routes.GetUserByEmail).Methods("GET")
	r.HandleFunc("/orders", routes.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{email}", routes.GetOrdersByEmail).Methods("GET")
	r.HandleFunc("/orders/id/{id}", routes.GetOrderByID).Methods("GET")
	r.HandleFunc("/events", routes.GetEvents).Methods("GET")
	r.HandleFunc("/events/id/{id}", routes.GetEventByID).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")
	r.HandleFunc("/orders/{id}", routes.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/events/{id}", routes.DeleteEvent).Methods("DELETE")
	r.HandleFunc("/orders", routes.CreateOrder).Methods("POST")
	r.HandleFunc("/events", routes.CreateEvent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
