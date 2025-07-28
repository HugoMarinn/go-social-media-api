package httphelper

import (
	"encoding/json"
	"net/http"
)

type SuccessReponseDTO struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteSuccessJSON(w http.ResponseWriter, status int, msg string, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(SuccessReponseDTO{
		Message: msg,
		Code:    status,
		Data:    data,
	})
}
