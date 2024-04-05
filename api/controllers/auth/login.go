package controllers

import (
	"github.com/labstack/echo/v4"
)

func LogIn(c echo.Context) error {
	// var receivedUser models.User
	// json.NewDecoder(r.Body).Decode(&receivedUser)

	// existingUser, err := repositories.RetrieveUserByEmail(receivedUser.Email)

	// if err != nil {

	// 	log.Error("could not retrieve user details for given email", "email", receivedUser.Email)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	enc, _ := json.Marshal(map[string]interface{}{
	// 		"message": err.Error(),
	// 	})
	// 	w.Write(enc)
	// 	return
	// }

	// pwMatchErr := bcrypt.CompareHashAndPassword(
	// 	[]byte(existingUser.Password),
	// 	[]byte(receivedUser.Password),
	// )

	// if pwMatchErr != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	enc, _ := json.Marshal(map[string]interface{}{
	// 		"message": "password does not match",
	// 	})
	// 	w.Write(enc)
	// }

	// name := existingUser.Name
	// email := existingUser.Email
	// username := existingUser.Username

	// tokenStruct, err := generateToken(name, email, username)

	// if err != nil {
	// 	log.Error("could not generate acces token for given email", "email", receivedUser.Email)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	enc, _ := json.Marshal(map[string]interface{}{
	// 		"message": "could not login successfully, try again",
	// 	})
	// 	w.Write(enc)
	// 	return
	// }

	// accessToken := tokenStruct.AccessToken
	// refreshToken := tokenStruct.RefreshToken
	// cookie := http.Cookie{
	// 	HttpOnly: true,
	// 	Name:     "refreshToken",
	// 	Value:    refreshToken,
	// 	// Domain:   "jpoly1219devbox.xyz",
	// 	Path: "/auth/",
	// }

	// http.SetCookie(w, &cookie)
	// w.WriteHeader(http.StatusOK)
	// enc, _ := json.Marshal(map[string]interface{}{
	// 	"message":     "user successfully logged in",
	// 	"accessToken": accessToken,
	// })
	// w.Write(enc)
	return nil
}
