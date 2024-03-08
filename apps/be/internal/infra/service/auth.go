package service

import (
	"context"
	"fmt"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/domain/repo"
	"github.com/vizitiuRoman/hyf/internal/domain/service"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type authService struct {
	logger log.Logger

	rdb redis.Cache

	authAdapterFactory adapter.AuthTokenAdapterFactory
	authRepoFactory    repo.AuthTokenRepoFactory
}

func NewAuthService(
	ctx context.Context,

	logger log.Logger,

	rdb redis.Cache,

	authAdapterFactory adapter.AuthTokenAdapterFactory,
	authRepoFactory repo.AuthTokenRepoFactory,
) service.AuthService {
	return &authService{
		logger: logger.WithComponent(ctx, "auth_service"),
		rdb:    rdb,

		authAdapterFactory: authAdapterFactory,
		authRepoFactory:    authRepoFactory,
	}
}

func (s *authService) Login(ctx context.Context, in *pb.LoginIn) (*model.AuthToken, error) {

	userID := int64(1)
	// verify user

	authAdapter := s.authAdapterFactory.Create(ctx)
	authToken, err := authAdapter.FromLoginProtoAndUserID(in, userID)
	if err != nil {
		return nil, err
	}

	authTokenRepo := s.authRepoFactory.Create(ctx, s.rdb)
	err = authTokenRepo.Create(ctx, authToken)
	if err != nil {
		return nil, fmt.Errorf("cannot login: %w", err)
	}

	return authToken, nil
}

func (s *authService) Register(ctx context.Context, in *pb.RegisterIn) (*model.AuthToken, error) {
	//TODO implement me
	panic("implement me")
}
