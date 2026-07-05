package middleware

import (
	"context"
	"jun4_4/internal/userid"
	"net/http"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")

		if userID == "" {
			userID = "guest"
		}

		ctx := context.WithValue(r.Context(), userid.UserIDCtxKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
