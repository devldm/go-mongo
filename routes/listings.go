package listingRoutes

import (
	listingHandler "go-mongo/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupListingRoutes(router fiber.Router) {
	note := router.Group("/listings")

	// Get all listings
	note.Get("/", listingHandler.GetListings)
	// Get one listing
	note.Get("/:id", listingHandler.GetListing)

	// TODO: use these to expand offering in future
	// // Create a listing
	// note.Post("/", listingHandler.CreateListing)
	// // // Update one listing
	// note.Put("/:id", listingHandler.UpdateListing)
	// // // Delete one listing
	// note.Delete("/:id", listingHandler.DeleteListing)
}
