package services

import (
	"context"
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
	config, err := config.LoadConfig()
	if err != nil {
		return db.User{}, err
	}

	name := pgtype.Text{}
	email := pgtype.Text{}
	image := pgtype.Text{}
	updatedAt := pgtype.Timestamptz{}
	oauthType := pgtype.Text{}
	oauthProvider := pgtype.Text{}
	oauthScope := pgtype.Text{}
	oauthAccessToken := pgtype.Text{}
	oauthRefreshToken := pgtype.Text{}
	oauthTokenType := pgtype.Text{}

	oauthType.Scan("oauth")
	oauthProvider.Scan(user.Provider)
	oauthScope.Scan(config.GITHUB_CLIENT_SCOPE)
	oauthAccessToken.Scan(user.AccessToken)
	oauthRefreshToken.Scan(user.RefreshToken)
	oauthTokenType.Scan("bearer")

	name.Scan(user.Name)
	email.Scan(user.Email)
	image.Scan(user.AvatarURL)
	updatedAt.Scan(time.Now())

	dbUser := new(db.CreateUserParams)
	dbUser.Email = email
	dbUser.Name = name
	dbUser.Image = image
	dbUser.UpdatedAt = updatedAt
	dbUser.OauthType = oauthType
	dbUser.OauthProvider = oauthProvider
	dbUser.OauthProviderAccountId = user.UserID
	dbUser.OauthScope = oauthScope
	dbUser.OauthAccessToken = oauthAccessToken
	dbUser.OauthRefreshToken = oauthRefreshToken
	dbUser.OauthTokenType = oauthTokenType

	return query.CreateUser(context.Background(), *dbUser)

}
