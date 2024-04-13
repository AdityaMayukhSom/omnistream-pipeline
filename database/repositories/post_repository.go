package repository

import "devstream.in/pixelated-pipeline/services/models"

type PostRepository interface {
	FindPostById(id string)
	FindPostsByUsername(username string)

	DeletePostById(id string)

	CreatePost(post models.Post)
	UpdatePost(post models.Post)
}
