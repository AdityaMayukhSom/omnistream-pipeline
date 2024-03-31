package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`

	Posts []Post
}
