package service

import (
	"fmt"

	"devstream.in/pixelated-pipeline/database"
	"devstream.in/pixelated-pipeline/services/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	LoginUser(credentials models.LoginCredential) models.Token
}

type UserServiceImpl struct {
}

func (us *UserServiceImpl) LoginUser(credentials models.LoginCredential) (*models.Token, error) {
	database := database.Init()
	existingUser, err := database.FindUserByUsername(credentials.Username)

	if err != nil {
		return nil, fmt.Errorf("provided credentials did not exist")
	}

	pwMatchErr := bcrypt.CompareHashAndPassword(
		[]byte(existingUser.Password),
		[]byte(credentials.Password),
	)

	if pwMatchErr != nil {
		return nil, fmt.Errorf("provided credentials do not match")
	}

	tokenService := NewTokenService()
	tokenStruct, err := tokenService.GenerateToken(
		existingUser.Name,
		existingUser.Email,
		existingUser.Username,
	)

	if err != nil {
		return nil, fmt.Errorf("could not generate acces token for given username %s", credentials.Username)
	}

	return tokenStruct, nil
}
