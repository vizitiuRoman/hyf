package service

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/vizitiuRoman/auth-service/internal/domain/model"
	"github.com/vizitiuRoman/auth-service/internal/infra/adapter"
	"github.com/vizitiuRoman/auth-service/internal/infra/repo"
	"github.com/vizitiuRoman/auth-service/internal/infra/repo/cache"
	pb "github.com/vizitiuRoman/auth-service/pkg/gen/auth/v1"
	"github.com/vizitiuRoman/libs/cache/redis"
	"github.com/vizitiuRoman/libs/logger"
	"github.com/vizitiuRoman/libs/pgclient"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	logger logger.Logger

	db  pgclient.DB
	rdb redis.Cache

	adapter *adapter.AuthAdapter

	authTokenRepoFactory *cache.AuthRepoFactory
	userRepoFactory      *repo.UserRepoFactory
}

func NewAuthService(
	ctx context.Context,

	logger logger.Logger,

	db pgclient.DB,
	rdb redis.Cache,

	adapter *adapter.AuthAdapter,

	authTokenRepoFactory *cache.AuthRepoFactory,
	userRepoFactory *repo.UserRepoFactory,
) *AuthService {
	return &AuthService{
		logger: logger.WithComponent(ctx, "auth_service"),

		db:  db,
		rdb: rdb,

		adapter: adapter,

		authTokenRepoFactory: authTokenRepoFactory,
		userRepoFactory:      userRepoFactory,
	}
}

func (s *AuthService) Login(ctx context.Context, in *pb.LoginIn) (*model.AuthToken, *model.AuthToken, error) {
	logger := s.logger.WithMethod(ctx, "Login").With(
		zap.String("email", in.Email),
	)

	userRepo := s.userRepoFactory.Create(ctx, s.db)
	user, err := userRepo.FindByEmail(ctx, in.Email)
	if err != nil {
		return nil, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		logger.Error("cannot compare hash and password", zap.Error(err))
		return nil, nil, err
	}

	authToken, refreshToken, err := s.adapter.FromUserID(ctx, user.ID)
	if err != nil {
		return nil, nil, err
	}

	authTokenRepo := s.authTokenRepoFactory.Create(ctx, s.rdb)
	err = authTokenRepo.Create(ctx, authToken)
	if err != nil {
		logger.Error("cannot save access token", zap.Error(err))
		return nil, nil, fmt.Errorf("cannot login: %w", err)
	}

	err = authTokenRepo.Create(ctx, refreshToken)
	if err != nil {
		logger.Error("cannot save refresh token", zap.Error(err))
		return nil, nil, fmt.Errorf("cannot login: %w", err)
	}

	return authToken, refreshToken, nil
}

func (s *AuthService) Register(ctx context.Context, in *pb.RegisterIn) (*model.AuthToken, *model.AuthToken, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.db)
	user, err := userRepo.Create(ctx, s.adapter.FromRegisterProtoToUser(in))
	if err != nil {
		return nil, nil, err
	}

	authToken, refreshToken, err := s.adapter.FromUserID(ctx, user.ID)
	if err != nil {
		return nil, nil, err
	}

	authTokenRepo := s.authTokenRepoFactory.Create(ctx, s.rdb)
	err = authTokenRepo.Create(ctx, authToken)
	if err != nil {
		s.logger.WithMethod(ctx, "Register").Error("cannot save access token", zap.Error(err))
		return nil, nil, fmt.Errorf("cannot register: %w", err)
	}

	err = authTokenRepo.Create(ctx, refreshToken)
	if err != nil {
		s.logger.WithMethod(ctx, "Register").Error("cannot save refresh token", zap.Error(err))
		return nil, nil, fmt.Errorf("cannot register: %w", err)
	}

	return authToken, refreshToken, nil
}

func (s *AuthService) Refresh(ctx context.Context, in *pb.RefreshIn) (*model.AuthToken, *model.AuthToken, error) {
	logger := s.logger.WithMethod(ctx, "Refresh").With(zap.String("refresh_token", in.RefreshToken))

	authTokenRepo := s.authTokenRepoFactory.Create(ctx, s.rdb)

	uuidClaim, userID, err := s.adapter.ClaimsFromRefreshJWT(ctx, in.RefreshToken)
	if err != nil {
		logger.Error("cannot claim from refresh jwt", zap.Error(err))
		return nil, nil, err
	}

	_, err = authTokenRepo.GetByUUID(ctx, uuid.FromStringOrNil(uuidClaim))
	if err != nil {
		logger.Error("cannot find refresh token by uuid", zap.Error(err), zap.String("uuid", uuidClaim))
		return nil, nil, err
	}

	authToken, refreshToken, err := s.adapter.FromUserID(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	err = authTokenRepo.Create(ctx, authToken)
	if err != nil {
		logger.Error("cannot save access token", zap.Error(err))
		return nil, nil, fmt.Errorf("cannot refresh: %w", err)
	}

	err = authTokenRepo.Create(ctx, refreshToken)
	if err != nil {
		logger.Error("cannot save refresh token", zap.Error(err))
		return nil, nil, fmt.Errorf("cannot refresh: %w", err)
	}

	return authToken, refreshToken, nil
}

func (s *AuthService) Logout(ctx context.Context, in *pb.LogoutIn) error {
	logger := s.logger.WithMethod(ctx, "Logout").With(zap.String("refresh_token", in.RefreshToken), zap.String("access_token", in.Token))

	authTokenRepo := s.authTokenRepoFactory.Create(ctx, s.rdb)

	accessUUID, _, err := s.adapter.ClaimsFromJWT(ctx, in.Token)
	if err != nil {
		logger.Error("cannot claim from jwt", zap.Error(err))
		return err
	}

	refreshUUID, _, err := s.adapter.ClaimsFromRefreshJWT(ctx, in.RefreshToken)
	if err != nil {
		logger.Error("cannot claim from refresh jwt", zap.Error(err))
		return err
	}

	err = authTokenRepo.Delete(ctx, uuid.FromStringOrNil(refreshUUID))
	if err != nil {
		logger.Error("cannot delete refresh token", zap.Error(err))
		return fmt.Errorf("cannot logout: %w", err)
	}

	err = authTokenRepo.Delete(ctx, uuid.FromStringOrNil(accessUUID))
	if err != nil {
		logger.Error("cannot delete access token", zap.Error(err))
		return fmt.Errorf("cannot logout: %w", err)
	}

	return nil
}
