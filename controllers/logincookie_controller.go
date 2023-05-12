package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)



func GetCookieHandler(c echo.Context) error {
	cookie, err := c.Cookie("JWTCookie")
	if err != nil {
		if err == http.ErrNoCookie {
			// handle jika cookie tidak ditemukan
			return c.String(http.StatusUnauthorized, err.Error())
		}
		// handle error lainnya
		return err
	}

	// handle jika cookie ditemukan
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "Cookie ditemukan")
}
