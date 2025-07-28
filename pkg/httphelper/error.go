package httphelper

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseDTO struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

func WriteErrorJSON(w http.ResponseWriter, msg string, status int, errors interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponseDTO{
		Message: msg,
		Code:    status,
		Errors:  errors,
	})
}
