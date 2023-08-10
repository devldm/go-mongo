package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongo/database"
	"go-mongo/models"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const coll = "listingsAndReviews"

func GetAllListings(c *fiber.Ctx) error {
	mg, err := database.ConnectDB()
	cursor, err := mg.Db.Collection(coll).Find(c.Context(), bson.D{{}}, options.Find().SetLimit(6))
	if err != nil {
		return err
		//return c.Status(500).SendString(err.Error())
	}
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document were found")
		return err
	}

	var listings []models.Listing = make([]models.Listing, 0)

	if err := cursor.All(c.Context(), &listings); err != nil {
		return err
		//return c.Status(500).SendString(err.Error())

	}
	return c.JSON(listings)
}

func GetById(c *fiber.Ctx) error {
	mg, err := database.ConnectDB()
	id := c.Params("id")
	var result bson.M
	mg.Db.Collection(coll).FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&result)
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
}
