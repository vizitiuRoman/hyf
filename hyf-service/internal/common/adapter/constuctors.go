package adapter

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/middleware"
	"github.com/vizitiuRoman/hyf/pkg/adapter/logger"
	"github.com/vizitiuRoman/hyf/pkg/adapter/pgclient"
	"go.uber.org/fx"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/config"
)

var Constructors = fx.Provide(
	config.NewConfig,

	NewFxLogger,
	NewFxMiddlewares,

	fx.Annotate(NewFxPgPool, fx.OnStop(func(db pgclient.DB) error { return db.Close() })),
)

func NewFxLogger(cfg *config.Config) (logger.Logger, error) {
	return logger.NewLogger(cfg.Logger)
}

func NewFxPgPool(ctx context.Context, logger logger.Logger, cfg *config.Config) (pgclient.DB, error) {
	return pgclient.NewPool(ctx, cfg.DB, logger)
}

func NewFxMiddlewares(ctx context.Context, logger logger.Logger) []grpc.Middleware {
	return []grpc.Middleware{
		{
			Priority:    grpc.Priority(99999),
			Interceptor: middleware.LoggerMiddleware(ctx, logger),
		},
	}
}
