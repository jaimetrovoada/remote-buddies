package handlers

import (
	"context"
	"log"
	"net/http"
	"remote-buddies/server/internal/db"

	"github.com/labstack/echo/v4"
)

func (s *Service) NearbyHandler(c echo.Context) error {
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	// _radius := c.Query("radius")
	params := new(db.ListLocationsParams)
	params.StMakepoint = lat
	params.StMakepoint_2 = lon
	locs, err := s.db.ListLocations(context.Background(), *params)

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

	params := new(db.CreateLocationParams)
	params.StPoint = userLocation.Lat
	params.StPoint_2 = userLocation.Lon

	s.db.CreateLocation(r.Context(), *params)

	log.Printf("lat: %v, lon: %v", userLocation.Lat, userLocation.Lon)

	return c.NoContent(http.StatusOK)
}
