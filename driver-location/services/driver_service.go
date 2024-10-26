package services

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"driver-location-matching/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverService struct {
	db *mongo.Collection
}

func NewDriverService(db *mongo.Collection) *DriverService {
	return &DriverService{db: db}
}

func (s *DriverService) AddDriverLocations(filePath string) error {
	file, err := os.Open("/Users/qlubturkiye/Downloads/Coordinates.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var drivers []interface{}

	// Skip the header row
	if _, err := reader.Read(); err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Ensure there are exactly two columns for latitude and longitude
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

		driver := models.Driver{
			ID: primitive.NewObjectID(),
			Location: models.GeoJSONPoint{
				Type:        "Point",
				Coordinates: []float64{longitude, latitude},
			},
		}

		drivers = append(drivers, driver)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = s.db.InsertMany(ctx, drivers)
	return err
}

func (s *DriverService) FindNearestDrivers(coordinates []float64, radius float64) ([]models.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates,
				},
				"$maxDistance": radius,
			},
		},
	}

	cursor, err := s.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var drivers []models.Driver
	for cursor.Next(ctx) {
		var driver models.Driver
		if err := cursor.Decode(&driver); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, cursor.Err()
}

func (s *DriverService) EnsureIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "location", Value: "2dsphere"}},
	}
	_, err := s.db.Indexes().CreateOne(ctx, indexModel)
	return err
}
