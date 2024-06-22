package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// ListAccounts lists all existing accounts
//
//	@summary		This returns basic hello world.
//	@description	This is the first line
//	@router			/hello-text [get]
func HelloTextHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Omnistream.")
}

func HelloFileHandler(c echo.Context) error {
	return c.File("public/home.txt")
}
func HelloJsonHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"Message": "Hello World",
		"Time":    time.Now(),
	})
}
