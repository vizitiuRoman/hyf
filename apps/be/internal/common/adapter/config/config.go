package config

import (
	"os"
	"time"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/db"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/server/grpc"
	"go.uber.org/config"
	"go.uber.org/zap"
)

type Config struct {
	Location string `yaml:"location"`

	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	BuildDate string `yaml:"build_date"`

	Cache  *cache.Config `yaml:"cache"`
	Logger *zap.Config   `yaml:"logger"`
	DB     *db.Config    `yaml:"db"`
	Server *grpc.Config  `yaml:"server"`

	Auth *struct {
		AccessTokenSecretKey  string `yaml:"access_token_secret_key"`
		RefreshTokenSecretKey string `yaml:"refresh_token_secret_key"`
	} `yaml:"auth"`
}

var cfg *Config

const defaultConfigYaml = "./config.yaml"

var (
	name      = "hyf-svc"
	version   = "undefined/local"
	buildDate = time.Now().Format(time.RFC3339)
)

func NewConfig() (*Config, error) {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = defaultConfigYaml
	}

	provider, err := NewProviderByOptions(config.File(configFile))
	if err != nil {
		return nil, err
	}
	if err = provider.Populate(&cfg); err != nil {
		panic(err)
	}

	cfg.Name = name
	cfg.Version = version
	cfg.BuildDate = buildDate

	return cfg, nil
}
