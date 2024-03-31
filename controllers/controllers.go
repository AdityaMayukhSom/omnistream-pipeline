package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"devstream.in/blogs/models"
	"devstream.in/blogs/repositories"

	"github.com/gorilla/mux"
)

func ReturnUserPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returning this user's post...")
	vars := mux.Vars(r)
	keys := vars["userid"]

	results, err := repositories.Db.Query(
		"SELECT * FROM posts WHERE author = $1 ORDER BY id DESC;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}

	var post models.Post
	var posts []models.Post
	for results.Next() {
		err = results.Scan(
			&post.ID,
			&post.Title,
			&post.Author,
			&post.Content,
		)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	nameResults, err := repositories.Db.Query(
		"SELECT username FROM users INNER JOIN posts ON users.id = posts.author WHERE author = $1;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}
	index := 0
	for nameResults.Next() {
		if index < len(posts) {
			err := nameResults.Scan(&posts[index].Author)
			if err != nil {
				panic(err.Error())
			}
			index++
		}
	}

	json.NewEncoder(w).Encode(posts)
}
