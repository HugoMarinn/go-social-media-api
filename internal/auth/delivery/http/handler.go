package http

import (
	"errors"
	"net/http"

	"github.com/HugoMarinn/go-social-media-api/internal/auth"
	"github.com/HugoMarinn/go-social-media-api/internal/auth/usecase"
	"github.com/HugoMarinn/go-social-media-api/pkg/httphelper"
)

type AuthHandler struct {
	uc auth.UseCase
}

func NewAuthHandler(uc auth.UseCase) auth.Handler {
	return &AuthHandler{uc}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := r.Context().Value("validatedBody").(auth.RegisterRequestDTO)

		newUser, err := h.uc.Register(&payload)
		if err != nil {
			if errors.Is(err, usecase.ErrEmailAlreadyTaken) {
				httphelper.WriteErrorJSON(
					w,
					"this email has already taken, try another...",
					http.StatusBadRequest,
					make([]string, 0),
				)
				return
			}
			httphelper.WriteErrorJSON(
				w,
				"we got a unexpected error on our server, sorry...",
				http.StatusInternalServerError,
				make([]string, 0),
			)
			return
		}

		httphelper.WriteSuccessJSON(
			w,
			http.StatusCreated,
			"user registered successfully",
			newUser,
		)
	}
}
