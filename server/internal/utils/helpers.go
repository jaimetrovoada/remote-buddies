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

func GetUser(email string, query *db.Queries) (db.User, error) {

	uEmail := pgtype.Text{}
	uEmail.Scan(email)
	return query.GetUser(context.Background(), uEmail)
}

func CreateNewUser(user goth.User, query *db.Queries) (db.User, error) {

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

	return query.CreateUser(context.Background(), *dbUser)

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
