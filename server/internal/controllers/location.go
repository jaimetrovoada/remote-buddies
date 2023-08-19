package controllers

import (
	"context"
	"log"
	"net/http"
	"remote-buddies/server/internal/db"
	"remote-buddies/server/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func (s *Service) NearbyHandler(c echo.Context) error {
	radius := 100
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	radius, err := strconv.Atoi(c.QueryParam("radius"))
	if err != nil {
		radius = 100
	}

	params := new(db.ListNearbyUsersParams)
	params.StMakepoint = lat
	params.StMakepoint_2 = lon
	params.StDwithin = radius
	locs, err := s.db.ListNearbyUsers(context.Background(), *params)

	if err != nil {
		log.Fatalf("Error reading locations: %v", err)
	}
	log.Println(locs)
	return c.JSON(http.StatusOK, locs)
}

type UserLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (s *Service) UserLocationHandler(c echo.Context) error {
	userLocation := new(UserLocation)
	r := c.Request()
	if err := c.Bind(userLocation); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaims)

	email := pgtype.Text{}
	email.Scan(claims.Email)

	params := new(db.UpdateUserLocationParams)
	params.StPoint = userLocation.Lat
	params.StPoint_2 = userLocation.Lon
	params.Email = email

	if err := s.db.UpdateUserLocation(r.Context(), *params); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
