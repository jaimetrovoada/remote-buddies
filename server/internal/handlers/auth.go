package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"remote-buddies/server/internal/config"
	"remote-buddies/server/internal/db"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth/gothic"
)

func (s *Service) AuthHandler(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		fmt.Printf("gothUser:\n%+v\n", gothUser)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func (s *Service) AuthCallbackHandler(res http.ResponseWriter, req *http.Request) {

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, jsonData, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user:\n%+v\n", buf.String())

	name := pgtype.Text{}
	email := pgtype.Text{}
	image := pgtype.Text{}
	updatedAt := pgtype.Timestamptz{}

	name.Scan(user.Name)
	email.Scan(user.Email)
	image.Scan(user.AvatarURL)
	updatedAt.Scan(time.Now())

	dbUser := new(db.CreateUserParams)
	dbUser.Email = email
	dbUser.Name = name
	dbUser.Image = image
	dbUser.UpdatedAt = updatedAt

	fmt.Println("dbuser", dbUser)

	result, err := s.db.CreateUser(context.Background(), *dbUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	scope := pgtype.Text{}
	accessToken := pgtype.Text{}
	refreshToken := pgtype.Text{}
	tokenType := pgtype.Text{}

	scope.Scan(config.GITHUB_CLIENT_SCOPE)
	accessToken.Scan(user.AccessToken)
	refreshToken.Scan(user.RefreshToken)
	tokenType.Scan("bearer")

	account := new(db.CreateAccountParams)
	account.UserId = fmt.Sprintf("%x", result.ID.Bytes)
	account.Type = "oauth"
	account.Provider = user.Provider
	account.ProviderAccountId = user.UserID
	account.Scope = scope
	account.AccessToken = accessToken
	account.RefreshToken = refreshToken
	account.TokenType = tokenType

	s.db.CreateAccount(context.Background(), *account)

	cookie := &http.Cookie{
		Name:     "oauth.session-token",
		Value:    "cookie_value",
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	// Set the cookie in the HTTP response
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "http://localhost:3000", http.StatusFound)
}
