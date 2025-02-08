package main

import (
	"context"
	"strconv"
	"time"

	"github.com/vizitiuRoman/libs/pgclient"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.uber.org/zap"
)

func NewFxPool(ctx context.Context, logger *zap.Logger) (pgclient.DB, error) {
	cfg := &pgclient.Config{
		DSN:             "postgres://postgres:password@localhost:54321/postgres?search_path=public&sslmode=disable",
		MaxIdleConns:    5,
		ConnMaxLifetime: 5,
		ConnMaxIdleTime: 5,
		MaxOpenConns:    5,
	}

	if err := runDockerPostgres(cfg, logger); err != nil {
		logger.Error("unable to run test nats docker", zap.Error(err))
		return nil, err
	}

	pool, err := pgclient.NewPool(ctx, cfg, logger, pgclient.WithTransactionLogging(logger))
	if err != nil {
		logger.Error("unable to create pool", zap.Error(err))
		return nil, err
	}

	return pool, nil
}

func runDockerPostgres(cfg *pgclient.Config, logger *zap.Logger) error {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return err
	}

	if err = pool.Client.Ping(); err != nil {
		logger.Error("could not connect to docker", zap.Error(err))
		return err
	}

	pgConfig, err := pgconn.ParseConfig(cfg.DSN)
	if err != nil {
		logger.Error("unable to parse postgres DSN", zap.Error(err))
		return err
	}

	resource, ok := pool.ContainerByName("exapi-postgres-client-client-test")
	if !ok {
		resource, err = pool.RunWithOptions(&dockertest.RunOptions{
			Name:       "exapi-postgres-client-client-test",
			Repository: "postgres",
			Tag:        "15",
			Env: []string{
				"POSTGRES_HOST=" + pgConfig.Host,
				"POSTGRES_PORT=5432",
				"POSTGRES_DB=" + pgConfig.Database,
				"POSTGRES_SCHEMA=" + pgConfig.RuntimeParams["search_path"],
				"POSTGRES_USER=" + pgConfig.User,
				"POSTGRES_PASSWORD=" + pgConfig.Password,
			},
			PortBindings: map[docker.Port][]docker.PortBinding{
				"5432/tcp": {{HostPort: strconv.Itoa(int(pgConfig.Port)) + "/tcp"}},
			},
		}, func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		})
		if err != nil {
			return err
		}
	}

	time.Sleep(time.Second * 3)
	return resource.Expire(5000)
}
