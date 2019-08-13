package funcenter

import (
	"context"
)

type contextKey string

const middlewaresContextKey = contextKey("middlewares")

func New(ctx context.Context, middlewares ...Middleware) (context.Context, error) {
	ctx = context.WithValue(ctx, middlewaresContextKey, middlewares)
	var err error
	for _, middleware := range GetMiddlewares(ctx) {
		if ctx, err = middleware.OnInit(ctx); err != nil {
			return nil, err
		}
	}
	return ctx, nil
}

func Enter(ctx context.Context, args ...interface{}) context.Context {
	for _, middleware := range GetMiddlewares(ctx) {
		ctx = middleware.OnEnter(ctx, args...)
	}
	return ctx
}

func Leave(ctx context.Context) {
	for _, middleware := range GetMiddlewares(ctx) {
		middleware.OnLeave(ctx)
	}
}

func Close(ctx context.Context) {
	for _, middleware := range GetMiddlewares(ctx) {
		middleware.OnClose(ctx)
	}
}

func GetMiddlewares(ctx context.Context) []Middleware {
	return ctx.Value(middlewaresContextKey).([]Middleware)
}
