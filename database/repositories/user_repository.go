package repository

import (
	"devstream.in/pixelated-pipeline/services/models"
)

type UserRepository interface {
	FindByUsername(username string)
	FindByEmail(email string)

	DeleteByUsername(username string)
	DeleteByEmail(email string)

	Save(user models.User)
}
