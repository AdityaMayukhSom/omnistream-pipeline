package middlewares

import (
	"net/http"

	apiConstant "devstream.in/pixelated-pipeline/api/constants"
	"devstream.in/pixelated-pipeline/api/dto"
	"devstream.in/pixelated-pipeline/config"
	"devstream.in/pixelated-pipeline/services"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

func WithAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("auth middleware hit")

		if c.Request().Method == http.MethodOptions {
			return c.JSON(http.StatusMethodNotAllowed, dto.ResponseError{
				ErrorMessage: "http method options not supported",
			})
		}

		var tokenStr string
		var accessCookieFound bool = false

		cookies := c.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == apiConstant.CookieNameAccessToken {
				accessCookieFound = true
				tokenStr = cookie.Value
				break
			}
		}

		if accessCookieFound {
			tokenService := services.NewTokenService()
			if username, err := tokenService.ValidateToken(tokenStr, config.GetAccessSecretKey()); err == nil {
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
	return func(c echo.Context) error {
		var tokenStr string
		var accessCookieFound bool = false

		cookies := c.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == apiConstant.CookieNameAccessToken {
				accessCookieFound = true
				tokenStr = cookie.Value
				break
			}
		}

		if accessCookieFound {
			tokenService := services.NewTokenService()
			if username, err := tokenService.ValidateToken(tokenStr, config.GetAccessSecretKey()); err == nil {
				c.Set(apiConstant.ContextAttributeKeyName, username)
				return c.Redirect(http.StatusSeeOther, apiConstant.DefaultAuthenticatedRoute)
			}
		}

		return next(c)
	}
}
