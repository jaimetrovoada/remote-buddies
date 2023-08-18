package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"remote-buddies/server/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func (s *Service) AuthHandler(c echo.Context) error {
	// try to get the user without re-authenticating
	w, r := c.Response(), c.Request()
	url, _ := url.Parse("http://localhost:3000/")
	qParams := url.Query()
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	} else {
		gothic.BeginAuthHandler(w, r)
	}
	return nil
}

func (s *Service) AuthCallbackHandler(c echo.Context) error {

	url, _ := url.Parse("http://localhost:3000/")
	qParams := url.Query()

	w, r := c.Response(), c.Request()
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	}

	exists, err := utils.CheckUserExists(user.Email, s.db)
	if err != nil {
		log.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	}
	if exists {
		qParams.Add("error", "UserExists")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	}

	id, err := utils.CreateNewUser(user, s.db)
	if err != nil {
		log.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	}

	if err := utils.CreateNewAccount(id, user, s.db); err != nil {
		log.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	}

	return c.Redirect(http.StatusFound, url.String())

}
