package service

import (
	"context"
	"fmt"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/infra/adapter"
	"github.com/vizitiuRoman/hyf/internal/infra/repo"
	"github.com/vizitiuRoman/hyf/pkg/adapter/logger"
	"github.com/vizitiuRoman/hyf/pkg/adapter/pgclient"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type TodoService struct {
	logger logger.Logger
	db     pgclient.DB

	adapter         *adapter.TodoAdapter
	todoRepoFactory *repo.TodoRepoFactory
}

func NewTodoService(
	ctx context.Context,
	logger logger.Logger,
	db pgclient.DB,

	adapter *adapter.TodoAdapter,
	todoRepoFactory *repo.TodoRepoFactory,
) *TodoService {
	return &TodoService{
		logger: logger.WithComponent(ctx, "todo_service"),

		db: db,

		adapter:         adapter,
		todoRepoFactory: todoRepoFactory,
	}
}

func (s *TodoService) Find(ctx context.Context, id int64) (*model.Todo, error) {
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	todo, err := todoRepo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot find todo: %w", err)
	}

	return todo, nil
}

func (s *TodoService) FindAll(ctx context.Context) ([]*model.Todo, error) {
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	todos, err := todoRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot find todos: %w", err)
	}

	return todos, nil
}

func (s *TodoService) Create(ctx context.Context, input *pb.CreateTodoIn) (*model.Todo, error) {
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	todo, err := todoRepo.Create(ctx, s.adapter.FromProto(input.Todo))
	if err != nil {
		return nil, fmt.Errorf("cannot create todo: %w", err)
	}

	return todo, nil
}

// Update --> @TODO --> return error when app is created with the same ID
func (s *TodoService) Update(ctx context.Context, input *pb.UpdateTodoIn) (*model.Todo, error) {
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	todo, err := todoRepo.Update(ctx, s.adapter.FromProto(input.Todo))
	if err != nil {
		return nil, fmt.Errorf("cannot update todo: %w", err)
	}

	return todo, nil
}

func (s *TodoService) Delete(ctx context.Context, id int64) error {
	todoRepo := s.todoRepoFactory.Create(ctx, s.db)

	err := todoRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot delete todo: %w", err)
	}

	return nil
}
