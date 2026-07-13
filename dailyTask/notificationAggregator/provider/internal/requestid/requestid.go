package requestid

import "context"

type requestIDKey struct{}

var Key = requestIDKey{}

func GetRequestID(ctx context.Context) string {
	if requestId, ok := ctx.Value(Key).(string); ok {
		return requestId
	}

	return ""
}
