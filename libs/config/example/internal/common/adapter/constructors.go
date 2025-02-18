package adapter

import (
	"go.uber.org/fx"

	"github.com/vizitiuRoman/libs/config/example/internal/common/adapter/config"
)

var Constructors = fx.Provide(
	config.NewConfig,
)
