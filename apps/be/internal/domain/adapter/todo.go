package adapter

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"github.com/vizitiuRoman/hyf/pkg/gen/sqlboiler/hyfdb"
)

type TodoAdapterFactory interface {
	Create(context.Context) TodoAdapter
}

type TodoAdapter interface {
	FromProto(todo *pb.Todo) *model.Todo
	ToProto(todo *model.Todo) *pb.Todo
	ToProtos(todos []*model.Todo) []*pb.Todo

	ToEntity(*model.Todo) *hyfdb.Todo
	ToEntities([]*model.Todo) hyfdb.TodoSlice
	FromEntity(*hyfdb.Todo) *model.Todo
	FromEntities(hyfdb.TodoSlice) []*model.Todo
}
