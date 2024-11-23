package adapter

import (
	"context"
	"time"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/pkg/adapter/logger"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"github.com/vizitiuRoman/hyf/pkg/gen/sqlboiler/hyfdb"
	"github.com/volatiletech/null/v8"
)

type TodoAdapter struct {
	logger logger.Logger
}

func NewTodoAdapter(ctx context.Context, logger logger.Logger) *TodoAdapter {
	return &TodoAdapter{
		logger: logger.WithComponent(ctx, "TodoAdapter"),
	}
}

func (a *TodoAdapter) FromProto(todo *pb.Todo) *model.Todo {
	return &model.Todo{
		ID:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

func (a *TodoAdapter) ToProto(todo *model.Todo) *pb.Todo {
	return &pb.Todo{
		Id:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

func (a *TodoAdapter) ToProtos(todos []*model.Todo) []*pb.Todo {
	output := make([]*pb.Todo, 0, len(todos))

	for _, todo := range todos {
		output = append(output, a.ToProto(todo))
	}

	return output
}

func (a *TodoAdapter) ToEntity(todo *model.Todo) *hyfdb.Todo {
	return &hyfdb.Todo{
		ID:               todo.ID,
		Title:            todo.Title,
		Description:      todo.Description,
		GroupID:          null.Int{},
		CreatorUserID:    null.Int{},
		Location:         null.String{},
		UrgencyLevel:     null.String{},
		CreatedAt:        time.Time{},
		CompletionStatus: null.Bool{},
	}
}

func (a *TodoAdapter) ToEntities(todos []*model.Todo) hyfdb.TodoSlice {
	entities := make(hyfdb.TodoSlice, 0, len(todos))

	for _, todo := range todos {
		entities = append(entities, a.ToEntity(todo))
	}

	return entities
}

func (a *TodoAdapter) FromEntity(todo *hyfdb.Todo) *model.Todo {
	return &model.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

func (a *TodoAdapter) FromEntities(todos hyfdb.TodoSlice) []*model.Todo {
	todosModel := make([]*model.Todo, 0, len(todos))

	for _, todo := range todos {
		todosModel = append(todosModel, a.FromEntity(todo))
	}

	return todosModel
}
