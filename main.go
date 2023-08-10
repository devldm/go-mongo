package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go-mongo/controllers"
	"go-mongo/database"
	"go-mongo/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mg models.MongoInstance

func setupRoutes(app *fiber.App) {
	app.Get("/listings", controllers.GetAllListings)
	app.Get("/listings/:id", controllers.GetById)

}

func initDatabase() models.MongoInstance {
	mg, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	return mg
}

func main() {
	mg = initDatabase()
	defer mg.Db.Client().Disconnect(context.TODO())

	app := fiber.New()

	

	app.Listen(":3000")
}
