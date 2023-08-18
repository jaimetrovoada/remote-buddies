package middleware

import (
	"net/http"
	"net/url"
	"remote-buddies/server/internal/utils"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	url, _ := url.Parse("http://localhost:3000/")
	qParams := url.Query()
	if _, ok := err.(*utils.AuthError); ok {
		// code = serr.Code
		qParams.Add("error", "AuthError")
		url.RawQuery = qParams.Encode()
		c.Redirect(http.StatusFound, url.String())
	}
	c.Logger().Error(err)
	c.JSON(code, err)
}
