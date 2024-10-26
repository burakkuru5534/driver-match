package middleware

import (
	"net/http"

	"driver-location-matching/services"
)

func AuthMiddleware(service *services.DriverService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			auth, err := service.Authenticate(token)
			if err != nil || !auth {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
