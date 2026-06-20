package handler

import (
	"net/http"
)

func NewRouter(h *UserHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", h.User)

	return mux
}
