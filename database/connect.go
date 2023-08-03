package database

import (
	"context"
	"go-mongo/config"
	"go-mongo/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (models.MongoInstance, error) {
	uri := config.Config("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	db := client.Database(config.Config("DB_NAME"))

	if err != nil {
		return models.MongoInstance{}, err
	}

	return models.MongoInstance{
		Client: client,
		Db:     db,
	}, nil
}
