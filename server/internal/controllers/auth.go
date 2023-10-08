package controllers

import (
	"remote-buddies/server/internal/config"
	"remote-buddies/server/internal/db"
	"remote-buddies/server/internal/errors"
	"remote-buddies/server/internal/services"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func (s *Service) AuthHandler(c echo.Context) error {
	// try to get the user without re-authenticating
	w, r := c.Response(), c.Request()
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		return &errors.AuthError{Message: "AuthError", Err: err}
	} else {
		gothic.BeginAuthHandler(w, r)
	}
	return nil
}

func (s *Service) AuthCallbackHandler(c echo.Context) error {

	config, err := config.LoadConfig()
	if err != nil {
		return err
	}
	url := config.FRONTEND_URL

	w, r := c.Response(), c.Request()
	gUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return &errors.AuthError{Message: "AuthError", Err: err}
	}

	var user db.User

	user, err = services.GetUser(gUser.Email, s.db)
	if err != nil {
		if err == pgx.ErrNoRows {
			user, err = services.CreateNewUser(gUser, s.db)
			if err != nil {
				return &errors.AuthError{Message: "AuthError", Err: err}
			}
		} else {

			return &errors.AuthError{Message: "AuthError", Err: err}
		}

	}

	return services.LoginUser(user, url, c)

}
