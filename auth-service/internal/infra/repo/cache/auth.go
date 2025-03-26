package cache

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/vizitiuRoman/auth-service/internal/common/adapter/helper/util"
	"github.com/vizitiuRoman/auth-service/internal/domain"
	"github.com/vizitiuRoman/auth-service/internal/domain/model"
	"github.com/vizitiuRoman/auth-service/internal/infra/adapter"
	"github.com/vizitiuRoman/libs/cache/redis"
	"github.com/vizitiuRoman/libs/logger"
	"go.uber.org/zap"
)

type AuthRepoFactory struct {
	logger logger.Logger

	adapter *adapter.AuthAdapter
}

func NewAuthRepoFactory(logger logger.Logger, adapter *adapter.AuthAdapter) *AuthRepoFactory {
	return &AuthRepoFactory{
		logger:  logger,
		adapter: adapter,
	}
}

func (f *AuthRepoFactory) Create(ctx context.Context, rdb redis.Cache) *AuthRepo {
	return &AuthRepo{
		logger: f.logger.WithComponent(ctx, "AuthRepo"),
		rdb:    rdb,

		adapter: f.adapter,
	}
}

type AuthRepo struct {
	logger logger.Logger
	rdb    redis.Cache

	adapter *adapter.AuthAdapter
}

func (r *AuthRepo) keyID(u uuid.UUID) string {
	return "auth/token/{id:" + u.String() + "}/"
}

func (r *AuthRepo) patternID(u uuid.UUID) string {
	return "auth/token/*{id:" + u.String() + "}/*"
}

func (r *AuthRepo) Create(ctx context.Context, token *model.AuthToken) error {
	authTokenString, err := r.adapter.MarshalJSON(ctx, token)
	if err != nil {
		return err
	}

	return r.rdb.Set(r.keyID(token.UUID), []byte(util.UnP(authTokenString)))
}

func (r *AuthRepo) Update(ctx context.Context, token *model.AuthToken) error {
	return r.Create(ctx, token)
}

func (r *AuthRepo) Delete(ctx context.Context, u uuid.UUID) error {
	return r.rdb.Del(r.patternID(u))
}

func (r *AuthRepo) GetByUUID(ctx context.Context, u uuid.UUID) (*model.AuthToken, error) {
	str, err := r.rdb.Get(r.keyID(u))
	if err != nil {
		return nil, err
	}

	if str == nil {
		r.logger.WithMethod(ctx, "GetByUUID").Error("auth token not found", zap.String("uuid", u.String()), zap.Error(domain.ErrNotFound))
		return nil, domain.ErrNotFound
	}

	return r.adapter.UnmarshalJSON(ctx, str)
}
