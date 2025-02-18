package config

import (
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/hyf/pkg/adapter/config"
	"github.com/vizitiuRoman/hyf/pkg/adapter/logger"
	"github.com/vizitiuRoman/hyf/pkg/adapter/pgclient"
)

type Config struct {
	Location string `yaml:"location"`

	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	BuildDate string `yaml:"build_date"`

	Logger *logger.Config   `yaml:"logger"`
	DB     *pgclient.Config `yaml:"db"`
	Server *grpc.Config     `yaml:"server"`

	Auth *struct {
		AccessTokenSecretKey  string `yaml:"access_token_secret_key"`
		RefreshTokenSecretKey string `yaml:"refresh_token_secret_key"`
	} `yaml:"auth"`
}

var cfg *Config

func NewConfig() (*Config, error) {
	return cfg, config.PopulateConfig(cfg)
}
