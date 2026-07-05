package handler

import "net/http"

func NewRouter(h *HelloHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", h.Hello)

	return mux
}
