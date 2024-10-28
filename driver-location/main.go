package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"location-service/controllers"
	"location-service/middleware"
	"location-service/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	utils.ConnectDB() // Connect to MongoDB

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/auth", controllers.Authenticate).Methods("POST")
	r.HandleFunc("/location", controllers.CreateLocation).Methods("POST")
	r.HandleFunc("/import", controllers.ImportDrivers).Methods("POST")
	r.HandleFunc("/driver/nearest", middleware.AuthMiddleware(controllers.GetNearestDriver)).Methods("POST")
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Location Service is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
