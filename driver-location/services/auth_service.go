package services

import (
	"errors"
	"location-service/utils"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthenticateUser(creds Credentials) (string, error) {
	// Mock authentication
	if creds.Username == "driver" && creds.Password == "password" {
		return utils.GenerateToken(creds.Username)
	}
	return "", errors.New("invalid credentials")
}
