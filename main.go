package main

import (
	"labora-api/config"
	"labora-api/controllers"
	"labora-api/services"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	err := services.EstablishDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/items", controllers.GetItems).Methods("GET")
	// router.HandleFunc("/items/details", controllers.GetDetails).Methods("GET")
	router.HandleFunc("/items/{id}", controllers.GetItem).Methods("GET")
	router.HandleFunc("/item", controllers.SearchItem).Methods("GET")
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")

	port := ":8000"
	if err := config.StartServer(port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
