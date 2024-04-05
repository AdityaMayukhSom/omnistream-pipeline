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

func ExtractToken(r *http.Request) string {
	tokenHeaderStr := r.Header.Get("Authorization")
	fmt.Println(tokenHeaderStr)
	strSlice := strings.Split(tokenHeaderStr, " ")
	var tokenStr string
	if len(strSlice) == 2 {
		tokenStr = strSlice[1]
	}

	fmt.Println(tokenStr)
	return tokenStr
}

func verifyToken(r *http.Request, secretKey string) (*jwt.Token, error) {
	tokenStr := ExtractToken(r)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("Failed to verify token.")
		return nil, err
	}

	return token, nil
}

func CheckTokenValidity(r *http.Request, secretKey string) (*jwt.Token, error) {
	token, err := verifyToken(r, secretKey)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		fmt.Println("Invalid token.")
		return nil, err
	}
	return token, nil
}

func generateToken(name, email, username interface{}) (*models.Token, error) {
	var err error

	accessSecretKey := config.DefaultConfig.AccessSecretKey
	refreshSecretKey := config.DefaultConfig.RefreshSecretKey

	tokenInfo := &models.Token{}
	tokenInfo.AccessUuid = uuid.NewString()
	tokenInfo.AccessExpire = time.Now().Add(time.Minute * 15).Unix()
	tokenInfo.RefreshUuid = uuid.NewString()
	tokenInfo.RefreshExpire = time.Now().Add(time.Hour * 24 * 7).Unix()

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
	refreshTokenClaims["refresh_uuid"] = uuid.NewString()
	refreshTokenClaims["user_username"] = username
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenInfo.RefreshToken, err = refreshToken.SignedString([]byte(refreshSecretKey))
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return tokenInfo, nil
}
