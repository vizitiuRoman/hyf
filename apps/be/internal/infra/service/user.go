package service

import (
	"context"
	"fmt"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/domain/repo"
	"github.com/vizitiuRoman/hyf/internal/domain/service"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type userService struct {
	logger log.Logger
	db     db.DB

	userAdapterFactory adapter.UserAdapterFactory
	userRepoFactory    repo.UserRepoFactory
}

func NewUserService(
	ctx context.Context,
	logger log.Logger,
	db db.DB,

	userAdapterFactory adapter.UserAdapterFactory,
	userRepoFactory repo.UserRepoFactory,
) service.UserService {
	return &userService{
		logger: logger.WithComponent(ctx, "user_service"),

		db: db,

		userAdapterFactory: userAdapterFactory,
		userRepoFactory:    userRepoFactory,
	}
}

func (s *userService) Find(ctx context.Context, id int64) (*model.User, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	user, err := userRepo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot find user: %w", err)
	}

	return user, nil
}

func (s *userService) FindAll(ctx context.Context) ([]*model.User, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	users, err := userRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot find users: %w", err)
	}

	return users, nil
}

func (s *userService) Create(ctx context.Context, input *pb.CreateUserIn) (*model.User, error) {
	userAdapter := s.userAdapterFactory.Create(ctx)
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	user, err := userRepo.Create(ctx, userAdapter.FromProto(input.User))
	if err != nil {
		return nil, fmt.Errorf("cannot create user: %w", err)
	}

	return user, nil
}

// Update --> @TODO --> return error when app is created with the same ID
func (s *userService) Update(ctx context.Context, input *pb.UpdateUserIn) (*model.User, error) {
	userAdapter := s.userAdapterFactory.Create(ctx)
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	user, err := userRepo.Update(ctx, userAdapter.FromProto(input.User))
	if err != nil {
		return nil, fmt.Errorf("cannot update user: %w", err)
	}

	return user, nil
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	userRepo := s.userRepoFactory.Create(ctx, s.db)

	err := userRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot delete user: %w", err)
	}

	return nil
}
