package controllers

import (
	"github.com/labstack/echo/v4"
)

func UpdatePost(c echo.Context) error {

	// vars := mux.Vars(r)
	// keys := vars["id"]

	// var updatedPost models.Post
	// json.NewDecoder(r.Body).Decode(&updatedPost)

	// results, err := repositories.Db.Query(
	// 	"UPDATE posts SET title=$1, content=$2 WHERE id=$3;",
	// 	updatedPost.Title, updatedPost.Content, keys,
	// )
	// if err != nil {
	// 	panic(err.Error())
	// }

	// var post models.Post
	// for results.Next() {
	// 	err = results.Scan(&post.ID, &post.Title, &post.Author, &post.Content)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }

	// json.NewEncoder(w).Encode(post)
	return nil
}
