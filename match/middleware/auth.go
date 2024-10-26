package middleware

import (
	"context"
	"match-service/utils"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		userID, err := utils.ValidateToken(token[len("Bearer "):]) // Get the token without "Bearer "
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Store user ID in context
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
