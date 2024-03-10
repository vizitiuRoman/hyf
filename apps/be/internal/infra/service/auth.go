package service

import (
	"context"
	"fmt"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/domain/repo"
	"github.com/vizitiuRoman/hyf/internal/domain/service"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	logger log.Logger

	db  db.DB
	rdb redis.Cache

	authAdapterFactory adapter.AuthTokenAdapterFactory

	authTokenRepoFactory repo.AuthTokenRepoFactory
	userRepoFactory      repo.UserRepoFactory
}

func NewAuthService(
	ctx context.Context,

	logger log.Logger,

	db db.DB,
	rdb redis.Cache,

	authAdapterFactory adapter.AuthTokenAdapterFactory,

	authTokenRepoFactory repo.AuthTokenRepoFactory,
	userRepoFactory repo.UserRepoFactory,
) service.AuthService {
	return &authService{
		logger: logger.WithComponent(ctx, "auth_service"),

		db:  db,
		rdb: rdb,

		authAdapterFactory: authAdapterFactory,

		authTokenRepoFactory: authTokenRepoFactory,
		userRepoFactory:      userRepoFactory,
	}
}

func (s *authService) Login(ctx context.Context, in *pb.LoginIn) (*model.AuthToken, error) {
	logger := s.logger.WithMethod(ctx, "Login").With(
		zap.String("email", in.Email),
	)

	userRepo := s.userRepoFactory.Create(ctx, s.db)
	user, err := userRepo.FindByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		logger.Error("cannot compare hash and password", zap.Error(err))
		return nil, err
	}

	authAdapter := s.authAdapterFactory.Create(ctx)
	authToken, err := authAdapter.FromUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	authTokenRepo := s.authTokenRepoFactory.Create(ctx, s.rdb)
	err = authTokenRepo.Create(ctx, authToken)
	if err != nil {
		logger.Error("cannot login", zap.Error(err))
		return nil, fmt.Errorf("cannot login: %w", err)
	}

	return authToken, nil
}

func (s *authService) Register(ctx context.Context, in *pb.RegisterIn) (*model.AuthToken, error) {
	authAdapter := s.authAdapterFactory.Create(ctx)

	userRepo := s.userRepoFactory.Create(ctx, s.db)
	user, err := userRepo.Create(ctx, authAdapter.FromRegisterProtoToUser(in))
	if err != nil {
		return nil, err
	}

	authToken, err := authAdapter.FromUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	authTokenRepo := s.authTokenRepoFactory.Create(ctx, s.rdb)

	err = authTokenRepo.Create(ctx, authToken)
	if err != nil {
		s.logger.WithMethod(ctx, "Register").Error("cannot register", zap.Error(err))
		return nil, fmt.Errorf("cannot register: %w", err)
	}

	return authToken, nil
}
