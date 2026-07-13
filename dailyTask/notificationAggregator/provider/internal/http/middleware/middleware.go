package middleware

import (
	"context"
	"encoding/json"
	"log"
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

func ApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil || mediaType != "application/json" {
			errCode := http.StatusUnsupportedMediaType

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(errCode)

			data := model.ErrorResponse{Code: errCode, Message: "Unsupported Media Type"}
			if err := json.NewEncoder(w).Encode(data); err != nil {
				log.Printf("Ошибка записи json: %s", err)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}
