package main

import (
	"context"

	"github.com/vizitiuRoman/libs/pgclient"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(context.Background),
		fx.Provide(zap.NewExample),
		fx.Provide(NewFxPool),
		fx.Invoke(run),
	).Run()
}

func run(
	lf fx.Lifecycle,
	logger *zap.Logger,
	db pgclient.DB,
) {
	lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("application initiated")

			r, err := db.Query("SELECT now() AS now")
			if err != nil {
				logger.Error("unable to execute query", zap.Error(err))
				return err
			}

			var v any
			if r.Next() {
				if err = r.Scan(&v); err != nil {
					logger.Info("unable to scan result", zap.Error(err))
					return err
				}
			}

			logger.Info("query result", zap.Any("result", v))

			if err = r.Close(); err != nil {
				logger.Info("unable to close result", zap.Error(err))
				return err
			}

			if err := db.Close(); err != nil {
				logger.Info("unable to close db", zap.Error(err))
				return err
			}

			return nil
		},
	})
}
