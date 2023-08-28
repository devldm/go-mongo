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
	ACCOMMODATES          int                  `json:"accommodates" bson:"accommodates"`
	IMAGE                 Images               `json:"images" bson:"images"`
	ADDRESS               Address              `json:"address" bson:"address"`
	HOST                  Host                 `json:"host" bson:"host"`
	REVIEW_SCORE          ReviewScore          `json:"review_scores" bson:"review_scores"`
	NUMBER_OF_REVIEWS     int                  `json:"number_of_reviews" bson:"number_of_reviews"`
	REVIEWS               []ReviewData         `json:"reviews" bson:"reviews"`
	PRICE                 primitive.Decimal128 `json:"price" bson:"price"`
	ROOM_TYPE             string               `json:"room_type" bson:"room_type"`
	AMENITIES             []string             `json:"amenities" bson:"amenities"`
}
