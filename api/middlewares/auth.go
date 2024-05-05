package middlewares

import (
	"net/http"

	apiConstant "devstream.in/pixelated-pipeline/api/constants"
	"devstream.in/pixelated-pipeline/api/dto"
	"devstream.in/pixelated-pipeline/config"
	"devstream.in/pixelated-pipeline/services"
	"github.com/labstack/echo/v4"
)

func WithAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// log.Info("auth middleware hit")

		if c.Request().Method == http.MethodOptions {
			return c.JSON(http.StatusMethodNotAllowed, dto.ResponseError{
				ErrorMessage: "http method options not supported",
			})
		}

		if tokenCookie, err := c.Cookie(apiConstant.CookieNameAccessToken); err == nil {
			tokenService := services.NewTokenService()
			if username, err := tokenService.ValidateToken(tokenCookie.Value, config.GetAccessSecretKey()); err == nil {
				c.SetCookie(tokenCookie)
				c.Set(apiConstant.ContextAttributeKeyName, username)
				return next(c)
			}
		}

		// return c.JSON(http.StatusUnauthorized, dto.ResponseError{
		// 	ErrorMessage: "please login to continue",
		// })
		return c.Redirect(http.StatusSeeOther, apiConstant.DefaultUnauthenticatedRoute)
	}
}

func WithAlreadyAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	// log.Info("with already authenticated hit")
	return func(c echo.Context) error {
		if tokenCookie, err := c.Cookie(apiConstant.CookieNameAccessToken); err == nil {
			tokenService := services.NewTokenService()
			if username, err := tokenService.ValidateToken(tokenCookie.Value, config.GetAccessSecretKey()); err == nil {
				c.SetCookie(tokenCookie)
				c.Set(apiConstant.ContextAttributeKeyName, username)
				return c.Redirect(http.StatusSeeOther, apiConstant.DefaultAuthenticatedRoute)
			}
		}
		return next(c)
	}
}
