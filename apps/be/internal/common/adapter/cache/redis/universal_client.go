package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/helper/util"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"go.uber.org/zap"
)

type universalClient struct {
	ctx    context.Context
	cfg    *UniversalConfig
	logger log.Logger
	cln    redis.UniversalClient
}

func (c *universalClient) Init() error {
	if err := c.cln.Ping(c.ctx).Err(); err != nil {
		return err
	}

	c.logger.WithMethod(c.ctx, "Init").Info("redis client initialized")

	return nil
}

func NewUniversalClient(ctx context.Context, cfg *UniversalConfig, logger log.Logger) (Cache, error) {
	client := &universalClient{
		ctx:    ctx,
		cfg:    cfg,
		logger: logger.WithComponent(ctx, "redis-universal-client"),
	}

	redisClient := cfg.build()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	client.cln = redisClient

	return client, nil
}

func (c *universalClient) Set(ctx context.Context, key string, value *string) error {
	var s string
	if value == nil {
		s = ""
	} else {
		s = util.UnP(value)
	}
	err := c.cln.Set(ctx, key, s, time.Duration(c.cfg.DefaultTTL)*time.Second).Err()
	if err != nil {
		c.logger.WithMethod(ctx, "Set").Error("redis client set error", zap.Error(err), zap.String("key", key), zap.String("value", s))
		return err
	}

	return nil
}

func (c *universalClient) Get(ctx context.Context, key string) (*string, error) {
	val, err := c.cln.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		c.logger.WithMethod(ctx, "Get").Error("redis client get error", zap.Error(err), zap.String("key", key))
		return nil, err
	}

	return &val, nil
}

func (c *universalClient) MGet(ctx context.Context, keys ...string) ([]*string, error) {
	vals, err := c.cln.MGet(ctx, keys...).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		c.logger.WithMethod(ctx, "MGet").Error("redis client mget error", zap.Error(err), zap.Strings("keys", keys))
		return nil, err
	}

	res := make([]*string, len(vals))
	for i, val := range vals {
		if val == nil {
			continue
		}
		res[i] = util.P(val.(string))
	}

	return res, nil
}

func (c *universalClient) Del(ctx context.Context, keys ...string) error {
	err := c.cln.Del(ctx, keys...).Err()
	if err != nil {
		c.logger.WithMethod(ctx, "Del").Error("redis client del error", zap.Error(err), zap.Strings("keys", keys))
		return err
	}

	return nil
}

func (c *universalClient) KeysByPattern(ctx context.Context, pattern string) ([]string, error) {
	var (
		cursor   uint64
		uniqKeys = make(map[string]struct{})
		keys     []string
		err      error
		res      = make([]string, 0, 1)
	)

	for {
		keys, cursor, err = c.cln.Scan(ctx, cursor, pattern, c.cfg.ScanCount).Result()
		if err != nil {
			c.logger.WithMethod(ctx, "KeysByPattern").Error("redis client keys error", zap.Error(err))
			return nil, err
		}

		for _, key := range keys {
			if _, ok := uniqKeys[key]; ok {
				continue
			}
			uniqKeys[key] = struct{}{}
			res = append(res, key)
		}

		if cursor == 0 {
			break
		}
	}

	return res, nil
}

func (c *universalClient) Close() {
	if err := c.cln.Close(); err != nil {
		c.logger.Error("master redis connection close error", zap.Error(err))
		return
	}

	c.logger.Info("redis client closed")
}

func (c *universalClient) FlushDB(ctx context.Context) error {
	statusCmd := c.cln.FlushDB(ctx)
	if err := statusCmd.Err(); err != nil {
		c.logger.Error("redis client flushdb error", zap.Error(err))
	}

	return nil
}
