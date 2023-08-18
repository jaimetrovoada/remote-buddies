package utils

import (
	"context"
	"fmt"
	"log"
	"remote-buddies/server/internal/config"
	"remote-buddies/server/internal/db"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth"
)

func CheckUserExists(email string, query *db.Queries) (bool, error) {

	uEmail := pgtype.Text{}
	uEmail.Scan(email)
	exists, err := query.CheckUserExists(context.Background(), uEmail)
	if err != nil {
		return exists == 1, err
	}
	return exists == 1, nil
}

func CreateNewUser(user goth.User, query *db.Queries) (string, error) {

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

	result, err := query.CreateUser(context.Background(), *dbUser)
	if err != nil {
		return "", err
	}

	id := fmt.Sprintf("%x", result.ID.Bytes)

	return id, nil
}

func CreateNewAccount(id string, user goth.User, query *db.Queries) error {

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
	account.UserId = fmt.Sprintf("%x", id)
	account.Type = "oauth"
	account.Provider = user.Provider
	account.ProviderAccountId = user.UserID
	account.Scope = scope
	account.AccessToken = accessToken
	account.RefreshToken = refreshToken
	account.TokenType = tokenType

	query.CreateAccount(context.Background(), *account)
	return nil

}

// func CreateCookies(res http.ResponseWriter, req *http.Request, id string) {
// 	cookie := &http.Cookie{
// 		Name:     "oauth.session-token",
// 		Value:    id,
// 		Path:     "/",
// 		Domain:   "localhost",
// 		Expires:  time.Now().Add(24 * time.Hour),
// 		HttpOnly: true,
// 	}
//
// 	http.SetCookie(res, cookie)
// }
