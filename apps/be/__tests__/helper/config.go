package helper

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/config"

	uberconfig "go.uber.org/config"
)

type Config struct {
	Grpc struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"grpc"`

	User struct {
		Email    string `yaml:"email"`
		Password string `yaml:"password"`
	}
}

const defaultConfigYaml = "./helper/config.yaml"

var cfg *Config

func newConfig() (*Config, error) {
	provider, err := config.NewProviderByOptions(uberconfig.File(defaultConfigYaml))
	if err != nil {
		return nil, err
	}
	if err = provider.Populate(&cfg); err != nil {
		panic(err)
	}

	return cfg, nil
}

func WithConfig() TestingSuiteOption {
	return func(t *testing.T, suite suite.TestingSuite) {
		root := getRootDir()

		os.Setenv("ENV_FILE", path.Join(root, "./helper/.env"))
		os.Setenv("CONFIG_FILE", path.Join(root, "config.yaml"))
		conf, err := newConfig()
		if err != nil {
			dir, errWd := os.Getwd()
			if errWd == nil {
				err = errors.New(fmt.Sprintf("failed to load config: %s: %v", dir, err))
			}
			panic(err)
		}

		subSuiteType := reflect.ValueOf(suite).Elem()
		for i := 0; i < subSuiteType.NumField(); i++ {
			field := subSuiteType.Field(i)
			if field.Type() == reflect.TypeOf((*Config)(nil)) {
				if field.IsValid() && field.CanSet() {
					field.Set(reflect.ValueOf(conf))
					return
				}
			}
		}

		panic("field with type *config.Config not found in suite")
	}
}
