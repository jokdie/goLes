package userid

import "context"

type contextKey struct{}

var UserIDCtxKey contextKey

func GetUserIDCtxKey(ctx context.Context) string {
	if id, ok := ctx.Value(UserIDCtxKey).(string); ok {
		return id
	}

	return ""
}
