package main

import (
	"context"
	"github.com/vizitiuRoman/libs/logger/example/internal/common/adapter/meta"
	"os"

	"github.com/vizitiuRoman/libs/logger"
	"github.com/vizitiuRoman/libs/logger/example/internal/common/adapter/log"
	"go.uber.org/fx"
)

func main() {
	fxApp := fx.New(
		fx.NopLogger,
		fx.Provide(context.Background),
		fx.Provide(func() (logger.Logger, error) {
			return logger.NewLogger(&logger.Config{
				CasinoID:     123321,
				Level:        "debug",
				Encoding:     "json",
				LevelEncoder: "console",
				CallerSkip:   nil,
			}, logger.WithContextEncoder(log.ContextEncoder))
		}),
		fx.Invoke(func(ctx context.Context, logger logger.Logger) {
			logger.Debug("application initiated")
			logger.AddContext(context.WithValue(ctx, meta.SessionID, "some-session-id")).Info("session id is specified")
			logger.AddContext(ctx).Info("session id is not specified")
			os.Exit(0)
		}),
	)

	fxApp.Run()
}
