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
)

type Config struct {
}

func WithConfig() TestingSuiteOption {
	return func(t *testing.T, suite suite.TestingSuite) {
		root := getRootDir()

		os.Setenv("ENV_FILE", path.Join(root, ".env"))
		os.Setenv("CONFIG_FILE", path.Join(root, "config.yaml"))
		conf, err := config.NewConfig()
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
			if field.Type() == reflect.TypeOf((*config.Config)(nil)) {
				if field.IsValid() && field.CanSet() {
					field.Set(reflect.ValueOf(conf))
					return
				}
			}
		}

		panic("field with type *config.Config not found in suite")
	}
}
