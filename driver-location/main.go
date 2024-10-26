package main

import (
	"driver-location-matching/controllers"
	"driver-location-matching/database"
	"driver-location-matching/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//godotenv.Load(".env")

	db, err := database.Connect("mongodb://localhost:27017", "driver_location", "drivers")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	service := services.NewDriverService(db)
	service.EnsureIndexes()
	controller := controllers.NewDriverController(service)

	router := mux.NewRouter()
	//router.Use(middleware.AuthMiddleware(service))

	router.HandleFunc("/import", controller.ImportDrivers).Methods("POST")
	router.HandleFunc("/nearest", controller.FindNearestDrivers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}
