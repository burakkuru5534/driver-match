package services

import (
	"github.com/sony/gobreaker"
	"location-service/models"
	"location-service/repository"
)

var cb *gobreaker.CircuitBreaker

func init() {
	settings := gobreaker.Settings{
		Name:    "DriverService",
		Timeout: 5,
	}
	cb = gobreaker.NewCircuitBreaker(settings)
}

// FindNearestDriver retrieves the nearest driver based on the user's location.
func FindNearestDriver(userLocation models.GeoJSON) (models.DriverLocation, error) {
	var driver models.DriverLocation
	result, err := cb.Execute(func() (interface{}, error) {
		return repository.FindNearestDriver(userLocation) // Call to location repository
	})
	if err != nil {
		return driver, err
	}
	driver = result.(models.DriverLocation)
	return driver, nil
}

// SaveDriver saves the driver information to the database.
func SaveDriver(driver models.DriverLocation) error {
	err := repository.SaveDriver(driver) // Call to the driver repository
	return err
}

// GetDriver retrieves a driver by ID.
func GetDriver(driverID string) (models.DriverLocation, error) {
	driver, err := repository.GetDriver(driverID) // Call to the driver repository
	return driver, err
}
