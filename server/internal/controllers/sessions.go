package controllers

import (
	"net/http"
	"remote-buddies/server/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

func (s *Service) UserSessionsHandler(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaims)

	res := Response{
		Name:  claims.Name,
		Email: claims.Email,
		Image: claims.Image,
	}

	return c.JSON(http.StatusOK, res)
}
