package infra

import (
	"github.com/vizitiuRoman/auth-service/internal/infra/adapter"
	"github.com/vizitiuRoman/auth-service/internal/infra/repo"
	"github.com/vizitiuRoman/auth-service/internal/infra/repo/cache"
	"github.com/vizitiuRoman/auth-service/internal/infra/service"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	adapter.NewAuthAdapter,
	adapter.NewUserAdapter,

	cache.NewAuthRepoFactory,
	repo.NewUserRepoFactory,

	service.NewAuthService,
	service.NewUserService,
)
