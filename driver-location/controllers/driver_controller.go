package controllers

import (
	"encoding/json"
	"location-service/models"
	"location-service/services"
	"net/http"
)

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
