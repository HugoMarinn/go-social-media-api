package server

import "net/http"

func mapRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/healthy", func(w http.ResponseWriter, r *http.Request) {

	})

	return mux
}
