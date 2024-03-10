package repo

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
)

type UserRepoFactory interface {
	Create(context.Context, db.DB) UserRepo
}

type UserRepo interface {
	Find(context.Context, int64) (*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	Delete(context.Context, int64) error
}
