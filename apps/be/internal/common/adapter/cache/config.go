package cache

import "github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"

type Config struct {
	Redis *redis.UniversalConfig `yaml:"redis"`
}
