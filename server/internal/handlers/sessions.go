package handlers

import (
	"net/http"
	"remote-buddies/server/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (s *Service) UserSessionsHandler(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaims)
	name := claims.Name
	return c.JSON(http.StatusOK, map[string]string{"name": name})
}
