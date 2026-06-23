package handler

import "net/http"

func NewRouter(h *WorkHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /work", h.Work)

	return mux
}
