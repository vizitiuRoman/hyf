package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/vizitiuRoman/auth-service/internal/domain/model"
	pb "github.com/vizitiuRoman/auth-service/pkg/gen/auth/v1"
	authdb "github.com/vizitiuRoman/auth-service/pkg/gen/sqlboiler/authdb"
	"github.com/vizitiuRoman/libs/logger"
	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"
)

type UserAdapter struct {
	logger logger.Logger
}

func NewUserAdapter(ctx context.Context, logger logger.Logger) *UserAdapter {
	return &UserAdapter{
		logger: logger.WithComponent(ctx, "UserAdapter"),
	}
}

func (a *UserAdapter) FromProto(user *pb.User) *model.User {
	return &model.User{
		ID:       user.Id,
		Email:    user.Name,
		Password: user.Password,
		Name:     user.Name,
		LastName: user.LastName,
	}
}

func (a *UserAdapter) ToProto(user *model.User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
}

func (a *UserAdapter) ToProtos(users []*model.User) []*pb.User {
	output := make([]*pb.User, 0, len(users))

	for _, user := range users {
		output = append(output, a.ToProto(user))
	}

	return output
}

func (a *UserAdapter) ToEntity(user *model.User) (*authdb.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot generate password hash: %w", err)
	}

	return &authdb.User{
		ID:          user.ID,
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    string(hashedPassword),
		CreatedAt:   time.Now(),
		LastLoginAt: null.Time{},
	}, nil
}

func (a *UserAdapter) ToEntities(users []*model.User) (authdb.UserSlice, error) {
	entities := make(authdb.UserSlice, 0, len(users))

	for _, user := range users {
		u, err := a.ToEntity(user)
		if err != nil {
			return nil, err
		}

		entities = append(entities, u)
	}

	return entities, nil
}

func (a *UserAdapter) FromEntity(user *authdb.User) *model.User {
	return &model.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		LastName: user.LastName,
	}
}

func (a *UserAdapter) FromEntities(users authdb.UserSlice) []*model.User {
	usersModel := make([]*model.User, 0, len(users))

	for _, user := range users {
		usersModel = append(usersModel, a.FromEntity(user))
	}

	return usersModel
}
