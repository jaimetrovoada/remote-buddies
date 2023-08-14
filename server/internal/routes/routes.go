package routes

import (
	"remote-buddies/server/internal/db"
	"remote-buddies/server/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewRouter(db *db.Queries) *fiber.App {
	handlers := handlers.NewHandler(db)
	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api.Post("/users/:user/location", handlers.UserLocationHandler)
	api.Get("/nearby", handlers.NearbyHandler)
	return app
}
