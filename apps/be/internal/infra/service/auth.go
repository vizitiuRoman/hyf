package service

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
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

func (s *authService) Login(ctx context.Context, in *pb.LoginIn) (*model.AuthToken, *model.AuthToken, error) {
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

	authAdapter := s.authAdapterFactory.Create(ctx)
	authToken, refreshToken, err := authAdapter.FromUserID(ctx, user.ID)
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

func (s *authService) Register(ctx context.Context, in *pb.RegisterIn) (*model.AuthToken, *model.AuthToken, error) {
	authAdapter := s.authAdapterFactory.Create(ctx)

	userRepo := s.userRepoFactory.Create(ctx, s.db)
	user, err := userRepo.Create(ctx, authAdapter.FromRegisterProtoToUser(in))
	if err != nil {
		return nil, nil, err
	}

	authToken, refreshToken, err := authAdapter.FromUserID(ctx, user.ID)
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

func (s *authService) Refresh(ctx context.Context, in *pb.RefreshIn) (*model.AuthToken, *model.AuthToken, error) {
	logger := s.logger.WithMethod(ctx, "Refresh").With(zap.String("refresh_token", in.RefreshToken))

	authTokenRepo := s.authTokenRepoFactory.Create(ctx, s.rdb)
	authAdapter := s.authAdapterFactory.Create(ctx)

	uuidClaim, userID, err := authAdapter.ClaimsFromRefreshJWT(ctx, in.RefreshToken)
	if err != nil {
		logger.Error("cannot claim from refresh jwt", zap.Error(err))
		return nil, nil, err
	}

	_, err = authTokenRepo.GetByUUID(ctx, uuid.FromStringOrNil(uuidClaim))
	if err != nil {
		logger.Error("cannot find refresh token by uuid", zap.Error(err), zap.String("uuid", uuidClaim))
		return nil, nil, err
	}

	authToken, refreshToken, err := authAdapter.FromUserID(ctx, userID)
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

func (s *authService) Logout(ctx context.Context, in *pb.LogoutIn) error {
	logger := s.logger.WithMethod(ctx, "Logout").With(zap.String("refresh_token", in.RefreshToken), zap.String("access_token", in.Token))

	authTokenRepo := s.authTokenRepoFactory.Create(ctx, s.rdb)
	authTokenAdapter := s.authAdapterFactory.Create(ctx)

	accessUUID, _, err := authTokenAdapter.ClaimsFromJWT(ctx, in.Token)
	if err != nil {
		logger.Error("cannot claim from jwt", zap.Error(err))
		return err
	}

	refreshUUID, _, err := authTokenAdapter.ClaimsFromRefreshJWT(ctx, in.RefreshToken)
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
