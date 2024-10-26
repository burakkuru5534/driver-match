package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetNearestDriverForUser retrieves the nearest driver for a given user ID.
func GetNearestDriverForUser(userID, token string) (interface{}, error) {
	// Simulate user location retrieval from your database or any source
	userLocation := map[string]float64{
		"latitude":  40.7128,  // Example latitude
		"longitude": -74.0060, // Example longitude
	}

	// Prepare the request to the location service
	locationServiceURL := "http://localhost:8081/driver/nearest"
	jsonData, err := json.Marshal(map[string]interface{}{
		"type":        "Point",
		"coordinates": []float64{userLocation["longitude"], userLocation["latitude"]},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal location data: %w", err)
	}

	req, err := http.NewRequest("POST", locationServiceURL+"/driver/nearest", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add Authorization header with Bearer token
	req.Header.Set("Authorization", token) // Use the passed token
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call location service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get nearest driver: %s", resp.Status)
	}

	var nearestDriver interface{}
	if err := json.NewDecoder(resp.Body).Decode(&nearestDriver); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return nearestDriver, nil
}
