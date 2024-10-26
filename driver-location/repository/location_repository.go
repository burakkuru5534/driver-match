package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"location-service/models"
	"location-service/utils"
	"math"
	"time"
)

func SaveLocation(location models.DriverLocation) error {
	collection := utils.GetMongoCollection("locations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, location)
	return err
}

// FindNearestDriver retrieves the nearest driver based on a specified location.
func FindNearestDriver(userLocation models.GeoJSON) (models.DriverLocation, error) {
	collection := utils.GetMongoCollection("drivers") // Assume locations store driver locations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var nearestDriver models.DriverLocation
	var drivers []models.DriverLocation

	// Fetch all driver locations from the database
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nearestDriver, err
	}
	defer cursor.Close(ctx)

	// Iterate through the results
	for cursor.Next(ctx) {
		var driver models.DriverLocation
		if err := cursor.Decode(&driver); err != nil {
			return nearestDriver, err
		}
		drivers = append(drivers, driver)
	}

	// Logic to find the nearest driver
	nearestDistance := math.MaxFloat64 // Initialize with a large number
	for _, driver := range drivers {
		distance := calculateDistance(userLocation.Coordinates[0], userLocation.Coordinates[1],
			driver.Location.Coordinates[0], driver.Location.Coordinates[1])
		if distance < nearestDistance {
			nearestDistance = distance
			nearestDriver = driver
			nearestDriver.Distance = distance
		}
	}

	if nearestDistance == math.MaxFloat64 {
		return nearestDriver, mongo.ErrNoDocuments // No drivers found
	}

	return nearestDriver, nil
}

// calculateDistance calculates the distance between two geographic points using the Haversine formula.
func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Radius of the Earth in kilometers

	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c // Distance in kilometers
}
