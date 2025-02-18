package memory

import (
	"time"
)

type Cache interface {
	Set(key string, value any)
	SetWithTTL(key string, value any, ttl time.Duration)
	Get(key string) (value any, found bool)
	Del(key string)
	Close()
}
