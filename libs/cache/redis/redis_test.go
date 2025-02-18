package redis

import (
	"context"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

func TestUniversalClientSuite(t *testing.T) {
	suite.Run(t, new(UniversalClientSuite))
}

type UniversalClientSuite struct {
	suite.Suite
	client *redisCache
}

var (
	containerName   = "redis-mascot-for-test"
	localhostConfig = &Config{
		Addrs:                 []string{"127.0.0.1:6379"},
		ClientName:            "",
		DB:                    0,
		ScanCount:             0,
		DefaultTTLInSec:       0,
		Protocol:              0,
		Username:              "",
		Password:              "",
		SentinelUsername:      "",
		SentinelPassword:      "",
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		MaxRedirects:          0,
		ReadOnly:              false,
		RouteByLatency:        false,
		RouteRandomly:         false,
		MasterName:            "",
	}

	_ = &Config{
		Addrs:                 []string{"rfs-redis-main.dev.svc.cluster.local:26379"},
		ClientName:            "",
		DB:                    0,
		ScanCount:             0,
		DefaultTTLInSec:       60 * 60 * 24,
		Protocol:              0,
		Username:              "",
		Password:              "verystrongpassword",
		SentinelUsername:      "",
		SentinelPassword:      "",
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		MaxRedirects:          0,
		ReadOnly:              false,
		RouteByLatency:        false,
		RouteRandomly:         false,
		MasterName:            "",
	}
)

func (s *UniversalClientSuite) SetupSuite() {
	var err error

	log, err := zap.NewDevelopmentConfig().Build()
	s.NoError(err)

	cfg := localhostConfig

	if cfg == localhostConfig {
		dockerPool, err := dockertest.NewPool("")
		s.NoError(err)

		resource, ok := dockerPool.ContainerByName(containerName)
		if !ok {
			resource, err = dockerPool.RunWithOptions(&dockertest.RunOptions{
				Name:       containerName,
				Repository: "redis",
				Tag:        "6.2-alpine",
				PortBindings: map[docker.Port][]docker.PortBinding{
					"6379/tcp": {
						{
							HostIP:   "127.0.0.1",
							HostPort: "6379/tcp",
						},
					},
				},
			}, func(config *docker.HostConfig) {
				config.AutoRemove = true
				config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			})
			s.NoError(err)
		}

		s.NoError(resource.Expire(60))
	}

	client, err := NewRedisClient(context.Background(), cfg, log)
	s.NoError(err)

	s.client = client.(*redisCache)
}

func (s *UniversalClientSuite) TearDownSuite() {
	s.NoError(s.client.flushDB())
	s.NoError(s.client.client.Close())
}

func (s *UniversalClientSuite) Test() {
	k, v := "key", []byte("value")
	s.NoError(s.client.Set(k, v))
	newV, err := s.client.Get(k)
	s.NoError(err)
	s.Equal(v, newV)

	k, v = "keyTTL", []byte("value2")
	s.NoError(s.client.SetWithTTL(k, v, time.Second))
	newV, err = s.client.Get(k)
	s.NoError(err)
	s.Equal(v, newV)
	<-time.After(time.Millisecond * 1500)
	newV, err = s.client.Get(k)
	s.Nil(newV)
	s.Error(err)
	s.Equal(ErrNoDataFound, err)
}
