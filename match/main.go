package main

import (
	controllers "driver-location-match/contollers"
	"driver-location-match/database"
	"driver-location-match/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	db, err := database.Connect("mongodb://localhost:27017", "driver_location", "drivers")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	matchingService := services.NewMatchingService(db.Database())
	matchingController := controllers.NewMatchingController(matchingService)

	r := mux.NewRouter()
	//r.Use(middleware.AuthMiddleware)
	r.HandleFunc("/drivers/nearby", matchingController.FindNearestDrivers).Methods("GET")

	port := "8081"
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
