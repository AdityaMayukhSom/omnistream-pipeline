package controllers

import (
	"github.com/labstack/echo/v4"
)

func ReturnUserData(c echo.Context) error {
	// fmt.Println("returning user data...")
	// vars := mux.Vars(r)
	// username := vars["username"]

	// user, err := repositories.RetrieveUserByUsername(username)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	enc, _ := json.Marshal(map[string]any{
	// 		"message": err.Error(),
	// 		"user":    nil,
	// 	})
	// 	w.Write(enc)
	// 	return
	// }
	// user.Password = "" // do not send user password encrypted value
	// w.WriteHeader(http.StatusOK)
	// enc, _ := json.Marshal(map[string]any{
	// 	"message": "successfully retrieved user",
	// 	"user":    user,
	// })
	// w.Write(enc)

	return nil
}
