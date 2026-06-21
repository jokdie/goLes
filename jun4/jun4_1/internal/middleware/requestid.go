package middleware

import (
	"context"
	"jun4_1/internal/requestid"
	"net/http"

	"github.com/google/uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		r.Header.Set("X-Request-ID", requestID)
		w.Header().Set("X-Request-ID", requestID)

		ctx := context.WithValue(r.Context(), requestid.RequestIDKey, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
