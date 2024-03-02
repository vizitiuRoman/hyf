package repo

import (
	"context"

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

type todoRepoFactory struct {
	logger log.Logger

	todoAdapterFactory adapter.TodoAdapterFactory
}

func NewTodoRepoFactory(logger log.Logger, todoAdapterFactory adapter.TodoAdapterFactory) repo.TodoRepoFactory {
	return &todoRepoFactory{
		logger: logger,

		todoAdapterFactory: todoAdapterFactory,
	}
}

func (f *todoRepoFactory) Create(ctx context.Context, db db.DB) repo.TodoRepo {
	return &todoRepo{
		logger: f.logger.WithComponent(ctx, "TodoRepo"),
		db:     db,

		todoAdapter: f.todoAdapterFactory.Create(ctx),
	}
}

type todoRepo struct {
	logger log.Logger
	db     db.DB

	todoAdapter adapter.TodoAdapter
}

func (r *todoRepo) Find(ctx context.Context, id int64) (*model.Todo, error) {
	ent, err := hyfdb.Todos(hyfdb.TodoWhere.ID.EQ(int(id))).One(ctx, r.db)
	if err != nil {
		r.logger.WithMethod(ctx, "Find").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.todoAdapter.FromEntity(ent), nil
}

func (r *todoRepo) FindAll(ctx context.Context) ([]*model.Todo, error) {
	ent, err := hyfdb.Todos().All(ctx, r.db)
	if err != nil {
		r.logger.WithMethod(ctx, "FindAll").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.todoAdapter.FromEntities(ent), nil
}

func (r *todoRepo) Create(ctx context.Context, input *model.Todo) (*model.Todo, error) {
	ent := r.todoAdapter.ToEntity(input)

	err := ent.Insert(ctx, r.db, boil.Blacklist(hyfdb.TodoColumns.ID))
	if err != nil {
		r.logger.WithMethod(ctx, "Create").Error("execute query", zap.Error(err))
		return nil, err
	}

	return r.todoAdapter.FromEntity(ent), nil
}

func (r *todoRepo) Update(ctx context.Context, input *model.Todo) (*model.Todo, error) {
	ent := r.todoAdapter.ToEntity(input)

	rowsAff, err := ent.Update(ctx, r.db, boil.Infer())
	if err != nil {
		r.logger.WithMethod(ctx, "Update").Error("failed to update saga", zap.Error(err))
		return nil, err
	}
	if rowsAff != 1 {
		r.logger.WithMethod(ctx, "Update").Error("no rows affected", zap.Error(domain.ErrNotFound))
		return nil, domain.ErrNotFound
	}

	return r.todoAdapter.FromEntity(ent), nil
}

func (r *todoRepo) Delete(ctx context.Context, id int64) error {
	rowsAff, err := hyfdb.Todos(hyfdb.TodoWhere.ID.EQ(int(id))).DeleteAll(ctx, r.db)
	if err != nil {
		r.logger.WithMethod(ctx, "Delete").Error("failed to delete app", zap.Error(err))
		return err
	}
	if rowsAff == 0 {
		r.logger.WithMethod(ctx, "Delete").Error("failed to delete app", zap.Error(domain.ErrNotFound))
		return domain.ErrNotFound
	}

	return nil
}
