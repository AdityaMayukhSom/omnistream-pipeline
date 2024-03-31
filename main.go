package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"

	"devstream.in/blogs/auth"
	"devstream.in/blogs/config"
	"devstream.in/blogs/controllers"
	"devstream.in/blogs/middlewares"
	"devstream.in/blogs/repositories"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("could not load config", "err", err)
	}
	log.Info(conf.DatabaseConf.Source)
	repositories.SetupDatabase()

	r := mux.NewRouter()
	r.Use(middlewares.WithCorsMiddleware)

	r.HandleFunc("/posts", controllers.ReturnAllPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts/batch/{id}", controllers.ReturnBatchPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", controllers.ReturnSinglePost).Methods(http.MethodGet)

	// r.HandleFunc("/{username}", controllers.ReturnUserData).Methods(http.MethodGet)
	// r.HandleFunc("/{userid}/posts", controllers.ReturnUserPosts).Methods(http.MethodGet)

	restrictedR := r.Path("/").Subrouter()
	restrictedR.Use(middlewares.WithAuthentication)

	restrictedR.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods(http.MethodOptions, http.MethodPut)
	restrictedR.HandleFunc("/posts/{id}", controllers.DeletePost).Methods(http.MethodOptions, http.MethodDelete)
	restrictedR.HandleFunc("/post", controllers.CreateNewPost).Methods(http.MethodOptions, http.MethodPost)

	authR := r.PathPrefix("/auth").Subrouter()
	authR.HandleFunc("/signup", auth.SignUp).Methods(http.MethodOptions, http.MethodPost)
	authR.HandleFunc("/login", auth.LogIn).Methods(http.MethodOptions, http.MethodPost)
	authR.HandleFunc("/refresh", auth.Refresh).Methods(http.MethodOptions, http.MethodPost)

	// fs := http.FileServer(http.Dir("./frontend/dist"))
	// r.PathPrefix("/").Handler(fs)

	err = http.ListenAndServe(":"+fmt.Sprint(conf.Port), r)
	if err != nil {
		log.Fatal("could not start http server", "err", err)
	}

	repositories.CleanupDatabase()
}
