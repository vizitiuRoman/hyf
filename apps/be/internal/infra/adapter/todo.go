package adapter

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"github.com/vizitiuRoman/hyf/pkg/gen/sqlboiler/hyfdb"
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

func (t *todoAdapter) FromProto(todo *pb.Todo) *model.Todo {
	return &model.Todo{}
}

func (t *todoAdapter) ToProto(todo *model.Todo) *pb.Todo {
	return &pb.Todo{}
}

func (t *todoAdapter) ToProtos(todos []*model.Todo) []*pb.Todo {
	output := make([]*pb.Todo, 0, len(todos))

	for _, todo := range todos {
		output = append(output, &pb.Todo{
			Id:          todo.ID,
			Name:        "",
			Description: "",
		})
	}

	return output
}

func (t *todoAdapter) ToEntity(todo *model.Todo) *hyfdb.Todo {
	return &hyfdb.Todo{}
}

func (t *todoAdapter) ToEntities(todos []*model.Todo) hyfdb.TodoSlice {
	entities := make(hyfdb.TodoSlice, 0, len(todos))

	for _, todo := range todos {
		entities = append(entities, t.ToEntity(todo))
	}

	return entities
}

func (t *todoAdapter) FromEntity(todo *hyfdb.Todo) *model.Todo {
	return &model.Todo{}
}

func (t *todoAdapter) FromEntities(todos hyfdb.TodoSlice) []*model.Todo {
	todosModel := make([]*model.Todo, 0, len(todos))

	for _, todo := range todos {
		todosModel = append(todosModel, t.FromEntity(todo))
	}

	return todosModel
}
