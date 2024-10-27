package controllers

import (
	"encoding/json"
	"location-service/models"
	"location-service/services"
	"net/http"
)

func ImportDrivers(w http.ResponseWriter, r *http.Request) {
	var filePath models.FilePath
	if err := json.NewDecoder(r.Body).Decode(&filePath); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := services.AddDriverLocations(filePath.Path); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Drivers imported successfully"))
}
