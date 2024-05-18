package services

import (
	"fmt"

	"devstream.in/pixelated-pipeline/database"
	serviceConstant "devstream.in/pixelated-pipeline/services/constants"
	"devstream.in/pixelated-pipeline/services/models"
	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	// Returns a valid token for the user if the login is successful, i.e.
	// the credentials are valid and a nil error. If the credentials are
	// invalid or server error happens, returns a non nil error.
	LoginUser(credentials models.LoginCredential) (*models.Token, error)

	// Registers the user into the database. If a user with same credentials
	// already exists, returns a non null error value.
	RegisterUser(user models.User) error

	GetDetails(username string) (*models.User, error)
}

func NewUserService() UserService {
	return NewUserServiceImpl()
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type UserServiceImpl struct{}

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
		WithClaim(serviceConstant.ClaimKeyEmail, existingUser.Email),
		WithClaim(serviceConstant.ClaimKeyName, existingUser.Name),
		WithClaim(serviceConstant.ClaimKeyUsername, existingUser.Username),
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

func (us *UserServiceImpl) GetDetails(username string) (*models.User, error) {
	db := database.Init()
	user, err := db.FindUserByUsername(username)
	return &user, err
}
