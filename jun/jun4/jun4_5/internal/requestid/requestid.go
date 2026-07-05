package requestid

import "context"

type requestIDKey struct{}

var RequestIDKey requestIDKey

func GetRequestID(ctx context.Context) string {
	if requestId, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestId
	}

	return ""
}
