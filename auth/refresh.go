package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"devstream.in/blogs/config"
	"devstream.in/blogs/repositories"
	"github.com/dgrijalva/jwt-go"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	// get refreshToken from cookie
	var refreshToken, refreshUsername string

	c, err := r.Cookie("refreshToken")

	if err != nil {
		w.Write([]byte("error in reading cookie : " + err.Error() + "\n"))
	} else {
		refreshToken = c.Value
	}

	// verify refreshToken and check its validity
	refreshSecretKey := config.DefaultConfig.RefreshSecretKey

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(refreshSecretKey), nil
	})

	if err != nil {
		fmt.Println("Failed to verify token.")
		return
	}

	if !token.Valid {
		fmt.Println("Invalid token.")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		fmt.Println(err)
	}
	refreshUsername = claims["user_username"].(string)

	// query the user and use that data as parameters for generating new tokens
	retrievedUser, err := repositories.RetrieveUserByUsername(refreshUsername)

	if err != nil {

	}

	// generate new access and refresh tokens
	tokenStruct, err := generateToken(
		retrievedUser.Name,
		retrievedUser.Email,
		retrievedUser.Username,
	)

	if err != nil {
		fmt.Println("failed to generate token")
	}

	accessToken := tokenStruct.AccessToken
	refreshToken = tokenStruct.RefreshToken
	cookie := http.Cookie{
		HttpOnly: true,
		Name:     "refreshToken",
		Value:    refreshToken,
		// Domain:   "jpoly1219devbox.xyz",
		Path: "/auth/",
	}

	responseJSON := map[string]string{
		"accessToken": accessToken,
	}

	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(responseJSON)
}
