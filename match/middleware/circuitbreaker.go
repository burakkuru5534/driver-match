package middleware

import (
	"github.com/sony/gobreaker"
	"net/http"
)

var breaker = gobreaker.NewCircuitBreaker(gobreaker.Settings{})

func CircuitBreaker(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := breaker.Execute(func() (interface{}, error) {
			next.ServeHTTP(w, r)
			return nil, nil
		})
		if err != nil {
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		}
	}
}
