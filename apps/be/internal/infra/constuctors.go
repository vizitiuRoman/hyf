package infra

import (
	"github.com/vizitiuRoman/hyf/internal/infra/adapter"
	"github.com/vizitiuRoman/hyf/internal/infra/repo"
	"github.com/vizitiuRoman/hyf/internal/infra/repo/cache"
	"github.com/vizitiuRoman/hyf/internal/infra/service"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	cache.NewAuthTokenRepoFactory,

	repo.NewTodoRepoFactory,
	repo.NewUserRepoFactory,

	adapter.NewAuthTokenAdapterFactory,
	adapter.NewTodoAdapterFactory,
	adapter.NewUserAdapterFactory,

	service.NewAuthService,
	service.NewTodoService,
	service.NewUserService,
)
