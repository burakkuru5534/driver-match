package main

import (
	"log"
	controllers "match-service/contollers"
	"match-service/middleware"
	"net/http"
)

func main() {

	// Set up the router
	http.HandleFunc("/match/nearest", middleware.AuthMiddleware(controllers.GetNearestDriver))

	// Start the server
	port := "8082" // Set your port in the .env file
	log.Printf("Match service running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
