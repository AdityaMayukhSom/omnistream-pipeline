package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"devstream.in/blogs/repositories"
	"github.com/gorilla/mux"
)

func ReturnUserData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returning user data...")
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := repositories.RetrieveUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		enc, _ := json.Marshal(map[string]any{
			"message": err.Error(),
			"user":    nil,
		})
		w.Write(enc)
		return
	}
	user.Password = "" // do not send user password encrypted value
	w.WriteHeader(http.StatusOK)
	enc, _ := json.Marshal(map[string]any{
		"message": "successfully retrieved user",
		"user":    user,
	})
	w.Write(enc)
}
