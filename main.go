package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

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

var mg MongoInstance

const dbName = "sample_airbnb"

const coll = "listingsAndReviews"

func Connect() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Printf("Connecting to MongoDB at %s", uri)
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")

	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	app.Get("/listings/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var result bson.M
		err := mg.Db.Collection(coll).FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the id %s\n", id)
			return err
		}
		if err != nil {
			panic(err)
		}

		jsonData, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", jsonData)
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Get("/listings", func(c *fiber.Ctx) error {
		cursor, err := mg.Db.Collection(coll).Find(c.Context(), bson.D{{}}, options.Find().SetLimit(6))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document were found")
			return err
		}

		var listings []Listing = make([]Listing, 0)

		if err := cursor.All(c.Context(), &listings); err != nil {
			return c.Status(500).SendString(err.Error())

		}
		return c.JSON(listings)
	})

	app.Listen(":3000")
}
