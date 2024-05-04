package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"devstream.in/pixelated-pipeline/config"
	"devstream.in/pixelated-pipeline/services/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenService interface {
	GenerateToken(name, email, username interface{}) (*models.Token, error)
}

func NewTokenService() TokenService {
	return NewTokenServiceImpl()
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type TokenServiceImpl struct {
}

func NewTokenServiceImpl() *TokenServiceImpl {
	return &TokenServiceImpl{}
}

func (ts *TokenServiceImpl) ExtractToken(r *http.Request) (tokenStr string) {
	tokenHeaderStr := r.Header.Get("Authorization")
	strSlice := strings.Split(tokenHeaderStr, " ")
	if len(strSlice) == 2 {
		tokenStr = strSlice[1]
	}
	return tokenStr
}

func (ts *TokenServiceImpl) VerifyToken(tokenStr string, secretKey string) (*jwt.Token, error) {

	// custom function to return the
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// for the given token, which will be passes to this function, if the signing algorithm
		// for the parsed token is not of type HMAC, then it will return an error.
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// returns the key for validating the token, can be access token secret or refresh token
		// secret depending upon the usecase and the key to be verified.
		return []byte(secretKey), nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc)

	if err != nil {
		fmt.Println("Failed to verify token.")
		return nil, err
	}

	return token, nil
}

func (ts *TokenServiceImpl) CheckTokenValidity(tokenStr string, secretKey string) (token *jwt.Token, err error) {
	if token, err = ts.VerifyToken(tokenStr, secretKey); err != nil || !token.Valid {
		return nil, err
	}
	return token, nil
}

func (ts *TokenServiceImpl) GenerateToken(name, email, username interface{}) (*models.Token, error) {
	var err error

	accessSecretKey := config.GetAccessSecretKey()
	refreshSecretKey := config.GetRefreshSecretKey()

	tokenInfo := &models.Token{
		AccessUuid:    uuid.NewString(),
		RefreshUuid:   uuid.NewString(),
		AccessExpire:  time.Now().Add(time.Minute * 15).Unix(),
		RefreshExpire: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["access_uuid"] = tokenInfo.AccessUuid
	accessTokenClaims["user_name"] = name
	accessTokenClaims["user_email"] = email
	accessTokenClaims["user_username"] = username
	accessTokenClaims["exp"] = tokenInfo.AccessExpire
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	tokenInfo.AccessToken, err = accessToken.SignedString([]byte(accessSecretKey))

	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["refresh_uuid"] = tokenInfo.RefreshUuid
	refreshTokenClaims["user_name"] = name
	refreshTokenClaims["user_email"] = email
	refreshTokenClaims["user_username"] = username
	refreshTokenClaims["exp"] = tokenInfo.RefreshExpire
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenInfo.RefreshToken, err = refreshToken.SignedString([]byte(refreshSecretKey))

	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return tokenInfo, nil
}
