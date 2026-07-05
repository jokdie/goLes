package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf(
			"Начало запроса [%s]: %s %s",
			r.Header.Get("X-Request-ID"),
			r.Method,
			r.URL.Path,
		)

		wrappedWriter := &responseWriterWrapper{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrappedWriter, r)

		log.Printf(
			"[%d] %s %s за %s",
			wrappedWriter.statusCode,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
