package dto

type RequestCreatePost struct {
}

type RequestRegisterUser struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RequestLoginUser struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RequestUpdateUser struct {
}

type RequestUpdatePost struct {
}
