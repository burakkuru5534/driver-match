package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"match-service/models"
	"match-service/utils"
	"net/http"
)

// GetNearestDriver requests the nearest driver from the Location Service.
func GetNearestDriver(userLocation models.GeoJSONPoint, token string) (models.Driver, error) {
	url := "http://localhost:8081/driver/nearest"
	requestBody, err := json.Marshal(userLocation)
	if err != nil {
		return models.Driver{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.Driver{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	// Use circuit breaker to make the request
	resp, err := utils.RequestWithCircuitBreaker(req)
	if err != nil {
		// Fallback response if circuit is open or request fails
		fmt.Println("Using fallback: No nearest driver available")
		return models.Driver{}, fmt.Errorf("circuit breaker triggered or service unavailable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Driver{}, fmt.Errorf("failed to get nearest driver, status code: %d", resp.StatusCode)
	}

	var nearestDriver models.Driver
	if err := json.NewDecoder(resp.Body).Decode(&nearestDriver); err != nil {
		return models.Driver{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return nearestDriver, nil
}
