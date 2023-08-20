package models

type Address struct {
	STREET  string `json:"street" bson:"street"`
	COUNTRY string `json:"country" bson:"country"`
}
