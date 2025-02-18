package auth

import (
	"context"

	"github.com/gofrs/uuid"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/vizitiuRoman/auth-service/internal/infra/adapter"
	"github.com/vizitiuRoman/auth-service/internal/infra/repo/cache"
	"github.com/vizitiuRoman/libs/cache/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthInterceptor(_ context.Context, rdb redis.Cache, authTokenRepoFactory *cache.AuthRepoFactory, adapter *adapter.AuthAdapter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		switch {
		case info.FullMethod == "/AuthSVC/Login" || info.FullMethod == "/AuthSVC/Register":
			return handler(ctx, req)

		default:
			token, err := grpc_auth.AuthFromMD(ctx, "bearer")
			if err != nil {
				return nil, err
			}

			uuidClaim, _, err := adapter.ClaimsFromJWT(ctx, token)
			if err != nil {
				return ctx, status.Error(codes.Unauthenticated, "wrong token")
			}

			if uuidClaim == "" {
				return ctx, status.Error(codes.Unauthenticated, "wrong claim: uuid doesn't exist")
			}

			authTokenRepo := authTokenRepoFactory.Create(ctx, rdb)
			_, err = authTokenRepo.GetByUUID(ctx, uuid.FromStringOrNil(uuidClaim))
			if err != nil {
				return ctx, status.Error(codes.Unauthenticated, err.Error())
			}

			return handler(ctx, req)
		}
	}
}
