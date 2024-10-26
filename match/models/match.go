package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeoJSONPoint struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Driver struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Location GeoJSONPoint       `bson:"location"`
}
