package config

import (
	"go.uber.org/zap"

	"github.com/vizitiuRoman/libs/config"
)

type DB struct {
	DSN             string `yaml:"dsn"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type Server struct {
	GrpcPort int  `yaml:"grpc_port"`
	HttpPort int  `yaml:"http_port"`
	UseTls   bool `yaml:"use_tls"`
}

type Config struct {
	Location string `yaml:"location"`
	DB       DB     `yaml:"db"`
	Server   Server `yaml:"server"`
}

var cfg = &Config{}

func NewConfig(logger *zap.Logger) (*Config, error) {
	err := config.PopulateConfig(cfg, config.WithLogger(logger))

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
