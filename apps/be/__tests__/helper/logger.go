package helper

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
)

func WithLogger() TestingSuiteOption {
	return func(t *testing.T, suite suite.TestingSuite) {
		logger := log.MustDefaultConsoleLogger("debug")
		subSuiteType := reflect.ValueOf(suite).Elem()
		for i := 0; i < subSuiteType.NumField(); i++ {
			field := subSuiteType.Field(i)
			if field.Kind() == reflect.Interface && field.Type().Implements(reflect.TypeOf((*log.Logger)(nil)).Elem()) {
				field.Set(reflect.ValueOf(logger))
				return
			}
		}
		panic("interface field assignable to log.Logger not found in suite")
	}
}
