package handler

import "net/http"

func NewRouter(h *UserHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user/{id}", h.User)

	return mux
}
