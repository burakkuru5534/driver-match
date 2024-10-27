// services/driver_service_test.go
package services

import (
	"errors"
	"location-service/models"
	"location-service/repository"
	"testing"
)

// Mock variables to store original functions for restoring after tests
var (
	originalFindNearestDriver = repository.FindNearestDriver
)

// TestFindNearestDriver_Success tests the successful retrieval of the nearest driver.
func TestFindNearestDriver_Success(t *testing.T) {
	// Mock the FindNearestDriver function to return a dummy driver location
	repository.FindNearestDriver = func(userLocation models.GeoJSON) (models.DriverLocation, error) {
		return models.DriverLocation{
			Distance: 800,
			Location: models.GeoJSON{
				Type:        "Point",
				Coordinates: []float64{-122.4194, 37.7749},
			},
		}, nil
	}
	defer func() { repository.FindNearestDriver = originalFindNearestDriver }() // Restore original after test

	// Set up a sample user location
	userLocation := models.GeoJSON{
		Type:        "Point",
		Coordinates: []float64{-122.4194, 37.7749},
	}

	// Call FindNearestDriver
	driver, err := FindNearestDriver(userLocation)
	if err != nil {
		t.Errorf("expected no error; got %v", err)
	}

	// Validate driver location
	if driver.Distance > 5000 {
		t.Errorf("expected distance 1.5; got %v", driver.Distance)
	}
}

// TestFindNearestDriver_Error tests the error case for FindNearestDriver.
func TestFindNearestDriver_Error(t *testing.T) {
	// Mock the FindNearestDriver function to return an error
	repository.FindNearestDriver = func(userLocation models.GeoJSON) (models.DriverLocation, error) {
		return models.DriverLocation{}, errors.New("no drivers found")
	}
	defer func() { repository.FindNearestDriver = originalFindNearestDriver }() // Restore original after test

	// Set up a sample user location
	userLocation := models.GeoJSON{
		Type:        "Point",
		Coordinates: []float64{-122.4194, 37.7749},
	}

	// Call FindNearestDriver and expect an error
	_, err := FindNearestDriver(userLocation)
	if err == nil {
		t.Error("expected error; got nil")
	}
}
