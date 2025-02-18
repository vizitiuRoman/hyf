package redis

import (
	"time"
)

// Config is used to configure a Redis Client.
// see: github.com/redis/go-redis/v9/universal.go
type Config struct {
	// Either a single address or a seed list of host:port addresses
	// of cluster/sentinel nodes.
	Addrs []string `yaml:"addrs"`

	// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
	ClientName string `yaml:"client_name"`

	// Database to hyf-service selected after connecting to the server.
	// Only single-node and failover clients.
	DB int `yaml:"db"`

	// ScanCount is the value passed as COUNT option of the SCAN command.
	ScanCount int64 `yaml:"scan_count"`

	// DefaultTTLInSec of every key that is being set in seconds.
	DefaultTTLInSec int64 `yaml:"default_ttl_in_sec"`

	// Common options.

	Protocol         int    `yaml:"protocol"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	SentinelUsername string `yaml:"sentinel_username"`
	SentinelPassword string `yaml:"sentinel_password"`

	MaxRetries      int           `yaml:"max_retries"`
	MinRetryBackoff time.Duration `yaml:"min_retry_backoff"`
	MaxRetryBackoff time.Duration `yaml:"max_retry_backoff"`

	DialTimeout           time.Duration `yaml:"dial_timeout"`
	ReadTimeout           time.Duration `yaml:"read_timeout"`
	WriteTimeout          time.Duration `yaml:"write_timeout"`
	ContextTimeoutEnabled bool          `yaml:"context_timeout_enabled"`

	// PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
	PoolFIFO bool `yaml:"pool_fifo"`

	PoolSize        int           `yaml:"pool_size"`
	PoolTimeout     time.Duration `yaml:"pool_timeout"`
	MinIdleConns    int           `yaml:"min_idle_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`

	// Only cluster clients.

	MaxRedirects   int  `yaml:"max_redirects"`
	ReadOnly       bool `yaml:"read_only"`
	RouteByLatency bool `yaml:"route_by_latency"`
	RouteRandomly  bool `yaml:"route_randomly"`

	// The sentinel master name.
	// Only failover clients.

	MasterName string `yaml:"master_name"`
}
