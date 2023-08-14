package handlers

import (
	"context"
	"log"
	"remote-buddies/server/internal/db"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	db *db.Queries
}

func NewHandler(db *db.Queries) *Service {
	return &Service{db}
}

func (s *Service) NearbyHandler(c *fiber.Ctx) error {
	lat := c.Query("lat")
	lon := c.Query("lon")
	// _radius := c.Query("radius")
	params := new(db.ListLocationsParams)
	params.StMakepoint = lat
	params.StMakepoint_2 = lon
	locs, err := s.db.ListLocations(context.Background(), *params)

	if err != nil {
		log.Fatalf("Error reading locations: %v", err)
	}
	log.Println(locs)
	return c.JSON(locs)
}

type UserLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (s *Service) UserLocationHandler(c *fiber.Ctx) error {
	userLocation := new(UserLocation)

	if err := c.BodyParser(userLocation); err != nil {
		return err
	}

	params := new(db.CreateLocationParams)
	params.StPoint = userLocation.Lat
	params.StPoint_2 = userLocation.Lon

	s.db.CreateLocation(c.Context(), *params)

	log.Printf("lat: %v, lon: %v", userLocation.Lat, userLocation.Lon)

	c.Status(fiber.StatusAccepted)
	return nil
}
