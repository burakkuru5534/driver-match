package controllers

import (
	"encoding/json"
	"location-service/services"
	"net/http"
)

// @Summary Auth Endpoint
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body User true "Credentials credentials"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Router /auth [post]
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

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
