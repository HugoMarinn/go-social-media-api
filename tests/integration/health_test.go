package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HugoMarinn/go-social-media-api/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestHealthyEndpoint(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/api/v1/healthy", server.HealthyHandler)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/v1/healthy")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var body map[string]string
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "OK!", body["status"])
}
