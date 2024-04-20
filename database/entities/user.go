package entities

import "time"

type User struct {
	Name      string `gorm:"not null;size:64"`
	Email     string `gorm:"uniqueIndex;not null;size:255"`
	Username  string `gorm:"primaryKey;not null;unique"`
	Password  string `gorm:"not null;size:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts     []Post
}
