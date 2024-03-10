package cache

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/domain/repo"
	"go.uber.org/zap"
)

type authTokenRepoFactory struct {
	logger log.Logger

	authTokenAdapterFactory adapter.AuthTokenAdapterFactory
}

func NewAuthTokenRepoFactory(logger log.Logger, authTokenAdapterFactory adapter.AuthTokenAdapterFactory) repo.AuthTokenRepoFactory {
	return &authTokenRepoFactory{
		logger:                  logger,
		authTokenAdapterFactory: authTokenAdapterFactory,
	}
}

func (f *authTokenRepoFactory) Create(ctx context.Context, rdb redis.Cache) repo.AuthTokenRepo {
	return &authTokenRepo{
		logger: f.logger.WithComponent(ctx, "AuthTokenRepo"),
		rdb:    rdb,

		authTokenAdapterFactory: f.authTokenAdapterFactory,
	}
}

type authTokenRepo struct {
	logger log.Logger
	rdb    redis.Cache

	authTokenAdapterFactory adapter.AuthTokenAdapterFactory
}

func (r *authTokenRepo) keyID(u uuid.UUID) string {
	return "auth/token/{id:" + u.String() + "}/"
}

func (r *authTokenRepo) patternID(u uuid.UUID) string {
	return "auth/token/*{id:" + u.String() + "}/*"
}

func (r *authTokenRepo) Create(ctx context.Context, token *model.AuthToken) error {
	authTokenString, err := r.authTokenAdapterFactory.Create(ctx).MarshalJSON(ctx, token)
	if err != nil {
		return err
	}

	return r.rdb.Set(ctx, r.keyID(token.UUID), authTokenString)
}

func (r *authTokenRepo) Update(ctx context.Context, token *model.AuthToken) error {
	return r.Create(ctx, token)
}

func (r *authTokenRepo) Delete(ctx context.Context, u uuid.UUID) error {
	keys, err := r.rdb.KeysByPattern(ctx, r.patternID(u))
	if err != nil {
		return err
	}

	return r.rdb.Del(ctx, keys...)
}

func (r *authTokenRepo) GetByUUID(ctx context.Context, u uuid.UUID) (*model.AuthToken, error) {
	str, err := r.rdb.Get(ctx, r.keyID(u))
	if err != nil {
		return nil, err
	}

	if str == nil {
		r.logger.WithMethod(ctx, "GetByUUID").Error("auth token not found", zap.String("uuid", u.String()), zap.Error(domain.ErrNotFound))
		return nil, domain.ErrNotFound
	}

	return r.authTokenAdapterFactory.Create(ctx).UnmarshalJSON(ctx, str)
}
