package service

import (
	"fmt"

	"devstream.in/pixelated-pipeline/database"
	"devstream.in/pixelated-pipeline/services/models"
	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	LoginUser(credentials models.LoginCredential) (*models.Token, error)
	RegisterUser(user models.User) error
}

func NewUserService() UserService {
	return NewUserServiceImpl()
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (us *UserServiceImpl) LoginUser(credentials models.LoginCredential) (*models.Token, error) {
	db := database.Init()
	existingUser, err := db.FindUserByUsername(credentials.Username)

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

func (us *UserServiceImpl) RegisterUser(user models.User) (err error) {
	var db database.Database = database.Init()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error("could not hash given password", "password", user.Password)
		return err
	}

	// TODO: technically a new object should be created using the hashed password
	// and that object should be transferred to repository layer, but this is for
	// convinience...
	user.Password = string(hashedPassword)
	_, err = db.CreateUser(user)
	return err
}
