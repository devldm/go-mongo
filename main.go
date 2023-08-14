package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go-mongo/database"
	"go-mongo/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const coll = "listingsAndReviews"

func main() {
	mg, err := database.ConnectDB()
	if err != nil {
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
		cursor, err := mg.Db.Collection(coll).Find(c.Context(), bson.D{{}}, options.Find().SetLimit(20))
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
	})

	app.Listen(":3000")
}
