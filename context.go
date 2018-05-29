package noxus_sdk

import "context"

type appClientKey struct{}

func NewContext(ctx context.Context, appClient *AppClient) context.Context {
	return context.WithValue(ctx, appClientKey{}, appClient)
}

func FromContext(ctx context.Context) *AppClient {
	return ctx.Value(appClientKey{}).(*AppClient)
}
