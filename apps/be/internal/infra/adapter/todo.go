package adapter

import (
	"context"
	"time"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"github.com/vizitiuRoman/hyf/pkg/gen/sqlboiler/hyfdb"
	"github.com/volatiletech/null/v8"
)

type todoAdapterFactory struct {
	logger log.Logger
}

func NewTodoAdapterFactory(logger log.Logger) adapter.TodoAdapterFactory {
	return &todoAdapterFactory{
		logger: logger,
	}
}

func (f *todoAdapterFactory) Create(ctx context.Context) adapter.TodoAdapter {
	return &todoAdapter{
		logger: f.logger.WithComponent(ctx, "TodoAdapter"),
	}
}

type todoAdapter struct {
	logger log.Logger
}

func (a *todoAdapter) FromProto(todo *pb.Todo) *model.Todo {
	return &model.Todo{
		ID:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

func (a *todoAdapter) ToProto(todo *model.Todo) *pb.Todo {
	return &pb.Todo{
		Id:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

func (a *todoAdapter) ToProtos(todos []*model.Todo) []*pb.Todo {
	output := make([]*pb.Todo, 0, len(todos))

	for _, todo := range todos {
		output = append(output, a.ToProto(todo))
	}

	return output
}

func (a *todoAdapter) ToEntity(todo *model.Todo) *hyfdb.Todo {
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

func (a *todoAdapter) ToEntities(todos []*model.Todo) hyfdb.TodoSlice {
	entities := make(hyfdb.TodoSlice, 0, len(todos))

	for _, todo := range todos {
		entities = append(entities, a.ToEntity(todo))
	}

	return entities
}

func (a *todoAdapter) FromEntity(todo *hyfdb.Todo) *model.Todo {
	return &model.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

func (a *todoAdapter) FromEntities(todos hyfdb.TodoSlice) []*model.Todo {
	todosModel := make([]*model.Todo, 0, len(todos))

	for _, todo := range todos {
		todosModel = append(todosModel, a.FromEntity(todo))
	}

	return todosModel
}
