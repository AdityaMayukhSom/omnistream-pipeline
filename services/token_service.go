package services

import (
	"errors"
	"fmt"
	"time"

	"devstream.in/pixelated-pipeline/config"
	serviceConstant "devstream.in/pixelated-pipeline/services/constants"
	"devstream.in/pixelated-pipeline/services/models"
	"github.com/charmbracelet/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Use [WithClaim] function to generate token claim. Used to pass key value pairs
// which will be added to the token during generation of the token.
type TokenClaim struct {
	key   string
	value any
}

// Utility function to use like functional options pattern for generating the
// token with the given claims.
func WithClaim(key string, value any) TokenClaim {
	return TokenClaim{key: key, value: value}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type TokenService interface {
	// Generates a new instance of token containing both access token and refresh token.
	// It also adds the claims which were passed as a list to the method.
	GenerateToken(claims ...TokenClaim) (*models.Token, error)

	// Validates the [tokenStr] with the [secretKey]. If the token is valid and not expired
	// yet, returns a nil error along with the username of the person. Otherwise if the token
	// cannot be parsed, expired or the algorithm used to sign does not match, returns an non
	// nil error.
	ValidateToken(tokenStr string, secretKey string) (string, error)
}

// Returns a new instance of a specific implementation of the token service interface.
func NewTokenService() TokenService {
	return NewTokenServiceImpl()
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// A specific implementation of the token service interface.
type TokenServiceImpl struct{}

func NewTokenServiceImpl() *TokenServiceImpl {
	return &TokenServiceImpl{}
}

func (ts *TokenServiceImpl) ValidateToken(tokenStr string, secretKey string) (string, error) {
	// Custom function to return the key which was used to sign the jwt.Token
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// The given tokenStr will be parsed and passed to this function. If the signing algorithm
		// of the parsed token is not of type HMAC, it will return an error.
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		// Returns the key for validating the token. It can be access token secret or refresh token
		// secret depending upon the usecase and the key to be verified.
		return []byte(secretKey), nil
	}

	// if the tokenStr is valid, it will be parsed without any error
	token, err := jwt.Parse(tokenStr, keyFunc)

	if err != nil {
		log.Error("failed to verify token", "token", tokenStr)
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims[serviceConstant.ClaimKeyUsername].(string), nil
}

func (ts *TokenServiceImpl) GenerateToken(claims ...TokenClaim) (tokenInfo *models.Token, err error) {
	accessSecretKey := config.GetAccessSecretKey()
	refreshSecretKey := config.GetRefreshSecretKey()

	tokenInfo = &models.Token{
		AccessUuid:    uuid.NewString(),
		RefreshUuid:   uuid.NewString(),
		AccessExpire:  time.Now().Add(time.Minute * 15),
		RefreshExpire: time.Now().Add(time.Hour * 24 * 7),
	}

	accessTokenClaims := jwt.MapClaims{}

	for _, claim := range claims {
		accessTokenClaims[claim.key] = claim.value
	}

	accessTokenClaims[serviceConstant.ClaimKeyAuthorized] = true
	accessTokenClaims[serviceConstant.ClaimKeyAccessUUID] = tokenInfo.AccessUuid
	accessTokenClaims[serviceConstant.ClaimKeyExpiration] = tokenInfo.AccessExpire
	accessTokenClaims[serviceConstant.ClaimKeyIssuer] = serviceConstant.ClaimValueIssuer

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	tokenInfo.AccessToken, err = accessToken.SignedString([]byte(accessSecretKey))

	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshTokenClaims := jwt.MapClaims{}

	for _, claim := range claims {
		refreshTokenClaims[claim.key] = claim.value
	}

	refreshTokenClaims[serviceConstant.ClaimKeyRefreshUUID] = tokenInfo.RefreshUuid
	refreshTokenClaims[serviceConstant.ClaimKeyExpiration] = tokenInfo.RefreshExpire
	refreshTokenClaims[serviceConstant.ClaimKeyIssuer] = serviceConstant.ClaimValueIssuer

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenInfo.RefreshToken, err = refreshToken.SignedString([]byte(refreshSecretKey))

	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return tokenInfo, nil
}
