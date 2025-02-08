package application

import (
	"context"

	grpc2 "github.com/vizitiuRoman/auth-service/internal/application/grpc"
	"github.com/vizitiuRoman/auth-service/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/auth-service/internal/infra/adapter"
	"github.com/vizitiuRoman/auth-service/internal/infra/service"
	"github.com/vizitiuRoman/libs/logger"
	"go.uber.org/fx"
)

type DescriptorOut struct {
	fx.Out
	Option *grpc.ServerDescriptor `group:"fx_descriptors"`
}

type FxGrpcIn struct {
	fx.In
	Descriptors []*grpc.ServerDescriptor `group:"fx_descriptors"`
}

func NewFxGRPC(in FxGrpcIn) []*grpc.ServerDescriptor {
	return in.Descriptors
}

func NewFxAuthSVC(ctx context.Context, logger logger.Logger, adapter *adapter.AuthAdapter, authService service.AuthService) DescriptorOut {
	return DescriptorOut{
		Option: grpc2.NewAuthSVCServerDescriptor(ctx, logger, adapter, authService),
	}
}
