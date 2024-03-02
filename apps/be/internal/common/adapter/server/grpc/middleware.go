package grpc

import (
	"google.golang.org/grpc"
)

type Priority uint32

const (
	MaxPriority = Priority(^uint32(0) >> 1)
	MinPriority = Priority(0)
)

type Middleware struct {
	Priority    Priority
	Interceptor grpc.UnaryServerInterceptor
}

type ByOrder []Middleware

func (b ByOrder) Len() int {
	return len(b)
}

func (b ByOrder) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByOrder) Less(i, j int) bool {
	return b[i].Priority > b[j].Priority
}
