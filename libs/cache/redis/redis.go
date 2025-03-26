package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/vizitiuRoman/libs/cache"
	"go.uber.org/zap"
)

type redisCache struct {
	ctx        context.Context
	cfg        *Config
	defaultTTL time.Duration
	log        cache.Logger
	client     redis.UniversalClient
}

func NewRedisClient(ctx context.Context, cfg *Config, log cache.Logger) (Cache, error) {
	c := &redisCache{
		ctx: ctx,
		cfg: cfg,
		log: log,
		client: redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:                 cfg.Addrs,
			ClientName:            cfg.ClientName,
			DB:                    cfg.DB,
			Dialer:                nil,
			OnConnect:             nil,
			Protocol:              cfg.Protocol,
			Username:              cfg.Username,
			Password:              cfg.Password,
			SentinelUsername:      cfg.SentinelUsername,
			SentinelPassword:      cfg.SentinelPassword,
			MaxRetries:            cfg.MaxRetries,
			MinRetryBackoff:       cfg.MinRetryBackoff,
			MaxRetryBackoff:       cfg.MaxRetryBackoff,
			DialTimeout:           cfg.DialTimeout,
			ReadTimeout:           cfg.ReadTimeout,
			WriteTimeout:          cfg.WriteTimeout,
			ContextTimeoutEnabled: cfg.ContextTimeoutEnabled,
			PoolFIFO:              cfg.PoolFIFO,
			PoolSize:              cfg.PoolSize,
			PoolTimeout:           cfg.PoolTimeout,
			MinIdleConns:          cfg.MinIdleConns,
			MaxIdleConns:          cfg.MaxIdleConns,
			MaxActiveConns:        0,
			ConnMaxIdleTime:       cfg.ConnMaxIdleTime,
			ConnMaxLifetime:       cfg.ConnMaxLifetime,
			TLSConfig:             nil,
			MaxRedirects:          cfg.MaxRedirects,
			ReadOnly:              cfg.ReadOnly,
			RouteByLatency:        cfg.RouteByLatency,
			RouteRandomly:         cfg.RouteRandomly,
			MasterName:            cfg.MasterName,
			DisableIndentity:      false,
			IdentitySuffix:        ""}),
	}

	if cfg.DefaultTTLInSec > 0 {
		c.defaultTTL = time.Duration(cfg.DefaultTTLInSec) * time.Second
	} else {
		c.defaultTTL = 0 * time.Second
	}

	if err := c.client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	c.log.Info("redis client created")

	return c, nil
}

func (c *redisCache) Set(key string, value []byte) error {
	if err := c.client.Set(c.ctx, key, value, c.defaultTTL).Err(); err != nil {
		c.log.Error("cannot set", zap.Error(err), zap.String("key", key), zap.Duration("ttl", c.defaultTTL))
		return err
	}

	return nil
}

func (c *redisCache) SetWithTTL(key string, value []byte, ttl time.Duration) error {
	if err := c.client.Set(c.ctx, key, value, ttl).Err(); err != nil {
		c.log.Error("cannot set", zap.Error(err), zap.String("key", key), zap.Duration("ttl", ttl))
		return err
	}

	return nil
}

func (c *redisCache) Get(key string) ([]byte, error) {
	result, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrNoDataFound
		}
		c.log.Error("cannot get", zap.Error(err), zap.String("key", key))
	}

	if result == "" {
		return nil, err
	}

	return []byte(result), err
}

func (c *redisCache) Del(key string) error {
	if err := c.client.Del(c.ctx, key).Err(); err != nil {
		c.log.Error("cannot delete", zap.Error(err), zap.String("key", key))
		return err
	}

	return nil
}

func (c *redisCache) Close() {
	if err := c.client.Close(); err != nil {
		c.log.Error("cannot close redis connection", zap.Error(err))
		return
	}

	c.log.Info("redis client closed")
}

func (c *redisCache) flushDB() error {
	statusCmd := c.client.FlushDB(c.ctx)
	if err := statusCmd.Err(); err != nil {
		c.log.Error("cannot flush redis db", zap.Error(err))
		return err
	}

	return nil
}
