package services

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	models2 "location-service/models"
	"location-service/utils"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddDriverLocations(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var drivers []interface{}

	// Read the CSV header
	if _, err := reader.Read(); err != nil {
		return err // Skip header
	}

	// Read each record
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Ensure we have the right number of fields
		if len(record) < 2 {
			return errors.New("invalid CSV format")
		}

		latitude, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return err
		}
		longitude, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return err
		}

		driver := models2.DriverLocation{
			ID: primitive.NewObjectID(),
			Location: models2.GeoJSON{
				Type:        "Point",
				Coordinates: []float64{longitude, latitude},
			},
		}

		drivers = append(drivers, driver)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := utils.GetMongoCollection("drivers")
	_, err = collection.InsertMany(ctx, drivers)
	return err
}
