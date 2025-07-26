package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HugoMarinn/go-social-media-api/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestHealthyEndpoint(t *testing.T) {
	handler := server.MapRoutes()

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/v1/healthy")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var body map[string]string
	err = json.NewDecoder(res.Body).Decode(&body)
	assert.NoError(t, err)
	assert.Equal(t, "OK!", body["status"])
}
