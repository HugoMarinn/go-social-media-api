package httphelper

import (
	"encoding/json"
	"net/http"
)

func ExtractRequestBody[T any](r *http.Request) (T, error) {
	var body T
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		var zero T
		return zero, err
	}
	return body, nil
}
