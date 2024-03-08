package service

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type TodoService interface {
	Find(context.Context, int64) (*model.Todo, error)
	FindAll(context.Context) ([]*model.Todo, error)
	Create(context.Context, *pb.CreateTodoIn) (*model.Todo, error)
	Update(context.Context, *pb.UpdateTodoIn) (*model.Todo, error)
	Delete(context.Context, int64) error
}
