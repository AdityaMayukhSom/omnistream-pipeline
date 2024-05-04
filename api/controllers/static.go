package controllers

import "github.com/labstack/echo/v4"

func DisplayLoginPage(c echo.Context) error {
	return c.File("views/login.html")
}

func DisplaySignupPage(c echo.Context) error {
	return c.File("views/signup.html")
}
