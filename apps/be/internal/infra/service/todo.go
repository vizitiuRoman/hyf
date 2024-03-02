package service

import (
	"context"
	"fmt"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/domain/repo"
	"github.com/vizitiuRoman/hyf/internal/domain/service"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type todoService struct {
	logger log.Logger
	db     db.DB

	todoAdapterFactory adapter.TodoAdapterFactory
	todoRepoFactory    repo.TodoRepoFactory
	c                  redis.Cache
}

func NewTodoService(
	ctx context.Context,
	logger log.Logger,
	db db.DB,

	todoAdapterFactory adapter.TodoAdapterFactory,
	todoRepoFactory repo.TodoRepoFactory,

	c redis.Cache,
) service.TodoService {
	return &todoService{
		db: db,

		todoAdapterFactory: todoAdapterFactory,
		todoRepoFactory:    todoRepoFactory,

		logger: logger.WithComponent(ctx, "todo_service"),

		c: c,
	}
}

func (s *todoService) Find(ctx context.Context, id int64) (*model.Todo, error) {
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	todo, err := todoRepo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot create todo: %w", err)
	}

	return todo, nil
}

func (s *todoService) FindAll(ctx context.Context) ([]*model.Todo, error) {
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	todos, err := todoRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot find todos: %w", err)
	}

	return todos, nil
}

func (s *todoService) Create(ctx context.Context, input *pb.CreateTodoIn) (*model.Todo, error) {
	todoAdapter := s.todoAdapterFactory.Create(ctx)
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	todo, err := todoRepo.Create(ctx, todoAdapter.FromProto(input.Todo))
	if err != nil {
		return nil, fmt.Errorf("cannot create todo: %w", err)
	}

	return todo, nil
}

// Update --> @TODO --> return error when app is created with the same ID
func (s *todoService) Update(ctx context.Context, input *pb.UpdateTodoIn) (*model.Todo, error) {
	todoAdapter := s.todoAdapterFactory.Create(ctx)
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	todo, err := todoRepo.Update(ctx, todoAdapter.FromProto(input.Todo))
	if err != nil {
		return nil, fmt.Errorf("cannot update todo: %w", err)
	}

	return todo, nil
}

func (s *todoService) Delete(ctx context.Context, id int64) error {
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	err := todoRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot delete todo: %w", err)
	}

	return nil
}
