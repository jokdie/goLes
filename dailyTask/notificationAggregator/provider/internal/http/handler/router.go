package handler

import (
	"net/http"
	"provider/internal/http/middleware"
)

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /provider/email", h.email)
	mux.HandleFunc("POST /provider/sms", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("POST /provider/push", func(w http.ResponseWriter, r *http.Request) {})

	var handler http.Handler = mux

	handler = middleware.ApplicationJsonMiddleware(mux)
	handler = middleware.RequestIDMiddleware(handler)

	return handler
}
