package config

import (
	"github.com/vizitiuRoman/auth-service/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/libs/cache/redis"
	"github.com/vizitiuRoman/libs/config"
	"github.com/vizitiuRoman/libs/logger"
	"github.com/vizitiuRoman/libs/pgclient"
)

type Config struct {
	Location string `yaml:"location"`

	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	BuildDate string `yaml:"build_date"`

	Cache  *redis.Config    `yaml:"cache"`
	Logger *logger.Config   `yaml:"logger"`
	DB     *pgclient.Config `yaml:"db"`
	Server *grpc.Config     `yaml:"server"`

	Auth *struct {
		AccessTokenSecretKey  string `yaml:"access_token_secret_key"`
		RefreshTokenSecretKey string `yaml:"refresh_token_secret_key"`
	} `yaml:"auth"`
}

var cfg = &Config{}

func NewConfig() (*Config, error) {
	return cfg, config.PopulateConfig(cfg)
}
