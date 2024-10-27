// services/match_service_test.go
package services

import (
	"encoding/json"
	"errors"
	"match-service/models"
	"match-service/utils"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetNearestDriver_Success(t *testing.T) {
	// Define the original circuit breaker function and restore after test
	originalRequestWithCircuitBreaker := utils.RequestWithCircuitBreaker
	defer func() { utils.RequestWithCircuitBreaker = originalRequestWithCircuitBreaker }()

	// Mock server to simulate Location Service
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		driver := models.Driver{
			Distance: 800,
			Location: models.GeoJSONPoint{
				Type:        "Point",
				Coordinates: []float64{20.4194, 40.7749},
			},
		}
		json.NewEncoder(w).Encode(driver)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Mock circuit breaker function to call the test server URL
	utils.RequestWithCircuitBreaker = func(req *http.Request) (*http.Response, error) {
		// Parse server URL
		testURL, err := url.Parse(server.URL + "/driver/nearest")
		if err != nil {
			return nil, err
		}
		req.URL = testURL // Set request URL to test server
		return http.DefaultClient.Do(req)
	}

	// Test input
	userLocation := models.GeoJSONPoint{
		Type:        "Point",
		Coordinates: []float64{20.4194, 40.7749},
	}
	token := "Bearer mock-token"

	// Call GetNearestDriver
	driver, err := GetNearestDriver(userLocation, token)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Verify the response
	if driver.Distance > 5000 {
		t.Errorf("expected distance 1.5, got %v", driver.Distance)
	}
}

func TestGetNearestDriver_CircuitBreakerOpen(t *testing.T) {
	// Define the original circuit breaker function and restore after test
	originalRequestWithCircuitBreaker := utils.RequestWithCircuitBreaker
	defer func() { utils.RequestWithCircuitBreaker = originalRequestWithCircuitBreaker }()

	// Mock circuit breaker to return an error
	utils.RequestWithCircuitBreaker = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("circuit breaker open")
	}

	// Test input
	userLocation := models.GeoJSONPoint{
		Type:        "Point",
		Coordinates: []float64{-122.4194, 37.7749},
	}
	token := "Bearer mock-token"

	// Call GetNearestDriver and expect an error
	_, err := GetNearestDriver(userLocation, token)
	if err == nil || err.Error() != "circuit breaker triggered or service unavailable: circuit breaker open" {
		t.Errorf("expected circuit breaker error, got %v", err)
	}
}

func TestGetNearestDriver_Non200Response(t *testing.T) {
	// Define the original circuit breaker function and restore after test
	originalRequestWithCircuitBreaker := utils.RequestWithCircuitBreaker
	defer func() { utils.RequestWithCircuitBreaker = originalRequestWithCircuitBreaker }()

	// Mock server to simulate Location Service with a non-200 status
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Mock circuit breaker function to call the test server URL
	utils.RequestWithCircuitBreaker = func(req *http.Request) (*http.Response, error) {
		// Parse server URL
		testURL, err := url.Parse(server.URL + "/driver/nearest")
		if err != nil {
			return nil, err
		}
		req.URL = testURL // Set request URL to test server
		return http.DefaultClient.Do(req)
	}

	// Test input
	userLocation := models.GeoJSONPoint{
		Type:        "Point",
		Coordinates: []float64{-122.4194, 37.7749},
	}
	token := "Bearer mock-token"

	// Call GetNearestDriver and expect an error
	_, err := GetNearestDriver(userLocation, token)
	expectedErr := "failed to get nearest driver, status code: 404"
	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}
