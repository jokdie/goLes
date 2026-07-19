package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Service interface {
	GetUser(ctx context.Context, id int) error
}

func Router(s Service) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
		rawId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(rawId)
		if err != nil {
			code := http.StatusBadRequest
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(ErrorResponse{Code: code, Message: "Bad Request"})
			return
		}

		if err := s.GetUser(r.Context(), id); err != nil {
			w.Header().Set("Content-type", "application/json")
			code := http.StatusInternalServerError
			message := err.Error()

			switch {
			case errors.Is(err, context.Canceled):
				w.WriteHeader(499)
				return
			case errors.Is(err, context.DeadlineExceeded):
				code = http.StatusGatewayTimeout
				message = "context deadline exceeded"
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(ErrorResponse{Code: code, Message: message})

			return
		}

		code := http.StatusNoContent
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(code)
	})

	return mux
}
