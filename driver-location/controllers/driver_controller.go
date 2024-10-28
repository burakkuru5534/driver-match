package controllers

import (
	"encoding/json"
	"location-service/models"
	"location-service/services"
	"net/http"
)

// @Summary GetNearestDriver Endpoint
// @Description GetNearestDriver  and returns the driver location and distance of the driver
// @Tags GetNearestDriver
// @Accept  json
// @Produce  json
// @Param userLocation body models.GeoJSON true "models.GeoJSON credentials"
// @Success 200 {object} models.DriverLocation
// @Failure 400 {object} error
// @Router /driver/nearest [post]
func GetNearestDriver(w http.ResponseWriter, r *http.Request) {
	var userLocation models.GeoJSON
	if err := json.NewDecoder(r.Body).Decode(&userLocation); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	nearestDriver, err := services.FindNearestDriver(userLocation)
	if err != nil {
		http.Error(w, "No drivers available", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(nearestDriver)
}
