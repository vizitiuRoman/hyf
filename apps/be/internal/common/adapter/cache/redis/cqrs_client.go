package redis

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/helper/util"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"go.uber.org/zap"
)

type cqrsClient struct {
	ctx       context.Context
	cfg       *Config
	logger    log.Logger
	masterRDB *redis.Client
	slaveRDB  *redis.Client
}

func (c *cqrsClient) Init() error {
	if err := c.masterRDB.Ping(c.ctx).Err(); err != nil {
		return err
	}

	if err := c.slaveRDB.Ping(c.ctx).Err(); err != nil {
		return err
	}

	c.logger.WithMethod(c.ctx, "Init").Info("redis client initialized")

	return nil
}

func NewCQSRClient(ctx context.Context, cfg *Config, logger log.Logger) (Cache, error) {
	cln := &cqrsClient{
		ctx:    ctx,
		cfg:    cfg,
		logger: logger.WithComponent(ctx, "redis-cqrs-client"),
	}

	rdb := cfg.build()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	masterIP, isMaster, err := cln.getMasterIP(ctx, rdb, cln.logger)
	if err != nil {
		return nil, err
	}
	slaveIP, isSlave, err := cln.getSlaveIP(ctx, rdb, cln.logger)
	if err != nil {
		return nil, err
	}

	if isMaster {
		cln.masterRDB = rdb

		if slaveIP != "" {
			slaveCfg := *cfg
			slaveCfg.Addr = slaveIP + ":6379"

			cln.slaveRDB = slaveCfg.build()
			if err := cln.slaveRDB.Ping(ctx).Err(); err != nil {
				return nil, err
			}
		} else {
			cln.slaveRDB = rdb
		}
	}

	if isSlave {
		cln.slaveRDB = rdb

		masterCfg := *cfg
		masterCfg.Addr = masterIP + ":6379"

		cln.masterRDB = masterCfg.build()
		if err := cln.slaveRDB.Ping(ctx).Err(); err != nil {
			return nil, err
		}
	}

	return cln, nil
}

func (c *cqrsClient) getMasterIP(ctx context.Context, rdb *redis.Client, logger log.Logger) (string, bool, error) {
	vals, err := rdb.Info(ctx, "replication").Result()
	if err != nil {
		logger.WithMethod(ctx, "getMasterIP").Error("unable to get replication info", zap.Error(err))
		return "", false, err
	}

	master := false
	masterIP := ""
	for _, s := range strings.Split(vals, "\n") {
		if strings.HasPrefix(s, "master_host") {
			masterIP = strings.TrimSpace(strings.Split(s, ":")[1])
		}
		if strings.Contains(s, "role:master") {
			master = true
		}
	}

	return masterIP, master, nil
}

func (c *cqrsClient) getSlaveIP(ctx context.Context, rdb *redis.Client, logger log.Logger) (string, bool, error) {
	vals, err := rdb.Info(ctx, "replication").Result()
	if err != nil {
		logger.WithMethod(ctx, "getSlaveIP").Error("unable to get replication info", zap.Error(err))
		return "", false, err
	}

	slave := true
	slaveIP := ""
	for _, text := range strings.Split(vals, "\n") {
		if strings.HasPrefix(text, "slave") {
			match := regexp.MustCompile(`ip=([0-9.]+)`).FindStringSubmatch(text)
			if len(match) > 1 {
				slaveIP = match[1]
			}
		}
		if strings.Contains(text, "role:master") {
			slave = false
		}
	}

	return slaveIP, slave, nil
}

func (c *cqrsClient) Set(ctx context.Context, key string, value *string) error {
	var s string
	if value == nil {
		s = ""
	} else {
		s = util.UnP(value)
	}
	err := c.masterRDB.Set(ctx, key, s, time.Duration(c.cfg.DefaultTTL)*time.Second).Err()
	if err != nil {
		c.logger.WithMethod(ctx, "Set").Error("redis client set error", zap.Error(err), zap.String("key", key), zap.String("value", s))
		return err
	}

	return nil
}

func (c *cqrsClient) Get(ctx context.Context, key string) (*string, error) {
	val, err := c.slaveRDB.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		c.logger.WithMethod(ctx, "Get").Error("redis client get error", zap.Error(err), zap.String("key", key))
		return nil, err
	}

	return &val, nil
}

func (c *cqrsClient) MGet(ctx context.Context, keys ...string) ([]*string, error) {
	vals, err := c.slaveRDB.MGet(ctx, keys...).Result()
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

func (c *cqrsClient) Del(ctx context.Context, keys ...string) error {
	err := c.masterRDB.Del(ctx, keys...).Err()
	if err != nil {
		c.logger.WithMethod(ctx, "Del").Error("redis client del error", zap.Error(err), zap.Strings("keys", keys))
		return err
	}

	return nil
}

//func (c *client) KeyValuesByPattern(ctx context.Context, pattern string) ([][2]string, error) {
//	stringSliceCmd := c.cln.Keys(ctx, pattern)
//	if err := stringSliceCmd.Err(); err != nil {
//		c.logger.WithMethod(ctx, "KeyValuesByPattern").Error("redis client keys error", zap.Error(err))
//		return nil, err
//	}
//
//	keys := stringSliceCmd.Val()
//	res := make([][2]string, len(keys))
//
//	for _, key := range keys {
//		val, err := c.cln.GetAll(ctx, key).Result()
//		if errors.Is(err, redis.Nil) {
//			continue
//		} else if err != nil {
//			return nil, err
//		}
//
//		res = append(res, [2]string{key, val})
//	}
//
//	return res, nil
//}

func (c *cqrsClient) KeysByPattern(ctx context.Context, pattern string) ([]string, error) {
	stringSliceCmd := c.slaveRDB.Keys(ctx, pattern)
	if err := stringSliceCmd.Err(); err != nil {
		c.logger.WithMethod(ctx, "KeysByPattern").Error("redis client keys error", zap.Error(err), zap.String("pattern", pattern))
		return nil, err
	}

	return stringSliceCmd.Val(), nil
}

func (c *cqrsClient) Close() {
	var isError bool
	if err := c.masterRDB.Close(); err != nil {
		c.logger.Error("master redis connection close error", zap.Error(err))
		isError = true
	}
	if err := c.slaveRDB.Close(); err != nil {
		c.logger.Error("master redis connection close error", zap.Error(err))
		isError = true
	}

	if !isError {
		c.logger.Info("redis client closed")
	}
}

func (c *cqrsClient) FlushDB(ctx context.Context) error {
	statusCmd := c.masterRDB.FlushDB(ctx)
	if err := statusCmd.Err(); err != nil {
		c.logger.Error("redis client flushdb error", zap.Error(err))
	}

	return nil
}
