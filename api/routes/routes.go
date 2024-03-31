package routes

import (
	"github.com/AdityaMayukhSom/alex_mux_go/api/middleware"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Register(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		r.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
