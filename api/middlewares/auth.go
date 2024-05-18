package middlewares

import (
	"net/http"
	"time"

	apiConstant "devstream.in/pixelated-pipeline/api/constants"
	"devstream.in/pixelated-pipeline/api/dto"
	"devstream.in/pixelated-pipeline/config"
	"devstream.in/pixelated-pipeline/services"
	"github.com/go-http-utils/headers"
	"github.com/labstack/echo/v4"
)

func WithAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add(headers.CacheControl, "private, no-store, max-age=0, no-cache, must-revalidate, post-check=0, pre-check=0")
		c.Response().Header().Add(headers.Expires, time.Unix(0, 0).Format(http.TimeFormat))
		c.Response().Header().Add(headers.Pragma, "no-cache")

		if tokenCookie, err := c.Cookie(apiConstant.CookieNameAccessToken); err == nil {
			tokenService := services.NewTokenService()
			if username, name, err := tokenService.ValidateToken(tokenCookie.Value, config.GetAccessSecretKey()); err == nil {
				c.Set(apiConstant.ContextAttributeKeyName, name)
				c.Set(apiConstant.ContextAttributeKeyUsername, username)
				return next(c)
			}
		}

		return c.Redirect(http.StatusSeeOther, apiConstant.DefaultUnauthenticatedRoute)
	}
}

func WithAlreadyAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add(headers.CacheControl, "private, no-store, max-age=0, no-cache, must-revalidate, post-check=0, pre-check=0")
		c.Response().Header().Add(headers.Expires, time.Unix(0, 0).Format(http.TimeFormat))
		c.Response().Header().Add(headers.Pragma, "no-cache")

		if tokenCookie, err := c.Cookie(apiConstant.CookieNameAccessToken); err == nil && tokenCookie.Valid() == nil {
			tokenService := services.NewTokenService()
			if username, name, err := tokenService.ValidateToken(tokenCookie.Value, config.GetAccessSecretKey()); err == nil {
				return c.JSON(http.StatusOK, dto.ResponseLoginUser{
					Type: apiConstant.RESPONSE_TYPE_ALREADY_AUTHENTICATED,
					User: dto.UserDTO{
						Name:     name,
						Username: username,
					},
				})
			}
		}

		return next(c)
	}
}
