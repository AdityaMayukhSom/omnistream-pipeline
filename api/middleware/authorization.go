package middleware

import (
	"errors"
	"net/http"

	"github.com/AdityaMayukhSom/alex_mux_go/api/handlers"
	"github.com/AdityaMayukhSom/alex_mux_go/internal/database"
	"go.uber.org/zap"
)

var (
	ErrorEmptyCredentials = errors.New("username or token empty")
	ErrorUsernameNotFound = errors.New("username not found")
	ErrorInvalidToken     = errors.New("token invalid")
	ErrorUnauthorized     = errors.New("unauthorised access blocked")
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := zap.L()

		username := r.URL.Query().Get("username")
		authToken := r.Header.Get("Authorization")

		if username == "" || authToken == "" {
			logger.Error(ErrorEmptyCredentials.Error())
			handlers.RequestErrorHandler(w, ErrorEmptyCredentials)
			return
		}

		db := database.GetInstance(database.POSTGRESQL)
		if db == nil {
			logger.Error("could not connect to database")
			handlers.InternalErrorHandler(w)
			return
		}

		credentials := db.GetCredentials(username)
		if credentials == nil {
			logger.Info(ErrorUsernameNotFound.Error())
			handlers.NotFoundErrorHandler(w, ErrorUsernameNotFound)
			return
		}

		if credentials.AuthToken != authToken {
			handlers.RequestErrorHandler(w, ErrorUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
