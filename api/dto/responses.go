package dto

import "time"

type ResponseLoginUser struct {
	Type string  `json:"type"`
	User UserDTO `json:"user"`
}

type ResponseLogoutUser struct {
	Type       string    `json:"type"`
	LogoutTime time.Time `json:"logoutTime"`
}

type ResponseRegisteredUser struct {
	Type string `json:"type"`
}

type ResponseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ResponseCreatedPost struct {
}

type ResponseGetMultiplePosts struct {
}

type ResponseGetMultipleUsers struct {
}

type ResponseGetSinglePost struct {
}

type ResponseGetSingleUser struct {
}

type ResponseUpdatedPost struct {
}

type ResponseUpdatedUser struct {
}
