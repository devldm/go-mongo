package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Listing struct {
	ID                    string               `bson:"_id"`
	LISTING_URL           string               `json:"listing_url" bson:"listing_url"`
	NAME                  string               `json:"name" bson:"name"`
	SUMMARY               string               `json:"summary" bson:"summary"`
	SPACE                 string               `json:"space" bson:"space"`
	NEIGHBORHOOD_OVERVIEW string               `json:"neighborhood_overview" bson:"neighborhood_overview"`
	NOTES                 string               `json:"notes" bson:"notes"`
	BEDS                  int                  `json:"beds" bson:"beds"`
	BEDROOMS              int                  `json:"bedrooms" bson:"bedrooms"`
	BATHROOMS             primitive.Decimal128 `json:"bathrooms" bson:"bathrooms"`
}
