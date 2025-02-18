package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vizitiuRoman/auth-service/internal/domain"
	"github.com/vizitiuRoman/auth-service/internal/domain/model"
	"github.com/vizitiuRoman/auth-service/internal/infra/adapter"
	authdb "github.com/vizitiuRoman/auth-service/pkg/gen/sqlboiler/authdb"
	"github.com/vizitiuRoman/libs/logger"
	"github.com/vizitiuRoman/libs/pgclient"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type UserRepoFactory struct {
	logger logger.Logger

	adapter *adapter.UserAdapter
}

func NewUserRepoFactory(logger logger.Logger, adapter *adapter.UserAdapter) *UserRepoFactory {
	return &UserRepoFactory{
		logger: logger,

		adapter: adapter,
	}
}

func (f *UserRepoFactory) Create(ctx context.Context, db pgclient.DB) *UserRepo {
	return &UserRepo{
		logger: f.logger.WithComponent(ctx, "UserRepo"),
		db:     db,

		adapter: f.adapter,
	}
}

type UserRepo struct {
	logger logger.Logger
	db     pgclient.DB

	adapter *adapter.UserAdapter
}

func (r *UserRepo) Find(ctx context.Context, id int64) (*model.User, error) {
	ent, err := authdb.Users(authdb.UserWhere.ID.EQ(id)).One(ctx, r.db)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		r.logger.WithMethod(ctx, "Find").Error("user not found", zap.Int64("id", id), zap.Error(err))
		return nil, domain.ErrNotFound

	case err != nil:
		r.logger.WithMethod(ctx, "Find").Error("failed to find user", zap.Int64("id", id), zap.Error(err))
		return nil, err

	default:
		return r.adapter.FromEntity(ent), nil
	}
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	ent, err := authdb.Users(authdb.UserWhere.Email.EQ(email)).One(ctx, r.db)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		r.logger.WithMethod(ctx, "Find").Error("user not found", zap.String("email", email), zap.Error(err))
		return nil, domain.ErrNotFound

	case err != nil:
		r.logger.WithMethod(ctx, "Find").Error("failed to find user", zap.String("email", email), zap.Error(err))
		return nil, err

	default:
		return r.adapter.FromEntity(ent), nil
	}
}

func (r *UserRepo) FindAll(ctx context.Context) ([]*model.User, error) {
	ent, err := authdb.Users().All(ctx, r.db)
	if err != nil {
		r.logger.WithMethod(ctx, "FindAll").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.adapter.FromEntities(ent), nil
}

func (r *UserRepo) Create(ctx context.Context, input *model.User) (*model.User, error) {
	ent, err := r.adapter.ToEntity(input)
	if err != nil {
		r.logger.WithMethod(ctx, "Create").Error("failed to adapt entity", zap.Error(err))
		return nil, err
	}

	err = ent.Insert(ctx, r.db, boil.Blacklist(authdb.UserColumns.ID))
	if err != nil {
		r.logger.WithMethod(ctx, "Create").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.adapter.FromEntity(ent), nil
}

func (r *UserRepo) Update(ctx context.Context, input *model.User) (*model.User, error) {
	ent, err := r.adapter.ToEntity(input)
	if err != nil {
		r.logger.WithMethod(ctx, "Update").Error("failed to adapt entity", zap.Error(err))
		return nil, err
	}

	rowsAff, err := ent.Update(ctx, r.db, boil.Infer())
	if err != nil {
		r.logger.WithMethod(ctx, "Update").Error("failed to update user", zap.Error(err))
		return nil, err
	}
	if rowsAff != 1 {
		r.logger.WithMethod(ctx, "Update").Error("no rows affected", zap.Error(domain.ErrNotFound))
		return nil, domain.ErrNotFound
	}

	return r.adapter.FromEntity(ent), nil
}

func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	rowsAff, err := authdb.Users(authdb.UserWhere.ID.EQ(id)).DeleteAll(ctx, r.db)
	if err != nil {
		r.logger.WithMethod(ctx, "Delete").Error("failed to delete user", zap.Error(err))
		return err
	}
	if rowsAff == 0 {
		r.logger.WithMethod(ctx, "Delete").Error("failed to delete user", zap.Error(domain.ErrNotFound))
		return domain.ErrNotFound
	}

	return nil
}
