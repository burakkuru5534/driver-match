package utils

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

// CircuitBreaker for handling Location Service requests
var cb *gobreaker.CircuitBreaker

func init() {
	// Configure the circuit breaker settings
	cbSettings := gobreaker.Settings{
		Name:        "LocationServiceCircuitBreaker",
		MaxRequests: 3,                // Max requests to allow when half-open
		Interval:    60 * time.Second, // Interval for reset
		Timeout:     10 * time.Second, // Timeout for the open state
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3 // Threshold to open circuit
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			fmt.Printf("Circuit breaker state changed: %s -> %s\n", from.String(), to.String())
		},
	}

	// Initialize the circuit breaker
	cb = gobreaker.NewCircuitBreaker(cbSettings)
}

// RequestWithCircuitBreaker sends an HTTP request with the circuit breaker protection
func RequestWithCircuitBreaker(req *http.Request) (*http.Response, error) {
	// Use the circuit breaker to execute the HTTP request
	result, err := cb.Execute(func() (interface{}, error) {
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		// Check for HTTP status codes that indicate service failure
		if resp.StatusCode >= 500 {
			return nil, errors.New("service unavailable")
		}
		return resp, nil
	})

	// Check if the result or error came from the circuit breaker
	if err != nil {
		fmt.Println("Circuit breaker error:", err)
		return nil, err
	}
	return result.(*http.Response), nil
}
