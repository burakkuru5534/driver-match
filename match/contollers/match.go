package controllers

import (
	"encoding/json"
	"match-service/models"
	"match-service/services"
	"net/http"
)

// @Summary GetNearestDriverController Endpoint
// @Description GetNearestDriverController gets the nearest location of the driver from location service
// @Tags GetNearestDriverController
// @Accept  json
// @Produce  json
// @Param location body models.GeoJSONPoint true "models.GeoJSONPoint credentials"
// @Success 200 {object} models.Driver
// @Failure 400 {object} error
// @Router /match/nearest [post]
// GetNearestDriverController handles requests for the nearest driver.
func GetNearestDriverController(w http.ResponseWriter, r *http.Request) {
	// Parse user location from the request body
	var userLocation models.GeoJSONPoint
	if err := json.NewDecoder(r.Body).Decode(&userLocation); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve the token from the request headers
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return
	}

	// Call the service to get the nearest driver with the token
	nearestDriver, err := services.GetNearestDriver(userLocation, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the nearest driver location
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nearestDriver)
}
