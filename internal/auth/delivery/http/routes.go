package http

import (
	"github.com/HugoMarinn/go-social-media-api/internal/auth"
	"github.com/HugoMarinn/go-social-media-api/pkg/middlewares"
	"github.com/go-chi/chi/v5"
)

func MapRoutes(r chi.Router, h auth.Handler) {
	r.Post(
		"/register",
		middlewares.ValidateRequestBody[auth.RegisterRequestDTO](h.Register()),
	)
}
