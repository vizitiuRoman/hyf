package memory

import (
	"context"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/vizitiuRoman/libs/cache"
)

type memCache struct {
	ctx        context.Context
	cache      *ristretto.Cache[string, any]
	cfg        *Config
	defaultTTL time.Duration
	log        cache.Logger
}

func NewMemoryCache(ctx context.Context, cfg *Config, log cache.Logger) (Cache, error) {
	var res, err = ristretto.NewCache(&ristretto.Config[string, any]{
		NumCounters:            cfg.NumCounters,
		MaxCost:                cfg.MaxCost,
		BufferItems:            cfg.BufferItems,
		Metrics:                false,
		OnEvict:                nil,
		OnReject:               nil,
		OnExit:                 nil,
		KeyToHash:              nil,
		Cost:                   nil,
		IgnoreInternalCost:     cfg.IgnoreInternalCost,
		TtlTickerDurationInSec: cfg.TTLTickerDurationInSec,
	})
	if err != nil {
		return nil, err
	}

	c := &memCache{
		ctx:        ctx,
		cache:      res,
		cfg:        cfg,
		defaultTTL: 0,
		log:        log,
	}

	if cfg.DefaultTTLInSec > 0 {
		c.defaultTTL = time.Duration(cfg.DefaultTTLInSec) * time.Second
	} else {
		c.defaultTTL = 0 * time.Second
	}

	log.Info("memory cache created")

	return c, nil
}

func (c *memCache) Clear() {
	c.cache.Clear()
}

func (c *memCache) Set(key string, value any) {
	c.cache.SetWithTTL(key, value, 1, c.defaultTTL)
}

func (c *memCache) SetWithTTL(key string, value any, ttl time.Duration) {
	_ = c.cache.SetWithTTL(key, value, 1, ttl)
}

func (c *memCache) Get(key string) (value any, found bool) {
	return c.cache.Get(key)
}

func (c *memCache) Del(key string) {
	c.cache.Del(key)
}

func (c *memCache) Close() {
	c.cache.Close()

	c.log.Info("memory cache closed")
}
