package server

import (
	"encoding/json"
	"net/http"
)

type healthyReponseDTO struct {
	Status string `json:"status"`
}

func HealthyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthyReponseDTO{
		Status: "OK!",
	})
}
