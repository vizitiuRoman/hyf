package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// Config is used to configure a Redis Client.
// see: github.com/redis/go-redis/v9/options.go
type Config struct {
	// The network type, either tcp or unix.
	// Default is tcp.
	Network string `yaml:"network"`

	// host:port address.
	Addr string `yaml:"addr"`

	// TTL of every key that is being set in seconds.
	DefaultTTL int `yaml:"default_ttl"`

	// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
	ClientName string `yaml:"client_name"`

	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	Username string `yaml:"username"`

	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	Password string `yaml:"password"`

	// Database to be selected after connecting to the server.
	DB int `yaml:"db"`

	// Maximum number of retries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	MaxRetries int `yaml:"max_retries"`

	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	// Default is 0. the idle connections are not closed by default.
	MinIdleConns int `yaml:"min_idle_conns"`

	// Maximum number of idle connections.
	// Default is 0. the idle connections are not closed by default.
	MaxIdleConns int `yaml:"max_idle_conns"`
}

type UniversalConfig struct {
	// Either a single address or a seed list of host:port addresses
	// of cluster/sentinel nodes.
	Addrs []string `yaml:"addrs"`

	// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
	ClientName string `yaml:"client_name"`

	// Database to be selected after connecting to the server.
	// Only single-node and failover clients.
	DB int `yaml:"db"`

	// ScanCount is the value passed as COUNT option of the SCAN command.
	ScanCount int64 `yaml:"scan_count"`

	// TTL of every key that is being set in seconds.
	DefaultTTL int `yaml:"default_ttl"`

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

func (cfg *UniversalConfig) build() redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      cfg.Addrs,
		ClientName: cfg.ClientName,
		DB:         cfg.DB,

		Protocol:         cfg.Protocol,
		Username:         cfg.Username,
		Password:         cfg.Password,
		SentinelUsername: cfg.SentinelUsername,
		SentinelPassword: cfg.SentinelPassword,

		MaxRetries:      cfg.MaxRetries,
		MinRetryBackoff: cfg.MinRetryBackoff,
		MaxRetryBackoff: cfg.MaxRetryBackoff,

		DialTimeout:           cfg.DialTimeout,
		ReadTimeout:           cfg.ReadTimeout,
		WriteTimeout:          cfg.WriteTimeout,
		ContextTimeoutEnabled: cfg.ContextTimeoutEnabled,

		PoolFIFO: cfg.PoolFIFO,

		PoolSize:        cfg.PoolSize,
		PoolTimeout:     cfg.PoolTimeout,
		MinIdleConns:    cfg.MinIdleConns,
		MaxIdleConns:    cfg.MaxIdleConns,
		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
		MaxRedirects:    cfg.MaxRedirects,
		ReadOnly:        cfg.ReadOnly,
		RouteByLatency:  cfg.RouteByLatency,
		RouteRandomly:   cfg.RouteRandomly,
		MasterName:      cfg.MasterName,
	})
}

func (cfg *Config) build() *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:               cfg.Network,
		Addr:                  cfg.Addr,
		ClientName:            cfg.ClientName,
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              cfg.Username,
		Password:              cfg.Password,
		CredentialsProvider:   nil,
		DB:                    cfg.DB,
		MaxRetries:            cfg.MaxRetries,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          cfg.MinIdleConns,
		MaxIdleConns:          cfg.MaxIdleConns,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
	})
}
