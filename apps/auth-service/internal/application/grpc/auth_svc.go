package grpc

import (
	"context"

	"github.com/vizitiuRoman/auth-service/internal/common/adapter/server/grpc"
	"github.com/vizitiuRoman/auth-service/internal/infra/adapter"
	"github.com/vizitiuRoman/auth-service/internal/infra/service"
	pb "github.com/vizitiuRoman/auth-service/pkg/gen/auth/v1"
	"github.com/vizitiuRoman/libs/logger"
)

type authSVC struct {
	pb.UnimplementedAuthSVCServer
	ctx    context.Context
	logger logger.Logger

	adapter *adapter.AuthAdapter

	authService service.AuthService
}

func NewAuthSVCServerDescriptor(
	ctx context.Context,
	logger logger.Logger,

	adapter *adapter.AuthAdapter,

	authService service.AuthService,
) *grpc.ServerDescriptor {
	server := &authSVC{
		ctx:    ctx,
		logger: logger,

		adapter: adapter,

		authService: authService,
	}

	return &grpc.ServerDescriptor{
		Server:               server,
		GRPCRegistrar:        pb.RegisterAuthSVCServer,
		GRPCGatewayRegistrar: pb.RegisterAuthSVCHandlerFromEndpoint,
		Methods: []grpc.MethodDescriptor{
			{
				Method: (*authSVC).Login,
			},
			{
				Method: (*authSVC).Register,
			},
			{
				Method: (*authSVC).Refresh,
			},
		},
	}
}

func (s *authSVC) Login(ctx context.Context, input *pb.LoginIn) (*pb.AuthOut, error) {
	accessToken, refreshToken, err := s.authService.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.adapter.ToProto(accessToken, refreshToken), nil
}

func (s *authSVC) Register(ctx context.Context, input *pb.RegisterIn) (*pb.AuthOut, error) {
	accessToken, refreshToken, err := s.authService.Register(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.adapter.ToProto(accessToken, refreshToken), nil
}

func (s *authSVC) Refresh(ctx context.Context, input *pb.RefreshIn) (*pb.AuthOut, error) {
	accessToken, refreshToken, err := s.authService.Refresh(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.adapter.ToProto(accessToken, refreshToken), nil
}

func (s *authSVC) Logout(ctx context.Context, input *pb.LogoutIn) (*pb.LogoutOut, error) {
	err := s.authService.Logout(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.LogoutOut{}, nil
}
