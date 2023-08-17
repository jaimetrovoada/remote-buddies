package handlers

import (
	"remote-buddies/server/internal/db"
)

type Service struct {
	db *db.Queries
}

func NewHandler(db *db.Queries) *Service {
	return &Service{db}
}
