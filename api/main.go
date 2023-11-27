package main

import (
	"net/http"

	routes "github.com/AdityaMayukhSom/alex_mux_go/api/routes"
	logging "github.com/AdityaMayukhSom/alex_mux_go/internal/logging"
	chi "github.com/go-chi/chi/v5"
	zap "go.uber.org/zap"
)

func main() {
	logging.Init()
	logger := zap.L()

	var r *chi.Mux = chi.NewRouter()
	routes.Register(r)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		logger.Fatal(
			"could not instantiate ap logger for logging",
			zap.String("url", "localhost"),
			zap.Int("port", 8080),
		)

	}
}
