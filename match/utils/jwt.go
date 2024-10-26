package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("secret_key") // Make sure this is consistent in both services

func GenerateToken(username string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

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
