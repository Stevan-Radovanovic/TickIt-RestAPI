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
	r.HandleFunc("/users/{email}", routes.GetUserByEmail).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
