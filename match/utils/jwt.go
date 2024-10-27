package utils

import (
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key") // Make sure this is consistent in both services

// Updated ValidateToken to return userID and error
func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["iss"].(string) // Using "iss" as username or ID
		return userID, nil
	}

	return "", err
}
