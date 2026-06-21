package requestid

import "context"

type contextKey struct{}

var RequestIDKey contextKey

func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(RequestIDKey).(string); ok {
		return id
	}

	return ""
}
