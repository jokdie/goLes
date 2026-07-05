package handler

import "net/http"

func NewRouter(h *TestHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", h.Do)

	return mux
}
