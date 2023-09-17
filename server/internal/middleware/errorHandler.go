package middleware

import (
	"errors"
	"net/http"
	"net/url"
	cErrors "remote-buddies/server/internal/errors"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	url, _ := url.Parse("http://localhost:3000/")
	qParams := url.Query()
	var authError *cErrors.AuthError
	if errors.As(err, &authError) {
		c.Logger().Error("cause: ", authError.Err)
		qParams.Add("error", authError.Message)
		url.RawQuery = qParams.Encode()
		c.Redirect(http.StatusFound, url.String())
		return
	}

	c.Logger().Error(err)
	c.JSON(code, err)
}
