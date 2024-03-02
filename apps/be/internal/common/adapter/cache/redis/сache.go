package redis

import "context"

type Cache interface {
	Set(context.Context, string, *string) error
	Get(context.Context, string) (*string, error)
	MGet(context.Context, ...string) ([]*string, error)
	Del(context.Context, ...string) error
	KeysByPattern(context.Context, string) ([]string, error)
	Init() error
	Close()
	FlushDB(context.Context) error
}
