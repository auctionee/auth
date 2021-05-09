package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ctxlog struct{}

func WithLogger(ctx context.Context, logger *logrus.Logger) context.Context {
	return context.WithValue(ctx, ctxlog{}, logger)
}

func GetLogger(ctx context.Context) *logrus.Logger {
	l, ok := ctx.Value(ctxlog{}).(*logrus.Logger)
	if !ok {
		return logrus.New()
	}
	return l
}
func NewCtxWithLogger() context.Context{
	return context.WithValue(context.Background(), ctxlog{},logrus.New() )
}