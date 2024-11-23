package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vizitiuRoman/hyf/internal/domain"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	"github.com/vizitiuRoman/hyf/internal/infra/adapter"
	"github.com/vizitiuRoman/hyf/pkg/adapter/logger"
	"github.com/vizitiuRoman/hyf/pkg/adapter/pgclient"
	"github.com/vizitiuRoman/hyf/pkg/gen/sqlboiler/hyfdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type TodoRepoFactory struct {
	logger logger.Logger

	adapter *adapter.TodoAdapter
}

func NewTodoRepoFactory(logger logger.Logger, adapter *adapter.TodoAdapter) *TodoRepoFactory {
	return &TodoRepoFactory{
		logger: logger,

		adapter: adapter,
	}
}

func (f *TodoRepoFactory) Create(ctx context.Context, db pgclient.DB) *TodoRepo {
	return &TodoRepo{
		logger: f.logger.WithComponent(ctx, "TodoRepo"),
		db:     db,

		adapter: f.adapter,
	}
}

type TodoRepo struct {
	logger logger.Logger
	db     pgclient.DB

	adapter *adapter.TodoAdapter
}

func (r *TodoRepo) Find(ctx context.Context, id int64) (*model.Todo, error) {
	ent, err := hyfdb.Todos(hyfdb.TodoWhere.ID.EQ(id)).One(ctx, r.db)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		r.logger.WithMethod(ctx, "Find").Error("todo not found", zap.Int64("id", id), zap.Error(err))
		return nil, domain.ErrNotFound

	case err != nil:
		r.logger.WithMethod(ctx, "Find").Error("failed to find todo", zap.Int64("id", id), zap.Error(err))
		return nil, err

	default:
		return r.adapter.FromEntity(ent), nil
	}
}

func (r *TodoRepo) FindAll(ctx context.Context) ([]*model.Todo, error) {
	ent, err := hyfdb.Todos().All(ctx, r.db)
	if err != nil {
		r.logger.WithMethod(ctx, "FindAll").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.adapter.FromEntities(ent), nil
}

func (r *TodoRepo) Create(ctx context.Context, input *model.Todo) (*model.Todo, error) {
	ent := r.adapter.ToEntity(input)

	err := ent.Insert(ctx, r.db, boil.Blacklist(hyfdb.TodoColumns.ID))
	if err != nil {
		r.logger.WithMethod(ctx, "Create").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.adapter.FromEntity(ent), nil
}

func (r *TodoRepo) Update(ctx context.Context, input *model.Todo) (*model.Todo, error) {
	ent := r.adapter.ToEntity(input)

	rowsAff, err := ent.Update(ctx, r.db, boil.Infer())
	if err != nil {
		r.logger.WithMethod(ctx, "Update").Error("failed to update todo", zap.Error(err))
		return nil, err
	}
	if rowsAff != 1 {
		r.logger.WithMethod(ctx, "Update").Error("no rows affected", zap.Error(domain.ErrNotFound))
		return nil, domain.ErrNotFound
	}

	return r.adapter.FromEntity(ent), nil
}

func (r *TodoRepo) Delete(ctx context.Context, id int64) error {
	rowsAff, err := hyfdb.Todos(hyfdb.TodoWhere.ID.EQ(id)).DeleteAll(ctx, r.db)
	if err != nil {
		r.logger.WithMethod(ctx, "Delete").Error("failed to delete todo", zap.Error(err))
		return err
	}
	if rowsAff == 0 {
		r.logger.WithMethod(ctx, "Delete").Error("failed to delete todo", zap.Error(domain.ErrNotFound))
		return domain.ErrNotFound
	}

	return nil
}
