// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID                     pgtype.UUID        `json:"id"`
	Name                   pgtype.Text        `json:"name"`
	Email                  pgtype.Text        `json:"email"`
	EmailVerified          pgtype.Timestamptz `json:"emailVerified"`
	Image                  pgtype.Text        `json:"image"`
	Coords                 interface{}        `json:"coords"`
	Interests              []string           `json:"interests"`
	CreatedAt              pgtype.Timestamptz `json:"created_at"`
	UpdatedAt              pgtype.Timestamptz `json:"updated_at"`
	OauthType              pgtype.Text        `json:"oauth_type"`
	OauthProvider          pgtype.Text        `json:"oauth_provider"`
	OauthProviderAccountId string             `json:"oauth_providerAccountId"`
	OauthRefreshToken      pgtype.Text        `json:"oauth_refresh_token"`
	OauthAccessToken       pgtype.Text        `json:"oauth_access_token"`
	OauthExpiresAt         pgtype.Int4        `json:"oauth_expires_at"`
	OauthTokenType         pgtype.Text        `json:"oauth_token_type"`
	OauthScope             pgtype.Text        `json:"oauth_scope"`
	OauthIDToken           pgtype.Text        `json:"oauth_id_token"`
	OauthSessionState      pgtype.Text        `json:"oauth_session_state"`
	OauthTokenSecret       pgtype.Text        `json:"oauth_token_secret"`
	OauthToken             pgtype.Text        `json:"oauth_token"`
}
