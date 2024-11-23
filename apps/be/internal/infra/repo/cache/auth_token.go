package cache

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"
	"github.com/vizitiuRoman/hyf/internal/domain"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/infra/adapter"
	"github.com/vizitiuRoman/hyf/pkg/adapter/logger"
	"go.uber.org/zap"
)

type AuthTokenRepoFactory struct {
	logger logger.Logger

	adapter *adapter.AuthTokenAdapter
}

func NewAuthTokenRepoFactory(logger logger.Logger, adapter *adapter.AuthTokenAdapter) *AuthTokenRepoFactory {
	return &AuthTokenRepoFactory{
		logger:  logger,
		adapter: adapter,
	}
}

func (f *AuthTokenRepoFactory) Create(ctx context.Context, rdb redis.Cache) *AuthTokenRepo {
	return &AuthTokenRepo{
		logger: f.logger.WithComponent(ctx, "AuthTokenRepo"),
		rdb:    rdb,

		adapter: f.adapter,
	}
}

type AuthTokenRepo struct {
	logger logger.Logger
	rdb    redis.Cache

	adapter *adapter.AuthTokenAdapter
}

func (r *AuthTokenRepo) keyID(u uuid.UUID) string {
	return "auth/token/{id:" + u.String() + "}/"
}

func (r *AuthTokenRepo) patternID(u uuid.UUID) string {
	return "auth/token/*{id:" + u.String() + "}/*"
}

func (r *AuthTokenRepo) Create(ctx context.Context, token *model.AuthToken) error {
	authTokenString, err := r.adapter.MarshalJSON(ctx, token)
	if err != nil {
		return err
	}

	return r.rdb.Set(ctx, r.keyID(token.UUID), authTokenString)
}

func (r *AuthTokenRepo) Update(ctx context.Context, token *model.AuthToken) error {
	return r.Create(ctx, token)
}

func (r *AuthTokenRepo) Delete(ctx context.Context, u uuid.UUID) error {
	keys, err := r.rdb.KeysByPattern(ctx, r.patternID(u))
	if err != nil {
		return err
	}

	return r.rdb.Del(ctx, keys...)
}

func (r *AuthTokenRepo) GetByUUID(ctx context.Context, u uuid.UUID) (*model.AuthToken, error) {
	str, err := r.rdb.Get(ctx, r.keyID(u))
	if err != nil {
		return nil, err
	}

	if str == nil {
		r.logger.WithMethod(ctx, "GetByUUID").Error("auth token not found", zap.String("uuid", u.String()), zap.Error(domain.ErrNotFound))
		return nil, domain.ErrNotFound
	}

	return r.adapter.UnmarshalJSON(ctx, str)
}
