package application

import (
	"context"

	grpc2 "github.com/vizitiuRoman/hyf/internal/application/grpc"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/service"
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

func NewFxTodoSVC(ctx context.Context, logger log.Logger, db db.DB, todoAdapterFactory adapter.TodoAdapterFactory, todoService service.TodoService) DescriptorOut {
	return DescriptorOut{
		Option: grpc2.NewTodoSVCServerDescriptor(ctx, logger, db, todoAdapterFactory, todoService),
	}
}

//
//func NewFxUserSVC(ctx context.Context, logger log.Logger, db db.DB, todoAdapterFactory adapter.UserAdapterFactory, todoService service.UserService) DescriptorOut {
//	return DescriptorOut{
//		Option: grpc2.NewUserSVCServerDescriptor(ctx, logger, db, todoAdapterFactory, todoService),
//	}
//}
