package routes

import (
	"remote-buddies/server/internal/config"
	"remote-buddies/server/internal/controllers"
	"remote-buddies/server/internal/db"
	"remote-buddies/server/internal/middleware"
	"remote-buddies/server/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(db *db.Queries) *echo.Echo {
	controllers := controllers.NewHandler(db)
	app := echo.New()

	app.Use(echoMiddleware.CORS())
	app.Use(echoMiddleware.Logger())
	app.Use(echoMiddleware.Recover())

	app.HTTPErrorHandler = middleware.CustomHTTPErrorHandler

	api := app.Group("/api")
	vConfig, _ := config.LoadConfig(".")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utils.JwtCustomClaims)
		},
		SigningKey: []byte(vConfig.JWT_SECRET),
	}

	api.POST("/users/location", controllers.UserLocationHandler, echojwt.WithConfig(config))
	api.GET("/nearby", controllers.NearbyHandler, echojwt.WithConfig(config))
	api.GET("/auth", controllers.AuthHandler)
	api.GET("/auth/callback", controllers.AuthCallbackHandler)
	api.GET("/sessions/user", controllers.UserSessionsHandler, echojwt.WithConfig(config))
	return app
}
