package middleware

import (
	"context"
	"time"

	"github.com/vizitiuRoman/libs/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func LoggerMiddleware(_ context.Context, logger logger.Logger) grpc.UnaryServerInterceptor {
	interceptor := grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		l := logger.With(
			zap.Any("debug.request.data", req),
			zap.String("debug.request.method", info.FullMethod),
		)

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			l = logger.With(zap.Any("debug.request.metadata", md))
		}

		start := time.Now()

		resp, respErr := handler(ctx, req)
		code, ok := status.FromError(respErr)
		if !ok {
			code = status.New(codes.Code(600), "")
		}

		l.Debug("request served",
			zap.Any("debug.response.body", resp),
			zap.Int("debug.response.status_code", int(code.Code())),
			zap.String("debug.response.status_message", code.Message()),
			zap.Duration("debug.request.duration", time.Since(start)),
		)

		return resp, respErr
	})

	return chainServerUnaryInterceptors(interceptor)

}
