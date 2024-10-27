package repository

import (
	"context"
	"location-service/models"
	"location-service/utils"
	"time"
)

// Define SaveLocation as a variable function
var SaveLocation = func(location models.DriverLocation) error {
	collection := utils.GetMongoCollection("drivers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, location)
	return err
}
