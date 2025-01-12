package health

import "context"

type ctxKey string

const ctxAuthSuccessKey ctxKey = "github.com/alexliesenfeld/health#authenticationSuccess"

func getAuthResult(ctx context.Context) *bool {
	authOK, ok := ctx.Value(ctxAuthSuccessKey).(bool)
	if !ok {
		return nil
	}
	return &authOK
}

func withAuthResult(ctx context.Context, value bool) context.Context {
	return context.WithValue(ctx, ctxAuthSuccessKey, value)
}
