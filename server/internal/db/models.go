// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID                string      `json:"id"`
	UserId            string      `json:"userId"`
	Type              string      `json:"type"`
	Provider          string      `json:"provider"`
	ProviderAccountId string      `json:"providerAccountId"`
	RefreshToken      pgtype.Text `json:"refresh_token"`
	AccessToken       pgtype.Text `json:"access_token"`
	ExpiresAt         pgtype.Int4 `json:"expires_at"`
	TokenType         pgtype.Text `json:"token_type"`
	Scope             pgtype.Text `json:"scope"`
	IDToken           pgtype.Text `json:"id_token"`
	SessionState      pgtype.Text `json:"session_state"`
}

type Location struct {
	ID     pgtype.UUID `json:"id"`
	Coords interface{} `json:"coords"`
}

type Session struct {
	ID           string           `json:"id"`
	SessionToken string           `json:"sessionToken"`
	UserId       string           `json:"userId"`
	Expires      pgtype.Timestamp `json:"expires"`
}

type User struct {
	ID            string           `json:"id"`
	Name          pgtype.Text      `json:"name"`
	Email         pgtype.Text      `json:"email"`
	EmailVerified pgtype.Timestamp `json:"emailVerified"`
	Image         pgtype.Text      `json:"image"`
}

type VerificationToken struct {
	Identifier string           `json:"identifier"`
	Token      string           `json:"token"`
	Expires    pgtype.Timestamp `json:"expires"`
}
