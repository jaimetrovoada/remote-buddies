package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"remote-buddies/server/internal/db"
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
	gUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	}

	var user db.User

	if qUser, err := utils.GetUser(gUser.Email, s.db); err != nil {
		log.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	} else {
		if user.Email.String != "" {
			user = qUser
			return utils.LoginUser(user, url.String(), c)
		}
	}

	if qUser, err := utils.CreateNewUser(gUser, s.db); err != nil {
		log.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	} else {
		if user.Email.String != "" {
			user = qUser
		}
	}

	id := fmt.Sprintf("%d", user.ID.Bytes)
	if err := utils.CreateNewAccount(id, gUser, s.db); err != nil {
		log.Println(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		return c.Redirect(http.StatusFound, url.String())
	}

	return utils.LoginUser(user, url.String(), c)

}
