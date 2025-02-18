package main

import (
	"context"
	"time"

	"github.com/vizitiuRoman/libs/pgclient"
	"go.uber.org/zap"
)

type loggingHook struct {
	logger *zap.Logger
	start  time.Time
}

type loggingHookFactory struct {
	logger *zap.Logger
}

func (h *loggingHook) Before(context.Context) {
	h.start = time.Now()
	h.logger.Info("before hook")
}

func (h *loggingHook) After(ctx context.Context, query string, args ...any) {
	h.logger.With(zap.String("go.method", "After")).Debug(
		"sql query execution",
		zap.String("sql.query.text", query),
		zap.Any("sql.query.args.any", args),
		zap.String("sql.query.duration.string", time.Since(h.start).String()),
	)
}

func NewLoggingHookFactory(logger *zap.Logger) pgclient.HookFactory {
	return &loggingHookFactory{
		logger: logger,
	}
}

func (f *loggingHookFactory) Create() pgclient.Hook {
	return &loggingHook{
		logger: f.logger,
	}
}
