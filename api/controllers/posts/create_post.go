package controllers

import (
	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {

	// var post models.Post
	// json.NewDecoder(r.Body).Decode(&post)

	// usernameToIdResults, err := repositories.Db.Query(
	// 	"SELECT id FROM users WHERE username = $1;",
	// 	post.Author,
	// )

	// if err != nil {
	// 	panic(err.Error())
	// }
	// for usernameToIdResults.Next() {
	// 	err = usernameToIdResults.Scan(&post.Author)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }

	// results, err := repositories.Db.Query(
	// 	"INSERT INTO posts(title, author, content) VALUES ($1, $2, $3);",
	// 	post.Title, post.Author, post.Content,
	// )
	// if err != nil {
	// 	panic(err.Error())
	// }

	// for results.Next() {
	// 	err = results.Scan(&post.Title, &post.Author, &post.Content)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }

	// json.NewEncoder(w).Encode(post)
	return nil
}
