package models

type User struct {
	Username string
	Password string
	Name     string
	Email    string
}

type LoginCredential struct {
	Username string
	Password string
}
