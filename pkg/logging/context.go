package logging

import "context"

type ctxLogger struct{}

func ContextWithLogger(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxLogger{}, newLogger())
}

func Logger(ctx context.Context) Log {
	if l, ok := ctx.Value(ctxLogger{}).(*logger); ok {
		return l
	}
	return newLogger()
}
