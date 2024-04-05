package middlewares

// import (
// 	"encoding/json"
// 	"net/http"

// 	"devstream.in/blogs/auth"
// 	"devstream.in/blogs/config"
// 	"github.com/charmbracelet/log"
// )

// func WithAuthentication(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Info("authentication server hit")

// 		if r.Method == "OPTIONS" {
// 			return
// 		}

// 		_, err := auth.CheckTokenValidity(r, config.DefaultConfig.AccessSecretKey)
// 		if err != nil {
// 			json.NewEncoder(w).Encode("Unauthorized")
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
