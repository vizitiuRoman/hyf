package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func chainServerUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var i int
		var next grpc.UnaryHandler
		next = func(ctx context.Context, req interface{}) (interface{}, error) {
			if i == len(interceptors)-1 {
				return interceptors[i](ctx, req, info, handler)
			}
			i++
			return interceptors[i-1](ctx, req, info, next)
		}
		return next(ctx, req)
	}
}
