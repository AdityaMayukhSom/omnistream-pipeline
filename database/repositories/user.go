package repositories

import (
	"time"

	"devstream.in/pixelated-pipeline/services/models"
)

type UserEntity struct {
	Name      string       `gorm:"not null;size:64"`
	Email     string       `gorm:"uniqueIndex;not null;size:255"`
	Username  string       `gorm:"primaryKey;not null;unique;size:64"`
	Password  string       `gorm:"not null;size:32"`
	Posts     []PostEntity `gorm:"foreignKey:user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name used by UserEntity, see Tabler interface of gorm
func (UserEntity) TableName() string {
	return "users"
}

type UserRepository interface {
	FindUserByUsername(username string) (models.User, error)
	FindUserByEmail(email string) (models.User, error)

	DeleteUserByUsername(username string) (bool, error)
	DeleteUserByEmail(email string) (bool, error)

	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}
