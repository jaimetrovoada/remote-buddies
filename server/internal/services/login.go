package services

import (
	"net/http"
	"remote-buddies/server/internal/config"
	"remote-buddies/server/internal/db"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
	jwt.RegisteredClaims
}

func LoginUser(user db.User, url string, ctx echo.Context) error {
	token, err := genJWTToken(user)
	if err != nil {
		return err
	}

	cookie := createCookies(token)
	ctx.SetCookie(cookie)

	return ctx.Redirect(http.StatusFound, url)
}

func createCookies(token string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "oauth.session-token",
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	return cookie
}

func genJWTToken(user db.User) (string, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	claims := &JwtCustomClaims{
		user.Name.String,
		user.Email.String,
		user.Image.String,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRET))

}
