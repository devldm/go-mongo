package models

type Address struct {
	STREET   string   `json:"street" bson:"street"`
	COUNTRY  string   `json:"country" bson:"country"`
	LOCATION Location `json:"location" bson:"location"`
}

type Location struct {
	COORDINATES []float64 `json:"coordinates" bson:"coordinates"`
}
