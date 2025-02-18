package adapter

import (
	"context"

	"github.com/vizitiuRoman/auth-service/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/auth-service/internal/common/adapter/server/middleware"
	"github.com/vizitiuRoman/auth-service/internal/common/adapter/server/middleware/auth"
	"github.com/vizitiuRoman/auth-service/internal/infra/adapter"
	"github.com/vizitiuRoman/auth-service/internal/infra/repo/cache"
	"github.com/vizitiuRoman/libs/cache/redis"
	"github.com/vizitiuRoman/libs/logger"
	"github.com/vizitiuRoman/libs/pgclient"
	"go.uber.org/fx"

	"github.com/vizitiuRoman/auth-service/internal/common/adapter/config"
)

var Constructors = fx.Provide(
	config.NewConfig,

	NewFxLogger,
	NewFxMiddlewares,
	NewFxRedisCache,

	fx.Annotate(NewFxPgPool, fx.OnStop(func(db pgclient.DB) error { return db.Close() })),
)

func NewFxRedisCache(lf fx.Lifecycle, ctx context.Context, logger logger.Logger, cfg *config.Config) (redis.Cache, error) {
	c, err := redis.NewRedisClient(ctx, cfg.Cache, logger)
	lf.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			c.Close()
			return nil
		},
	})

	return c, err
}

func NewFxLogger(cfg *config.Config) (logger.Logger, error) {
	return logger.NewLogger(cfg.Logger)
}

func NewFxPgPool(ctx context.Context, logger logger.Logger, cfg *config.Config) (pgclient.DB, error) {
	return pgclient.NewPool(ctx, cfg.DB, logger)
}

func NewFxMiddlewares(ctx context.Context, logger logger.Logger, rdb redis.Cache, authTokenRepoFactory *cache.AuthRepoFactory, adapter *adapter.AuthAdapter) []grpc.Middleware {
	return []grpc.Middleware{
		{
			Priority:    grpc.Priority(1000),
			Interceptor: auth.NewAuthInterceptor(ctx, rdb, authTokenRepoFactory, adapter),
		},
		{
			Priority:    grpc.Priority(99999),
			Interceptor: middleware.LoggerMiddleware(ctx, logger),
		},
	}
}
