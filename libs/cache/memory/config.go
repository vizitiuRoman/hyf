package memory

// Config is a configuration for cache
// see: github.com/dgraph-io/ristretto/cache.go
type Config struct {
	NumCounters            int64 `yaml:"num_counters"`
	MaxCost                int64 `yaml:"max_cost"`
	BufferItems            int64 `yaml:"buffer_items"`
	IgnoreInternalCost     bool  `yaml:"ignore_internal_cost"`
	TTLTickerDurationInSec int64 `yaml:"ttl_ticker_duration_in_sec"`
	// DefaultTTLInSec of every key that is being set in seconds.
	DefaultTTLInSec int64 `yaml:"default_ttl_in_sec"`
}
