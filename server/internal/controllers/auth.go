package controllers

import (
	"fmt"
	"remote-buddies/server/internal/db"
	"remote-buddies/server/internal/errors"
	"remote-buddies/server/internal/services"

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

	url := "http://localhost:3000/"

	w, r := c.Response(), c.Request()
	gUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return &errors.AuthError{Message: "AuthError", Err: err}
	}

	var user db.User

	if qUser, err := services.GetUser(gUser.Email, s.db); err != nil {
		return &errors.AuthError{Message: "AuthError", Err: err}
	} else {
		if user.Email.String != "" {
			user = qUser
			return services.LoginUser(user, url, c)
		}
	}

	if qUser, err := services.CreateNewUser(gUser, s.db); err != nil {
		return &errors.AuthError{Message: "AuthError", Err: err}
	} else {
		if user.Email.String != "" {
			user = qUser
		}
	}

	id := fmt.Sprintf("%d", user.ID.Bytes)
	if err := services.CreateNewAccount(id, gUser, s.db); err != nil {
		return &errors.AuthError{Message: "AuthError", Err: err}
	}

	return services.LoginUser(user, url, c)

}
