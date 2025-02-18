package redis

import (
	"time"
)

type Cache interface {
	Set(key string, value []byte) error
	SetWithTTL(key string, value []byte, ttl time.Duration) error
	Get(string) ([]byte, error)
	Del(string) error
	Close()
}
