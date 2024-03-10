package adapter

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"github.com/vizitiuRoman/hyf/pkg/gen/sqlboiler/hyfdb"
)

type UserAdapterFactory interface {
	Create(context.Context) UserAdapter
}

type UserAdapter interface {
	FromProto(todo *pb.User) *model.User
	ToProto(todo *model.User) *pb.User
	ToProtos(todos []*model.User) []*pb.User

	ToEntity(*model.User) (*hyfdb.User, error)
	ToEntities([]*model.User) (hyfdb.UserSlice, error)
	FromEntity(*hyfdb.User) *model.User
	FromEntities(hyfdb.UserSlice) []*model.User
}
