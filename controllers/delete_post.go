package controllers

import (
	"net/http"

	"devstream.in/blogs/repositories"
	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	keys := vars["id"]

	_, err := repositories.Db.Query(
		"DELETE FROM posts WHERE id = $1;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}
}
