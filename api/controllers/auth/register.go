package controllers

import (
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	// w.Header().Set("Content-Type", "application/json")

	// var user models.User
	// json.NewDecoder(r.Body).Decode(&user)

	// passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	// if err != nil {
	// 	log.Error("failed to generate password hash", err.Error())

	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	enc, _ := json.Marshal(map[string]interface{}{
	// 		"message": "could not signup successfully, try again",
	// 	})
	// 	w.Write(enc)

	// 	return
	// }

	// passwordHashStr := string(passwordHashByte)

	// results, err := repositories.Db.Query(
	// 	"INSERT INTO users(name, email, username, password) VALUES($1, $2, $3, $4);",
	// 	user.Name, user.Email, user.Username, passwordHashStr,
	// )

	// if err != nil {
	// 	log.Error("could not insert user into the database", err.Error())

	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	enc, _ := json.Marshal(map[string]interface{}{
	// 		"message": "could not signup successfully, try again",
	// 	})
	// 	w.Write(enc)

	// 	return
	// }

	// for results.Next() {
	// 	err = results.Scan(&user.Name, &user.Email, &user.Username, &user.Password)
	// 	if err != nil {
	// 		fmt.Println("scan failed; check the number of values in destination and the number of columns")
	// 	}
	// }

	// w.WriteHeader(http.StatusCreated)
	// enc, _ := json.Marshal(map[string]interface{}{
	// 	"message": "successfully signed up the user",
	// })
	// w.Write(enc)

	return nil
}
