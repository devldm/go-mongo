package router

import (
	listingRoutes "go-mongo/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")
	listingRoutes.SetupListingRoutes(api)

}
