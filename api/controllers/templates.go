package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RenderHelloWorldPage(c echo.Context) error {
	return c.Render(http.StatusOK, "helloworld.go.html", map[string]interface{}{
		"message": "Hello from Rendered Views",
	})
}

func RenderHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "homepage.go.html", map[string]interface{}{})
}
