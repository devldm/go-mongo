package listingHandler

import (
	"context"

	"fmt"
	"go-mongo/database"
	"go-mongo/models"
	"log"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const coll = "listingsAndReviews"

func GetListings(c *fiber.Ctx) error {
	mg, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	cursor, err := mg.Db.Collection(coll).Find(c.Context(), bson.D{{}}, options.Find().SetLimit(40))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document were found")
		return err
	}

	var listings []models.Listing = make([]models.Listing, 0)

	if err := cursor.All(c.Context(), &listings); err != nil {
		return c.Status(500).SendString(err.Error())

	}
	return c.JSON(listings)
}

func GetListing(c *fiber.Ctx) error {
	mg, err := database.ConnectDB()
	id := c.Params("id")
	var result bson.M
	findOneErr := mg.Db.Collection(coll).FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&result)

	if findOneErr == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id %s\n", id)
		return err
	}
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
