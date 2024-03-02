package adapter

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/middleware"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/middleware/auth"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/config"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/meta"
)

var Constructors = fx.Provide(
	config.NewConfig,

	NewFxLogger,
	NewFxMiddlewares,
	NewFxRedisCache,

	fx.Annotate(NewFxPgPool, fx.OnStop(func(db db.DB) error { return db.Close() })),
)

func NewFxRedisCache(lf fx.Lifecycle, ctx context.Context, logger log.Logger, cfg *config.Config) (redis.Cache, error) {
	c, err := redis.NewUniversalClient(ctx, cfg.Cache.Redis, logger)
	lf.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			c.Close()
			return nil
		},
	})

	return c, err
}

func NewFxLogger(cfg *config.Config) (log.Logger, error) {
	return log.NewLogger(cfg.Logger, zap.String(meta.AppVersionKey, cfg.Version+" "+cfg.BuildDate))
}

func NewFxPgPool(ctx context.Context, logger log.Logger, cfg *config.Config) (db.DB, error) {
	return db.NewPool(ctx, logger, cfg.DB)
}

func NewFxMiddlewares(ctx context.Context, logger log.Logger) []grpc.Middleware {
	return []grpc.Middleware{
		{
			Priority:    grpc.Priority(1000),
			Interceptor: auth.NewAuthInterceptor(ctx),
		},
		{
			Priority:    grpc.Priority(99999),
			Interceptor: middleware.LoggerMiddleware(ctx, logger),
		},
	}
}
