package controllers

import (
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	// json.NewDecoder(r.Body).Decode(&user)

	// passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	// passwordHashStr := string(passwordHashByte)

	// w.WriteHeader(http.StatusCreated)
	// enc, _ := json.Marshal(map[string]interface{}{
	// 	"message": "successfully signed up the user",
	// })
	// w.Write(enc)

	return nil
}
