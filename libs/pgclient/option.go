package pgclient

import (
	"context"
	"go.uber.org/zap"
	"time"
)

type Option func(cfg *pgDB)

func WithHookFactories(hookFactories []HookFactory) Option {
	return func(cln *pgDB) {
		cln.hookFactories = hookFactories
	}
}

func WithHookFactory(hookFactory HookFactory) Option {
	return func(cln *pgDB) {
		cln.hookFactories = append(cln.hookFactories, hookFactory)
	}
}

func WithTransactionLogging(log Logger) Option {
	return func(cln *pgDB) {
		cln.hookFactories = append(cln.hookFactories, newTransactionLoggingHook(log))
	}
}

type transactionLoggingHook struct {
	log   Logger
	start time.Time
}

type transactionLoggingHookFactory struct {
	log Logger
}

func (h *transactionLoggingHook) Before(context.Context) {
	h.start = time.Now()
}

func (h *transactionLoggingHook) After(ctx context.Context, query string, args ...any) {
	h.log.Debug(
		"sql query execution",
		zap.String("go.method", "After"),
		zap.String("debug.sql.query.string", query),
		zap.Any("debug.sql.query.args.object", args),
		zap.String("debug.sql.query.duration.string", time.Since(h.start).String()),
	)
}

func newTransactionLoggingHook(log Logger) HookFactory {
	return &transactionLoggingHookFactory{
		log: log,
	}
}

func (f *transactionLoggingHookFactory) Create() Hook {
	return &transactionLoggingHook{
		log: f.log,
	}
}
