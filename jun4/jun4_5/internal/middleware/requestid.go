package middleware

import (
	"context"
	"jun4_5/internal/requestid"
	"net/http"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		ctx := context.WithValue(r.Context(), requestid.RequestIDKey, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
