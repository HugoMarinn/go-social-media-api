package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(CorsMiddleware)
	r.Use(JSONResponseMiddleware)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/healthy", HealthyHandler)
	})

	return r
}
