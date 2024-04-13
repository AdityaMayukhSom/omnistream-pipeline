package repository

import (
	"devstream.in/pixelated-pipeline/services/models"
)

type UserRepository interface {
	FindUserByUsername(username string)
	FindUserByEmail(email string)

	DeleteUserByUsername(username string)
	DeleteUserByEmail(email string)

	CreateUser(user models.User)
	UpdateUser(user models.User)
}
