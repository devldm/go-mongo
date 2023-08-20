package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReviewScore struct {
	REVIEW_SCORES_RATING int `json:"review_scores_rating" bson:"review_scores_rating"`
}

type ReviewData struct {
	_ID           primitive.ObjectID `bson:"_id"`
	COMMENTS      string             `json:"comments" bson:"comments"`
	DATE          primitive.DateTime `json:"date" bson:"date"`
	LISTING_ID    string             `json:"listing_id" bson:"listing_id"`
	REVIEWER_ID   string             `json:"reviewer_id" bson:"reviewer_id"`
	REVIEWER_NAME string             `json:"reviewer_name" bson:"reviewer_name"`
}
