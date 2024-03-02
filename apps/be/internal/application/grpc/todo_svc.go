package grpc

import (
	"context"
	"errors"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/hyf/internal/domain"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/service"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type todoSVC struct {
	pb.UnimplementedTodoSVCServer
	ctx    context.Context
	logger log.Logger
	db     db.DB

	todoAdapterFactory adapter.TodoAdapterFactory

	todoService service.TodoService
}

func NewTodoSVCServerDescriptor(
	ctx context.Context,
	logger log.Logger,
	db db.DB,

	todoAdapterFactory adapter.TodoAdapterFactory,

	todoService service.TodoService,
) *grpc.ServerDescriptor {
	server := &todoSVC{
		ctx: ctx,
		db:  db,

		todoAdapterFactory: todoAdapterFactory,

		todoService: todoService,

		logger: logger,
	}

	return &grpc.ServerDescriptor{
		Server:               server,
		GRPCRegistrar:        pb.RegisterTodoSVCServer,
		GRPCGatewayRegistrar: pb.RegisterTodoSVCHandlerFromEndpoint,
		Methods: []grpc.MethodDescriptor{
			{
				Method: (*todoSVC).CreateTodo,
			},
			{
				Method: (*todoSVC).UpdateTodo,
			},
			{
				Method: (*todoSVC).DeleteTodo,
			},
			{
				Method: (*todoSVC).GetTodo,
			},
			{
				Method: (*todoSVC).GetTodos,
			},
		},
	}
}

func (s *todoSVC) CreateTodo(ctx context.Context, input *pb.CreateTodoIn) (*pb.CreateTodoOut, error) {
	todoAdapter := s.todoAdapterFactory.Create(ctx)

	todo, err := s.todoService.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTodoOut{Todo: todoAdapter.ToProto(todo)}, nil
}

func (s *todoSVC) UpdateTodo(ctx context.Context, input *pb.UpdateTodoIn) (*pb.UpdateTodoOut, error) {
	todo, err := s.todoService.Update(ctx, input)

	switch {
	case errors.Is(err, domain.ErrNotFound):
		return nil, status.Error(codes.NotFound, err.Error())

	case errors.Is(err, domain.ErrAlreadyExists):
		return nil, status.Error(codes.AlreadyExists, err.Error())

	case err != nil:
		return nil, err

	default:
		todoAdapter := s.todoAdapterFactory.Create(ctx)

		return &pb.UpdateTodoOut{Todo: todoAdapter.ToProto(todo)}, nil
	}
}

func (s *todoSVC) DeleteTodo(ctx context.Context, input *pb.DeleteTodoIn) (*emptypb.Empty, error) {
	err := s.todoService.Delete(ctx, input.Id)

	switch {
	case errors.Is(err, domain.ErrNotFound):
		return nil, status.Error(codes.NotFound, err.Error())

	case err != nil:
		return nil, err

	default:
		return &emptypb.Empty{}, nil
	}
}

func (s *todoSVC) GetTodo(ctx context.Context, input *pb.GetTodoIn) (*pb.GetTodoOut, error) {
	todo, err := s.todoService.Find(ctx, input.Id)

	switch {
	case errors.Is(err, domain.ErrNotFound):
		return nil, status.Error(codes.NotFound, err.Error())

	case err != nil:
		return nil, err

	default:
		todoAdapter := s.todoAdapterFactory.Create(ctx)

		return &pb.GetTodoOut{Todo: todoAdapter.ToProto(todo)}, nil
	}
}

func (s *todoSVC) GetTodos(ctx context.Context, _ *pb.GetTodosIn) (*pb.GetTodosOut, error) {
	todoAdapter := s.todoAdapterFactory.Create(ctx)

	todos, err := s.todoService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetTodosOut{Todos: todoAdapter.ToProtos(todos)}, nil
}
