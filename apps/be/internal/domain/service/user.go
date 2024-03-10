package service

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type UserService interface {
	Find(context.Context, int64) (*model.User, error)
	FindAll(context.Context) ([]*model.User, error)
	Create(context.Context, *pb.CreateUserIn) (*model.User, error)
	Update(context.Context, *pb.UpdateUserIn) (*model.User, error)
	Delete(context.Context, int64) error
}
