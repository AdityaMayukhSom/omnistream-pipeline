package controllers

import (
	"fmt"
	"net/http"

	"devstream.in/pixelated-pipeline/api/controllers"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	var req controllers.RequestRegisterUser

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, controllers.ResponseError{
			ErrorMessage: "could not parse request body",
		})
	}

	fmt.Println(req)

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "user successfully registered",
	})
}
