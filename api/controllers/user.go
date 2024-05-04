package controllers

import "github.com/labstack/echo/v4"

func EditUserData(c echo.Context) error {

	// vars := mux.Vars(r)
	// keys := vars["username"]

	// var updatedUser models.User
	// json.NewDecoder(r.Body).Decode(&updatedUser)

	// results, err := repositories.Db.Query(
	// 	"UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4;",
	// 	updatedUser.Name, updatedUser.Email, updatedUser.Password, keys,
	// )
	// if err != nil {
	// 	panic(err.Error())
	// }

	// for results.Next() {
	// 	err = results.Scan(
	// 		&updatedUser.Name,
	// 		&updatedUser.Email,
	// 		&updatedUser.Username,
	// 		&updatedUser.Password,
	// 	)

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }

	// json.NewEncoder(w).Encode(updatedUser)

	return nil
}

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
