package application

import (
	"context"

	grpc2 "github.com/vizitiuRoman/hyf/internal/application/grpc"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/hyf/internal/infra/adapter"
	"github.com/vizitiuRoman/hyf/internal/infra/service"
	"github.com/vizitiuRoman/hyf/pkg/adapter/logger"
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

func NewFxTodoSVC(ctx context.Context, logger logger.Logger, adapter *adapter.TodoAdapter, todoService service.TodoService) DescriptorOut {
	return DescriptorOut{
		Option: grpc2.NewTodoSVCServerDescriptor(ctx, logger, adapter, todoService),
	}
}

func NewFxAuthSVC(ctx context.Context, logger logger.Logger, adapter *adapter.AuthTokenAdapter, authService service.AuthService) DescriptorOut {
	return DescriptorOut{
		Option: grpc2.NewAuthSVCServerDescriptor(ctx, logger, adapter, authService),
	}
}

//
//func NewFxUserSVC(ctx context.Context, logger log.Logger, db db.DB, todoAdapterFactory adapter.UserAdapterFactory, todoService service.UserService) DescriptorOut {
//	return DescriptorOut{
//		Option: grpc2.NewUserSVCServerDescriptor(ctx, logger, db, todoAdapterFactory, todoService),
//	}
//}
