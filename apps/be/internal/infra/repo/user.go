package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/domain/repo"
	"github.com/vizitiuRoman/hyf/pkg/gen/sqlboiler/hyfdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type userRepoFactory struct {
	logger log.Logger

	userAdapterFactory adapter.UserAdapterFactory
}

func NewUserRepoFactory(logger log.Logger, userAdapterFactory adapter.UserAdapterFactory) repo.UserRepoFactory {
	return &userRepoFactory{
		logger: logger,

		userAdapterFactory: userAdapterFactory,
	}
}

func (f *userRepoFactory) Create(ctx context.Context, db db.DB) repo.UserRepo {
	return &userRepo{
		logger: f.logger.WithComponent(ctx, "UserRepo"),
		db:     db,

		userAdapter: f.userAdapterFactory.Create(ctx),
	}
}

type userRepo struct {
	logger log.Logger
	db     db.DB

	userAdapter adapter.UserAdapter
}

func (r *userRepo) Find(ctx context.Context, id int64) (*model.User, error) {
	ent, err := hyfdb.Users(hyfdb.UserWhere.ID.EQ(id)).One(ctx, r.db)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		r.logger.WithMethod(ctx, "Find").Error("user not found", zap.Int64("id", id), zap.Error(err))
		return nil, domain.ErrNotFound

	case err != nil:
		r.logger.WithMethod(ctx, "Find").Error("failed to find user", zap.Int64("id", id), zap.Error(err))
		return nil, err

	default:
		return r.userAdapter.FromEntity(ent), nil
	}
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	ent, err := hyfdb.Users(hyfdb.UserWhere.Email.EQ(email)).One(ctx, r.db)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		r.logger.WithMethod(ctx, "Find").Error("user not found", zap.String("email", email), zap.Error(err))
		return nil, domain.ErrNotFound

	case err != nil:
		r.logger.WithMethod(ctx, "Find").Error("failed to find user", zap.String("email", email), zap.Error(err))
		return nil, err

	default:
		return r.userAdapter.FromEntity(ent), nil
	}
}

func (r *userRepo) FindAll(ctx context.Context) ([]*model.User, error) {
	ent, err := hyfdb.Users().All(ctx, r.db)
	if err != nil {
		r.logger.WithMethod(ctx, "FindAll").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.userAdapter.FromEntities(ent), nil
}

func (r *userRepo) Create(ctx context.Context, input *model.User) (*model.User, error) {
	ent, err := r.userAdapter.ToEntity(input)
	if err != nil {
		r.logger.WithMethod(ctx, "Create").Error("failed to adapt entity", zap.Error(err))
		return nil, err
	}

	err = ent.Insert(ctx, r.db, boil.Blacklist(hyfdb.UserColumns.ID))
	if err != nil {
		r.logger.WithMethod(ctx, "Create").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.userAdapter.FromEntity(ent), nil
}

func (r *userRepo) Update(ctx context.Context, input *model.User) (*model.User, error) {
	ent, err := r.userAdapter.ToEntity(input)
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

	return r.userAdapter.FromEntity(ent), nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	rowsAff, err := hyfdb.Users(hyfdb.UserWhere.ID.EQ(id)).DeleteAll(ctx, r.db)
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
