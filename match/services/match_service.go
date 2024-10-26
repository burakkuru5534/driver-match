package services

import (
	"context"
	"time"

	"driver-location-match/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchingService struct {
	db *mongo.Collection
}

func NewMatchingService(db *mongo.Database) *MatchingService {
	return &MatchingService{db: db.Collection("drivers")}
}

func (s *MatchingService) FindNearestDrivers(coordinates []float64, radius float64) ([]models.Driver, error) {
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
