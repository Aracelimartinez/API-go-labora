package main

import (
	"labora-api/config"
	"labora-api/controllers"
	"labora-api/services"
	"log"
	"github.com/rs/cors"
	// "github.com/gorilla/handlers"
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

	// Configurar el middleware CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5432"},
		AllowedMethods: []string{"GET", "POST"},
	})

	handler := corsOptions.Handler(router)

	port := ":8000"
	if err := config.StartServer(port, handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
