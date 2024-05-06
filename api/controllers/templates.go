package controllers

import (
	"net/http"

	apiConstant "devstream.in/pixelated-pipeline/api/constants"
	"github.com/labstack/echo/v4"
)

// Inspired by gin.H
type H map[string]any

func RenderHelloWorldPage(c echo.Context) error {
	return c.Render(http.StatusOK, "helloworld", H{
		"name":     c.Get(apiConstant.ContextAttributeKeyName),
		"username": c.Get(apiConstant.ContextAttributeKeyUsername),
		"message":  "Hello, Sweetheart <3",
	})
}

func RenderHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home", H{
		"name":     c.Get(apiConstant.ContextAttributeKeyName),
		"username": c.Get(apiConstant.ContextAttributeKeyUsername),
	})
}

func RenderWritePage(c echo.Context) error {
	return c.Render(http.StatusOK, "write", H{
		"name":     c.Get(apiConstant.ContextAttributeKeyName),
		"username": c.Get(apiConstant.ContextAttributeKeyUsername),
	})
}
