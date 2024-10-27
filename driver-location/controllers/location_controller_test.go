// controllers/controller_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"location-service/models"
	"location-service/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateLocation_Success simulates a successful database save operation.
func TestCreateLocation_Success(t *testing.T) {
	// Temporarily override SaveLocation with a mock that simulates a successful save.
	originalSaveLocation := repository.SaveLocation
	repository.SaveLocation = func(location models.DriverLocation) error {
		return nil
	}
	defer func() { repository.SaveLocation = originalSaveLocation }() // Restore original after test

	// Create a valid DriverLocation with GeoJSON data.
	location := models.DriverLocation{
		Location: models.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-122.4194, 37.7749},
		},
	}
	payload, _ := json.Marshal(location)

	// Set up a request and recorder
	req, _ := http.NewRequest("POST", "/locations", bytes.NewBuffer(payload))
	recorder := httptest.NewRecorder()

	// Call CreateLocation
	CreateLocation(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("expected status %v; got %v", http.StatusCreated, status)
	}
}

// TestCreateLocation_SaveError simulates a database save error.
func TestCreateLocation_SaveError(t *testing.T) {
	// Temporarily override SaveLocation to simulate a database error
	originalSaveLocation := repository.SaveLocation
	repository.SaveLocation = func(location models.DriverLocation) error {
		return errors.New("database error")
	}
	defer func() { repository.SaveLocation = originalSaveLocation }() // Restore original after test

	// Create a valid DriverLocation with GeoJSON data.
	location := models.DriverLocation{
		Location: models.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-122.4194, 37.7749},
		},
	}
	payload, _ := json.Marshal(location)

	// Set up a request and recorder
	req, _ := http.NewRequest("POST", "/locations", bytes.NewBuffer(payload))
	recorder := httptest.NewRecorder()

	// Call CreateLocation
	CreateLocation(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("expected status %v; got %v", http.StatusInternalServerError, status)
	}
}
