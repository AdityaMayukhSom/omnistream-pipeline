package repository

import "devstream.in/pixelated-pipeline/services/models"

type PostRepository interface {
	FindByUsername(username string)
	FindByEmail(email string)

	DeleteByUsername(username string)
	DeleteByEmail(email string)

	Save(post models.Post)
}
