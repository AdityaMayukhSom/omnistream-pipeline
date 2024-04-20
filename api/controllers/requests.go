package controllers

type RequestCreatePost struct {
}

type RequestRegisterUser struct {
	Name     string
	Email    string
	Username string
	Password string
}

type RequestUpdateUser struct {
}

type RequestUpdatePost struct {
}
