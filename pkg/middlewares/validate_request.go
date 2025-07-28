package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HugoMarinn/go-social-media-api/pkg/httphelper"
	"github.com/HugoMarinn/go-social-media-api/pkg/validatorhelper"
)

func ValidateRequestBody[T any](next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload T

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			httphelper.WriteErrorJSON(
				w,
				"JSON not found or invalid",
				http.StatusBadRequest,
				make([]string, 0),
			)
			return
		}

		if err := validatorhelper.Validate.Struct(payload); err != nil {
			httphelper.WriteErrorJSON(
				w,
				"payload data is invalid",
				http.StatusUnprocessableEntity,
				validatorhelper.FormatValidationErrors(err),
			)
			return
		}

		ctx := context.WithValue(r.Context(), "validatedBody", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
