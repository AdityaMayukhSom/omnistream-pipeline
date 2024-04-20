package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Content     string `gorm:"not null;type:text"`
	IsPublished bool   `gorm:"not null"`
	UserID      string `gorm:"index"`
	User        User   `gorm:"foreignKey:user_id;references:username"`
}
