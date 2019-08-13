package zap // import "moul.io/funcenter/logging/zap"

import (
	"context"

	"go.uber.org/zap"
	"moul.io/funcenter"
)

type middleware struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) funcenter.Middleware {
	return middleware{logger: logger}
}

func (m middleware) OnClose(ctx context.Context) {
	m.logger.Info("OnClose")
}

func (m middleware) OnEnter(ctx context.Context, args ...interface{}) context.Context {
	m.logger.Info("OnEnter")
	return ctx
}

func (m middleware) OnLeave(ctx context.Context) {
	m.logger.Info("OnLeave")
}

func (m middleware) OnInit(ctx context.Context) (context.Context, error) {
	m.logger.Info("OnInit")
	return ctx, nil
}
