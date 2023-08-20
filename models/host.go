package models

type Host struct {
	HOST_NAME          string `json:"host_name" bson:"host_name"`
	HOST_THUMBNAIL_URL string `json:"host_thumbnail_url" bson:"host_thumbnail_url"`
}
