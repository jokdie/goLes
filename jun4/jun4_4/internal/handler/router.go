package handler

import "net/http"

func NewRouter(h *ProfileHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /profile", h.Profile)

	return mux
}
