package controllers

import (
	"net/http"
	"time"

	apiConstant "devstream.in/pixelated-pipeline/api/constants"
	"devstream.in/pixelated-pipeline/api/dto"
	"devstream.in/pixelated-pipeline/services"
	"devstream.in/pixelated-pipeline/services/models"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) (err error) {
	var req dto.RequestRegisterUser

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseError{
			Type: apiConstant.ERROR_TYPE_INVALID_BODY,
		})
	}

	userService := services.NewUserService()

	err = userService.RegisterUser(models.User{
		Username: req.Username,
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.ResponseRegisteredUser{
		Type: apiConstant.RESPONSE_TYPE_SUCCESSFUL_REGISTERED,
	})
}

func LogIn(c echo.Context) error {
	var req dto.RequestLoginUser

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseError{
			Type: apiConstant.ERROR_TYPE_INVALID_BODY,
		})
	}

	userService := services.NewUserService()

	// TODO: LoginCredential shouldn't be created manually, rather some mapper
	// object shall do the mapping between request to login credentials automatically
	tokenStruct, err := userService.LoginUser(models.LoginCredential{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ResponseError{
			Type:    apiConstant.ERROR_TYPE_INVALID_LOGIN_CREDENTIALS,
			Message: err.Error(),
		})
	}

	user, err := userService.GetDetails(req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Message: err.Error(),
		})
	}

	accessCookie := &http.Cookie{
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Name:     apiConstant.CookieNameAccessToken,
		Value:    tokenStruct.AccessToken,
		Path:     apiConstant.CookiePathAccessToken,
		Expires:  time.Now().Add(time.Minute * 15),
		// Secure:   false, // when true browser only transmit the cookie over https channel
	}

	refreshCookie := &http.Cookie{
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Name:     apiConstant.CookieNameRefreshToken,
		Value:    tokenStruct.RefreshToken,
		Path:     apiConstant.CookiePathRefreshToken, // matches api version 1, TODO : to match any api auth version
		Expires:  time.Now().Add(time.Hour * 24),
		// Secure:   false, // when true browser only transmit the cookie over https channel
	}

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusAccepted, dto.ResponseLoginUser{
		Type: apiConstant.RESPOSNE_TYPE_NEWLY_AUTHENTICATED,
		User: dto.UserDTO{
			Name:     user.Name,
			Username: user.Username,
		},
	})
}

func LogOut(c echo.Context) error {
	accessCookie := &http.Cookie{
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Name:     apiConstant.CookieNameAccessToken,
		Value:    "",
		MaxAge:   -1,
		Path:     apiConstant.CookiePathAccessToken,
		Expires:  time.Unix(0, 0),
	}

	refreshCookie := &http.Cookie{
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Name:     apiConstant.CookieNameRefreshToken,
		Value:    "",
		MaxAge:   -1,
		Path:     apiConstant.CookiePathRefreshToken,
		Expires:  time.Unix(0, 0),
	}

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, dto.ResponseLogoutUser{
		Type:       apiConstant.RESPONSE_TYPE_SUCCESSFUL_LOGOUT,
		LogoutTime: time.Now(),
	})
}

func Refresh(c echo.Context) error {
	// get refreshToken from cookie
	// var refreshToken, refreshUsername string

	// c, err := r.Cookie("refreshToken")

	// if err != nil {
	// 	w.Write([]byte("error in reading cookie : " + err.Error() + "\n"))
	// } else {
	// 	refreshToken = c.Value
	// }

	// // verify refreshToken and check its validity
	// refreshSecretKey := config.DefaultConfig.RefreshSecretKey

	// token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	return []byte(refreshSecretKey), nil
	// })

	// if err != nil {
	// 	fmt.Println("Failed to verify token.")
	// 	return
	// }

	// if !token.Valid {
	// 	fmt.Println("Invalid token.")
	// 	return
	// }

	// claims, ok := token.Claims.(jwt.MapClaims)

	// if !ok {
	// 	fmt.Println(err)
	// }
	// refreshUsername = claims["user_username"].(string)

	// // query the user and use that data as parameters for generating new tokens
	// retrievedUser, err := repositories.RetrieveUserByUsername(refreshUsername)

	// if err != nil {

	// }

	// // generate new access and refresh tokens
	// tokenStruct, err := generateToken(
	// 	retrievedUser.Name,
	// 	retrievedUser.Email,
	// 	retrievedUser.Username,
	// )

	// if err != nil {
	// 	fmt.Println("failed to generate token")
	// }

	// accessToken := tokenStruct.AccessToken
	// refreshToken = tokenStruct.RefreshToken
	// cookie := http.Cookie{
	// 	HttpOnly: true,
	// 	Name:     "refreshToken",
	// 	Value:    refreshToken,
	// 	// Domain:   "jpoly1219devbox.xyz",
	// 	Path: "/auth/",
	// }

	// responseJSON := map[string]string{
	// 	"accessToken": accessToken,
	// }

	// http.SetCookie(w, &cookie)
	// json.NewEncoder(w).Encode(responseJSON)
	return nil
}
