package middleware

import (
	"context"
	"net/http"
	"provider/internal/requestid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-GUID")

		w.Header().Set("X-GUID", requestID)

		ctx := context.WithValue(r.Context(), requestid.Key, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
