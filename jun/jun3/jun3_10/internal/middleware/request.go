package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceIDStr := r.Header.Get("X-Trace-ID")
		if traceIDStr == "" {
			if id, err := uuid.NewV7(); err == nil {
				traceIDStr = id.String()
			} else {
				traceIDStr = uuid.NewString()
			}
		}
		r.Header.Set("X-Request-ID", traceIDStr)
		w.Header().Set("X-Request-ID", traceIDStr)

		next.ServeHTTP(w, r)
	})
}
