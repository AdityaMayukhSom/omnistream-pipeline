package controllers

import (
	"net/http"

	"devstream.in/pixelated-pipeline/api/controllers"
	service "devstream.in/pixelated-pipeline/services"
	"devstream.in/pixelated-pipeline/services/models"
	"github.com/labstack/echo/v4"
)

func LogIn(c echo.Context) error {
	var req controllers.RequestLoginUser

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, controllers.ResponseError{
			ErrorMessage: "could not parse request body",
		})
	}

	userService := service.NewUserService()

	// TODO: LoginCredential shouldn't be created manually, rather some mapper
	// object shall do the mapping between request to login credentials automatically
	tokenStruct, err := userService.LoginUser(models.LoginCredential{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, controllers.ResponseError{
			ErrorMessage: err.Error(),
		})
	}

	accessCookie := http.Cookie{
		HttpOnly: true,
		Name:     "accessToken",
		Value:    tokenStruct.AccessToken,
	}

	refreshCookie := http.Cookie{
		HttpOnly: true,
		Name:     "refreshToken",
		Value:    tokenStruct.RefreshToken,
		Path:     "/auth/",
	}

	c.SetCookie(&accessCookie)
	c.SetCookie(&refreshCookie)
	c.Redirect(http.StatusFound, "/homepage.html")

	return nil
}
