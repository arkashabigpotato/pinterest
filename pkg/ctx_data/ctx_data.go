package ctx_data

import "context"

type UserData struct {
	UserID int
}

type CtxKey int

var key = CtxKey(1)

func ToContext(ctx context.Context, data UserData) context.Context {
	return context.WithValue(ctx, key, data)
}

func FromContext(ctx context.Context) UserData {
	if ctx.Value(key) == nil {
		return UserData{UserID: 0}
	}
	return ctx.Value(key).(UserData)
}
