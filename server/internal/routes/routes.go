package routes

import (
	"remote-buddies/server/internal/db"
	"remote-buddies/server/internal/handlers"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(db *db.Queries) *echo.Echo {
	handlers := handlers.NewHandler(db)
	app := echo.New()

	app.Use(echoMiddleware.CORS())
	app.Use(echoMiddleware.Logger())
	app.Use(echoMiddleware.Recover())

	// app.HTTPErrorHandler = middleware.CustomHTTPErrorHandler

	api := app.Group("/api")

	api.POST("/users/:user/location", handlers.UserLocationHandler)
	api.GET("/nearby", handlers.NearbyHandler)
	api.GET("/auth", handlers.AuthHandler)
	api.GET("/auth/callback", handlers.AuthCallbackHandler)
	return app
}
