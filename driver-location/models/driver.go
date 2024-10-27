package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DriverLocation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Distance float64            `bson:"distance,omitempty"`
	Location GeoJSON            `bson:"location"`
}

type GeoJSON struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type FilePath struct {
	Path string `bson:"path"`
}
