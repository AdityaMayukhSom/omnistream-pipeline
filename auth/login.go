package auth

import (
	"encoding/json"
	"net/http"
	"strconv"

	"devstream.in/blogs/models"
	"devstream.in/blogs/repositories"
	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	var receivedUser, existingUser models.User

	json.NewDecoder(r.Body).Decode(&receivedUser)

	results, err := repositories.Db.Query(
		"SELECT id, name, email, username, password FROM users WHERE email = $1;",
		receivedUser.Email,
	)

	if err != nil {
		log.Error("could not retrieve user details for given email", "email", receivedUser.Email)
		w.WriteHeader(http.StatusInternalServerError)
		enc, _ := json.Marshal(map[string]interface{}{
			"message": "could not login successfully, try again",
		})
		w.Write(enc)
		return
	}

	defer results.Close()

	if results.Next() {
		err = results.Scan(
			&existingUser.Id,
			&existingUser.Name,
			&existingUser.Email,
			&existingUser.Username,
			&existingUser.Password,
		)

		if err != nil {
			log.Error("could not parse user details for given email", "email", receivedUser.Email)
			w.WriteHeader(http.StatusInternalServerError)
			enc, _ := json.Marshal(map[string]interface{}{
				"message": "could not login successfully, try again",
			})
			w.Write(enc)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		enc, _ := json.Marshal(map[string]interface{}{
			"message": "there exists no user for given email",
		})
		w.Write(enc)
		return
	}

	pwMatchErr := bcrypt.CompareHashAndPassword(
		[]byte(existingUser.Password),
		[]byte(receivedUser.Password),
	)

	if pwMatchErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		enc, _ := json.Marshal(map[string]interface{}{
			"message": "password does not match",
		})
		w.Write(enc)
	}

	idInt, _ := strconv.Atoi(existingUser.Id)
	name := existingUser.Name
	email := existingUser.Email
	username := existingUser.Username

	tokenStruct, err := generateToken(idInt, name, email, username)

	if err != nil {
		log.Error("could not generate acces token for given email", "email", receivedUser.Email)
		w.WriteHeader(http.StatusInternalServerError)
		enc, _ := json.Marshal(map[string]interface{}{
			"message": "could not login successfully, try again",
		})
		w.Write(enc)
		return
	}

	accessToken := tokenStruct.AccessToken
	refreshToken := tokenStruct.RefreshToken
	cookie := http.Cookie{
		HttpOnly: true,
		Name:     "refreshToken",
		Value:    refreshToken,
		// Domain:   "jpoly1219devbox.xyz",
		Path: "/auth/",
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	enc, _ := json.Marshal(map[string]interface{}{
		"message":     "user successfully logged in",
		"accessToken": accessToken,
	})
	w.Write(enc)
}
