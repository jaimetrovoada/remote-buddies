package handlers

import (
	"net/http"
	"net/url"
	"remote-buddies/server/internal/utils"

	"github.com/gofiber/fiber/v2/log"
	"github.com/markbates/goth/gothic"
)

func (s *Service) AuthHandler(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	url, _ := url.Parse("http://localhost:3000/")
	qParams := url.Query()
	if _, err := gothic.CompleteUserAuth(res, req); err == nil {
		log.Error(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		http.Redirect(res, req, url.String(), http.StatusFound)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func (s *Service) AuthCallbackHandler(res http.ResponseWriter, req *http.Request) {

	url, _ := url.Parse("http://localhost:3000/")
	qParams := url.Query()

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		log.Error(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		http.Redirect(res, req, url.String(), http.StatusFound)
		return
	}

	exists, err := utils.CheckUserExists(user.Email, s.db)
	if err != nil {
		log.Error(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		http.Redirect(res, req, url.String(), http.StatusFound)
		return
	}
	if exists {
		qParams.Add("error", "UserExists")
		url.RawQuery = qParams.Encode()
		http.Redirect(res, req, url.String(), http.StatusFound)
		return
	}

	id, err := utils.CreateNewUser(user, s.db)
	if err != nil {
		log.Error(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		http.Redirect(res, req, url.String(), http.StatusFound)
		return
	}

	if err := utils.CreateNewAccount(id, user, s.db); err != nil {
		log.Error(err)
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		http.Redirect(res, req, url.String(), http.StatusFound)
		return
	}

	http.Redirect(res, req, url.String(), http.StatusFound)
}
