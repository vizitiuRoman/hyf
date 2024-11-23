package service

import (
	"context"
	"fmt"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/infra/adapter"
	"github.com/vizitiuRoman/hyf/internal/infra/repo"
	"github.com/vizitiuRoman/hyf/pkg/adapter/logger"
	"github.com/vizitiuRoman/hyf/pkg/adapter/pgclient"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type UserService struct {
	logger logger.Logger
	db     pgclient.DB

	adapter         *adapter.UserAdapter
	userRepoFactory repo.UserRepoFactory
}

func NewUserService(
	ctx context.Context,
	logger logger.Logger,
	db pgclient.DB,

	adapter *adapter.UserAdapter,
	userRepoFactory repo.UserRepoFactory,
) *UserService {
	return &UserService{
		logger: logger.WithComponent(ctx, "user_service"),

		db: db,

		adapter:         adapter,
		userRepoFactory: userRepoFactory,
	}
}

func (s *UserService) Find(ctx context.Context, id int64) (*model.User, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	user, err := userRepo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot find user: %w", err)
	}

	return user, nil
}

func (s *UserService) FindAll(ctx context.Context) ([]*model.User, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	users, err := userRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot find users: %w", err)
	}

	return users, nil
}

func (s *UserService) Create(ctx context.Context, input *pb.CreateUserIn) (*model.User, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	user, err := userRepo.Create(ctx, s.adapter.FromProto(input.User))
	if err != nil {
		return nil, fmt.Errorf("cannot create user: %w", err)
	}

	return user, nil
}

// Update --> @TODO --> return error when app is created with the same ID
func (s *UserService) Update(ctx context.Context, input *pb.UpdateUserIn) (*model.User, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	user, err := userRepo.Update(ctx, s.adapter.FromProto(input.User))
	if err != nil {
		return nil, fmt.Errorf("cannot update user: %w", err)
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	err := userRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot delete user: %w", err)
	}

	return nil
}
