package funcenter

import "context"

type dummyMiddleware struct {
	onInit  func(context.Context) (context.Context, error)
	onEnter func(context.Context, ...interface{}) context.Context
	onLeave func(context.Context)
	onClose func(context.Context)
}

func (m dummyMiddleware) OnInit(ctx context.Context) (context.Context, error) {
	if m.onInit != nil {
		return m.onInit(ctx)
	}
	return ctx, nil
}

func (m dummyMiddleware) OnEnter(ctx context.Context, args ...interface{}) context.Context {
	if m.onEnter != nil {
		return m.onEnter(ctx, args...)
	}
	return ctx
}

func (m dummyMiddleware) OnLeave(ctx context.Context) {
	if m.onLeave != nil {
		m.onLeave(ctx)
	}
}

func (m dummyMiddleware) OnClose(ctx context.Context) {
	if m.onClose != nil {
		m.onClose(ctx)
	}
}

func OnInitMiddleware(fn func(context.Context) (context.Context, error)) Middleware {
	return dummyMiddleware{onInit: fn}
}

func OnEnterMiddleware(fn func(context.Context, ...interface{}) context.Context) Middleware {
	return dummyMiddleware{onEnter: fn}
}

func OnLeaveMiddleware(fn func(context.Context)) Middleware {
	return dummyMiddleware{onLeave: fn}
}

func OnCloseMiddleware(fn func(context.Context)) Middleware {
	return dummyMiddleware{onClose: fn}
}
