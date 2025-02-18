package infra

import (
	"github.com/vizitiuRoman/hyf/internal/infra/adapter"
	"github.com/vizitiuRoman/hyf/internal/infra/repo"
	"github.com/vizitiuRoman/hyf/internal/infra/service"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	repo.NewTodoRepoFactory,

	adapter.NewTodoAdapter,

	service.NewTodoService,
)
