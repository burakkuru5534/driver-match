package controllers

import (
	"encoding/json"
	"match-service/services"
	"net/http"
)

// GetNearestDriver handles the request to get the nearest driver for an authenticated user.
func GetNearestDriver(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string) // Extract userID from context after authentication
	token := r.Header.Get("Authorization")         // Get the token from the request header

	nearestDriver, err := services.GetNearestDriverForUser(userID, token) // Pass the token to the service
	if err != nil {
		http.Error(w, "No drivers available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nearestDriver)
}
