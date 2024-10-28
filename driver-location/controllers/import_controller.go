package controllers

import (
	"encoding/json"
	"location-service/models"
	"location-service/services"
	"net/http"
)

// @Summary ImportDrivers Endpoint
// @Description ImportDrivers get filepath and import the csv to mongodb
// @Tags ImportDrivers
// @Accept  json
// @Produce  json
// @Param filePath body models.FilePath true "models.FilePath credentials"
// @Success 200 {object} models.FilePath
// @Failure 400 {object} error
// @Router /import [post]
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
