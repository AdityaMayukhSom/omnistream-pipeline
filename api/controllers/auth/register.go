package controllers

import (
	"fmt"
	"net/http"

	"devstream.in/pixelated-pipeline/api/controllers"
	service "devstream.in/pixelated-pipeline/services"
	"devstream.in/pixelated-pipeline/services/models"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) (err error) {
	var req controllers.RequestRegisterUser

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, controllers.ResponseError{
			ErrorMessage: "could not parse request body",
		})
	}

	fmt.Println(req)

	userService := service.NewUserService()

	err = userService.RegisterUser(models.User{
		Username: req.Username,
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, controllers.ResponseError{
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "user successfully registered",
	})
}
