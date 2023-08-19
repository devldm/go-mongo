package main

import (
	"go-mongo/database"
	"go-mongo/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a Fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on port 3000
	app.Listen(":3000")
}
