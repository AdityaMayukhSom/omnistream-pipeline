package repositories

import (
	"devstream.in/pixelated-pipeline/services/models"
	"gorm.io/gorm"
)

type PostEntity struct {
	gorm.Model
	Title       string     `gorm:"not null;type:varchar(255);uniqueIndex:authorAndTitleUnique"`
	Summary     string     `gorm:"type:text"`
	Content     string     `gorm:"not null;type:text"`
	IsPublished bool       `gorm:"not null;type:bool"`
	UserId      string     `gorm:"not null;uniqueIndex:authorAndTitleUnique"`
	User        UserEntity `gorm:"foreignKey:user_id;references:username"`
}

func (PostEntity) TableName() string {
	return "posts"
}

type PostRepository interface {
	FindPostById(id string)
	FindPostsByUsername(username string)

	DeletePostById(id string)

	CreatePost(post models.Post)
	UpdatePost(post models.Post)
}
