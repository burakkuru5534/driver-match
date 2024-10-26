package controllers

import (
	"encoding/json"
	"location-service/services"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var creds services.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := services.AuthenticateUser(creds)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
