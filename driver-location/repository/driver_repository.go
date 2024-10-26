package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"location-service/models"
	"location-service/utils"
	"time"
)

// SaveDriver saves the driver information to the database.
func SaveDriver(driver models.DriverLocation) error {
	collection := utils.GetMongoCollection("drivers") // Assuming a separate collection for drivers
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, driver)
	return err
}

// GetDriver retrieves a driver by ID.
func GetDriver(driverID string) (models.DriverLocation, error) {
	collection := utils.GetMongoCollection("drivers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var driver models.DriverLocation
	err := collection.FindOne(ctx, bson.M{"driver_id": driverID}).Decode(&driver)
	return driver, err
}
