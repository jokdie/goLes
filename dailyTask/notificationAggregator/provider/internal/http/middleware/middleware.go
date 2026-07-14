package middleware

import (
	"context"
	"encoding/json"
	"log/slog"
	"mime"
	"net/http"
	"provider/internal/model"
	"provider/internal/requestid"

	"github.com/google/uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-GUID")

		if requestID == "" {
			requestID = uuid.NewString()
		}

		w.Header().Set("X-GUID", requestID)

		ctx := context.WithValue(r.Context(), requestid.Key, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ApplicationJsonMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
			if err != nil || mediaType != "application/json" {
				errCode := http.StatusUnsupportedMediaType

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(errCode)

				data := model.ErrorResponse{Code: errCode, Message: "Unsupported Media Type"}
				if err := json.NewEncoder(w).Encode(data); err != nil {
					logger.Error(
						"failed to write json response",
						slog.Any("error", err),
					)
				}
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
