package helper

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestingSuiteOption func(t *testing.T, s suite.TestingSuite)

func MustSuiteWithOptions(t *testing.T, s suite.TestingSuite, opts ...TestingSuiteOption) suite.TestingSuite {
	for _, opt := range opts {
		opt(t, s)
	}

	return s
}
