package auth

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StoredToken struct {
	Key string
}

func NewStoredToken() *StoredToken {
	return &StoredToken{
		Key: "OMG",
	}
}

func (a *StoredToken) Authenticate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		// Unauthenticated error
		return nil, err
	}

	if token != a.Key {
		return ctx, status.Error(codes.Unauthenticated, "Wrong token")
	}

	return ctx, nil
}
