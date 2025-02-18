package main

import (
	"context"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/vizitiuRoman/libs/config/example/internal/common/adapter"
	"github.com/vizitiuRoman/libs/config/example/internal/common/adapter/config"
)

func main() {
	fxApp := fx.New(
		fx.Provide(func() context.Context { return context.Background() }),
		fx.StartTimeout(time.Second*3),
		fx.StopTimeout(time.Second*10),
		adapter.Constructors,
		fx.Provide(zap.NewExample),
		fx.Invoke(run),
	)

	fxApp.Run()
}

func run(
	lf fx.Lifecycle,
	config *config.Config,
	logger *zap.Logger,
) {
	lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("config initiated", zap.Any("config", config))
			return nil
		},
	})
}
