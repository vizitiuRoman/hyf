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

	adapter.NewAuthTokenAdapterFactory,
	adapter.NewTodoAdapterFactory,

	service.NewAuthService,
	service.NewTodoService,
)
