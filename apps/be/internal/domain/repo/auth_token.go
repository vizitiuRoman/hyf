package repo

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
)

type AuthTokenRepoFactory interface {
	Create(context.Context, redis.Cache) AuthTokenRepo
}

type AuthTokenRepo interface {
	Create(context.Context, *model.AuthToken) error
	Update(context.Context, *model.AuthToken) error
	Delete(context.Context, uuid.UUID) error
	GetByID(context.Context, uuid.UUID) (*model.AuthToken, error)
}
