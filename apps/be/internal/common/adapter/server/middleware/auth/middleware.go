package auth

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthInterceptor(_ context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, respErr := handler(ctx, req)

		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			// Unauthenticated error
			return nil, err
		}

		if token != "omg" {
			return ctx, status.Error(codes.Unauthenticated, "Wrong token")
		}

		return resp, respErr
	}
}
