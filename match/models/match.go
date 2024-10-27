package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeoJSONPoint struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Driver struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Distance float64            `bson:"distance,omitempty"`
	Location GeoJSONPoint       `bson:"location"`
}

// UserLocation represents a user's location with latitude and longitude coordinates.
type UserLocation struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}
