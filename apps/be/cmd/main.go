package main

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/vizitiuRoman/hyf/internal/application"
	"github.com/vizitiuRoman/hyf/internal/common/adapter"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/config"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/hyf/internal/infra"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(func() context.Context { return context.Background() }),

		fx.StartTimeout(time.Second*3),
		fx.StopTimeout(time.Second*10),

		adapter.Constructors,
		infra.Constructors,

		fx.Provide(runtime.NewServeMux),
		fx.Provide(grpc.NewGRPC),

		fx.Provide(func(cfg *config.Config) *grpc.Config { return cfg.Server }),

		fx.Provide(application.NewFxTodoSVC),
		fx.Provide(application.NewFxAuthSVC),

		fx.Provide(application.NewFxGRPC),

		fx.Invoke(newServer),
	)

	app.Run()
}

func newServer(
	lf fx.Lifecycle,
	server grpc.Server,
	serverDescriptors []*grpc.ServerDescriptor,
) error {
	err := server.Register(serverDescriptors)
	if err != nil {
		return err
	}

	lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err = server.Start(); err != nil {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Stop()
			return nil
		},
	})

	return nil
}
