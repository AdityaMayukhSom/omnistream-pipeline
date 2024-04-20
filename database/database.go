package database

import (
	postgresql "devstream.in/pixelated-pipeline/database/postgres"
	repository "devstream.in/pixelated-pipeline/database/repositories"
)

// data access objects or repositories hide all access to the data source.
// typically used per entity which
type Database interface {
	repository.UserRepository
	repository.PostRepository

	SetupDatabase()
	CleanupDatabase()
}

func NewDatabase() Database {
	return &postgresql.PostgresDatabase{}
}