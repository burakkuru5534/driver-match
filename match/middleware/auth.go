package middleware

import (
	"match-service/utils"
	"net/http"
)

// Middleware to authenticate requests
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if user is authenticated
		if authenticated, ok := claims["authenticated"].(bool); !ok || !authenticated {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r) // Continue to the next handler
	}
}
