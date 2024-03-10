package auth

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/cache/redis"
	"github.com/vizitiuRoman/hyf/internal/domain/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthInterceptor(_ context.Context, rdb redis.Cache, authTokenRepoFactory repo.AuthTokenRepoFactory) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		switch {
		case info.FullMethod == "/AuthSVC/Login" || info.FullMethod == "/AuthSVC/Register":
			return handler(ctx, req)

		default:
			resp, respErr := handler(ctx, req)

			token, err := grpc_auth.AuthFromMD(ctx, "bearer")
			if err != nil {
				return nil, err
			}

			t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				return []byte("@TODO"), nil
			})
			if err != nil {
				return nil, fmt.Errorf("validate: %w", err)
			}

			claims, ok := t.Claims.(jwt.MapClaims)
			if !ok || !t.Valid {
				return ctx, status.Error(codes.Unauthenticated, "wrong token")
			}

			if claims["uuid"] == nil {
				return ctx, status.Error(codes.Unauthenticated, "wrong claim: uuid doesn't exist")
			}

			authTokenRepo := authTokenRepoFactory.Create(ctx, rdb)
			_, err = authTokenRepo.GetByUUID(ctx, uuid.FromStringOrNil(claims["uuid"].(string)))
			if err != nil {
				return ctx, status.Error(codes.Unauthenticated, err.Error())
			}

			return resp, respErr
		}
	}
}
