package handler

import "net/http"

func NewRouter(h *UserHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", h.User)

	return mux
}
