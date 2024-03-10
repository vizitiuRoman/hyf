package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"github.com/vizitiuRoman/hyf/pkg/gen/sqlboiler/hyfdb"
	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"
)

type userAdapterFactory struct {
	logger log.Logger
}

func NewUserAdapterFactory(logger log.Logger) adapter.UserAdapterFactory {
	return &userAdapterFactory{
		logger: logger,
	}
}

func (f *userAdapterFactory) Create(ctx context.Context) adapter.UserAdapter {
	return &userAdapter{
		logger: f.logger.WithComponent(ctx, "UserAdapter"),
	}
}

type userAdapter struct {
	logger log.Logger
}

func (a *userAdapter) FromProto(user *pb.User) *model.User {
	return &model.User{
		ID:       user.Id,
		Email:    user.Name,
		Password: user.Password,
		Name:     user.Name,
		LastName: user.LastName,
	}
}

func (a *userAdapter) ToProto(user *model.User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
}

func (a *userAdapter) ToProtos(users []*model.User) []*pb.User {
	output := make([]*pb.User, 0, len(users))

	for _, user := range users {
		output = append(output, a.ToProto(user))
	}

	return output
}

func (a *userAdapter) ToEntity(user *model.User) (*hyfdb.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot generate password hash: %w", err)
	}

	return &hyfdb.User{
		ID:          user.ID,
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    string(hashedPassword),
		CreatedAt:   time.Now(),
		LastLoginAt: null.Time{},
	}, nil
}

func (a *userAdapter) ToEntities(users []*model.User) (hyfdb.UserSlice, error) {
	entities := make(hyfdb.UserSlice, 0, len(users))

	for _, user := range users {
		u, err := a.ToEntity(user)
		if err != nil {
			return nil, err
		}

		entities = append(entities, u)
	}

	return entities, nil
}

func (a *userAdapter) FromEntity(user *hyfdb.User) *model.User {
	return &model.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		LastName: user.LastName,
	}
}

func (a *userAdapter) FromEntities(users hyfdb.UserSlice) []*model.User {
	usersModel := make([]*model.User, 0, len(users))

	for _, user := range users {
		usersModel = append(usersModel, a.FromEntity(user))
	}

	return usersModel
}
