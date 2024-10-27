package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtSecret = []byte("your_secret_key") // Replace with your secret key

// User struct for demonstration purposes
type User struct {
	Username      string `json:"username"`
	Authenticated bool   `json:"authenticated"`
}

// Generate JWT token
func GenerateToken(user User) (string, error) {
	claims := jwt.MapClaims{
		"authenticated": user.Authenticated,
		"exp":           time.Now().Add(time.Hour * 72).Unix(), // Token valid for 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Validate JWT token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err // Return nil claims if token is invalid
	}

	return claims, nil
}
