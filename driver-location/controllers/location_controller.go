package controllers

import (
	"encoding/json"
	"location-service/models"
	"location-service/repository"
	"net/http"
)

func CreateLocation(w http.ResponseWriter, r *http.Request) {
	var location models.DriverLocation
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := repository.SaveLocation(location)
	if err != nil {
		http.Error(w, "Failed to save location", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
