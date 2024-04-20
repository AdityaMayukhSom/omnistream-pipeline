package repositories

import "gorm.io/gorm"

type CommentEntity struct {
	gorm.Model
	Text      string `gorm:"not null"`
	UserId    string
	PostId    string
	Commenter UserEntity `gorm:"foreignKey:user_id;references:username"`
	Post      PostEntity `gorm:"foreignKey:post_id;references:id"`
}

func (CommentEntity) TableName() string {
	return "comments"
}
