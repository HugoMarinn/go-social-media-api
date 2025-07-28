package server

import (
	"net/http"

	"github.com/HugoMarinn/go-social-media-api/internal/auth"
	authHttp "github.com/HugoMarinn/go-social-media-api/internal/auth/delivery/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MapRoutes(authHandler auth.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(CorsMiddleware)
	r.Use(JSONResponseMiddleware)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/healthy", HealthyHandler)

		r.Route("/auth", func(r chi.Router) {
			authHttp.MapRoutes(r, authHandler)
		})
	})

	return r
}
