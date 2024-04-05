package controllers

import (
	"github.com/labstack/echo/v4"
)

func ReturnAllPosts(c echo.Context) error {
	// fmt.Println("returning all posts...")
	// var posts = make([]models.Post, 0)
	// results, err := repositories.Db.Query("SELECT id, title, content FROM posts ORDER BY id DESC;")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// for results.Next() {
	// 	var post models.Post
	// 	err := results.Scan(&post.ID, &post.Title, &post.Content)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	posts = append(posts, post)
	// }

	// nameResults, err := repositories.Db.Query("SELECT username FROM users INNER JOIN posts ON users.id=posts.author ORDER BY posts.id DESC;")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// index := 0
	// for nameResults.Next() {
	// 	if index < len(posts) {
	// 		err := nameResults.Scan(&posts[index].Author)
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}
	// 		index++
	// 	}
	// }

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// json.NewEncoder(w).Encode(posts)
	return nil
}

func ReturnBatchPosts(c echo.Context) error {
	// vars := mux.Vars(r)
	// keys := vars["id"]

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-Type", "application/json")

	// keysInt, _ := strconv.Atoi(keys)
	// var posts = make([]models.Post, 0)

	// q := `
	// SELECT
	// 	posts.id,
	// 	posts.title,
	// 	posts.content,
	// 	users.username
	// FROM
	// 	posts
	// INNER JOIN
	// 	users
	// ON
	// 	posts.author = users.id
	// ORDER BY
	// 	id DESC
	// LIMIT
	// 	5
	// OFFSET
	// 	$1;
	// `
	// results, err := repositories.Db.Query(q, keysInt*5)

	// if err != nil {
	// 	log.Error("error while reading the data", err.Error())
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	enc, _ := json.Marshal(map[string]interface{}{
	// 		"message":    err.Error(),
	// 		"postsCount": 0,
	// 		"posts":      posts,
	// 	})
	// 	w.Write(enc)
	// 	return
	// }

	// for results.Next() {
	// 	var post models.Post
	// 	err := results.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
	// 	if err != nil {
	// 		log.Error("error while parsing the data", err.Error())
	// 	} else {
	// 		posts = append(posts, post)
	// 	}
	// }

	// w.WriteHeader(http.StatusAccepted)
	// enc, _ := json.Marshal(map[string]interface{}{
	// 	"message":    "posts retrieved successfully",
	// 	"postsCount": len(posts),
	// 	"posts":      posts,
	// })
	// w.Write(enc)

	return nil
}

func ReturnSinglePost(c echo.Context) error {
	// fmt.Println("returning one post...")
	// vars := mux.Vars(r)
	// keys := vars["id"]

	// results, err := repositories.Db.Query(
	// 	"SELECT id, title, content FROM posts WHERE id = $1;",
	// 	keys,
	// )
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var post models.Post
	// for results.Next() {
	// 	err = results.Scan(&post.ID, &post.Title, &post.Content)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }

	// nameResults, err := repositories.Db.Query(
	// 	"SELECT username FROM users INNER JOIN posts ON users.id=posts.author WHERE posts.id=$1;",
	// 	keys,
	// )
	// if err != nil {
	// 	panic(err.Error())
	// }
	// for nameResults.Next() {
	// 	err := nameResults.Scan(&post.Author)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// json.NewEncoder(w).Encode(post)

	return nil
}

func ReturnUserPosts(c echo.Context) error {
	// fmt.Println("returning this user's post...")
	// vars := mux.Vars(r)
	// keys := vars["userid"]

	// results, err := repositories.Db.Query(
	// 	"SELECT * FROM posts WHERE author = $1 ORDER BY id DESC;",
	// 	keys,
	// )
	// if err != nil {
	// 	panic(err.Error())
	// }

	// var post models.Post
	// var posts []models.Post
	// for results.Next() {
	// 	err = results.Scan(
	// 		&post.ID,
	// 		&post.Title,
	// 		&post.Author,
	// 		&post.Content,
	// 	)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	posts = append(posts, post)
	// }

	// nameResults, err := repositories.Db.Query(
	// 	"SELECT username FROM users INNER JOIN posts ON users.id = posts.author WHERE author = $1;",
	// 	keys,
	// )
	// if err != nil {
	// 	panic(err.Error())
	// }
	// index := 0
	// for nameResults.Next() {
	// 	if index < len(posts) {
	// 		err := nameResults.Scan(&posts[index].Author)
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}
	// 		index++
	// 	}
	// }

	// json.NewEncoder(w).Encode(posts)
	return nil
}
