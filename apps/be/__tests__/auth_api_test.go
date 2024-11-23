package __tests__

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vizitiuRoman/hyf/__tests__/helper"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	hyfv1 "github.com/vizitiuRoman/hyf/pkg/adapter/hyf/v1"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestAuthAPISuite(t *testing.T) {
	suite.Run(
		t,
		helper.MustSuiteWithOptions(t, new(AuthAPISuite), helper.WithConfig(), helper.WithLogger()),
	)
}

type AuthAPISuite struct {
	suite.Suite
	ctx    context.Context
	cancel context.CancelFunc

	authSVCClient pb.AuthSVCClient

	Cfg    *helper.Config
	Logger logger.Logger

	// suite
	mockUser *model.User

	refreshToken string
	accessToken  string
}

func (s *AuthAPISuite) SetupSuite() {
	var err error
	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.authSVCClient, err = hyfv1.NewAuthSVCClient(
		s.ctx,
		s.Logger,
		&hyfv1.Config{
			Host:     s.Cfg.Grpc.Host,
			GrpcPort: s.Cfg.Grpc.Port,
		},
		[]grpc.DialOption{}, nil,
	)
	s.Nil(err)

	s.mockUser = &model.User{
		Email:    s.Cfg.User.Email,
		Password: s.Cfg.User.Password,
		Name:     "omg",
		LastName: "oqo",
	}
}

func (s *AuthAPISuite) TestARegister() {
	out, err := s.authSVCClient.Register(s.ctx, &pb.RegisterIn{
		Email:    s.mockUser.Email,
		Password: s.mockUser.Password,
		Name:     s.mockUser.Name,
		LastName: s.mockUser.LastName,
	})
	s.Nil(err)

	s.NotNil(out)
	s.NotNil(out.Token)
	s.NotNil(out.ExpiresIn)

	s.refreshToken = out.RefreshToken
}

func (s *AuthAPISuite) TestBLogin() {
	out, err := s.authSVCClient.Login(s.ctx, &pb.LoginIn{
		Email:    s.mockUser.Email,
		Password: s.mockUser.Password,
	})
	s.Nil(err)

	s.NotNil(out)
	s.NotNil(out.Token)
	s.NotNil(out.ExpiresIn)

	s.refreshToken = out.RefreshToken
	s.accessToken = out.Token
}

func (s *AuthAPISuite) TestCRefresh() {
	md := metadata.New(map[string]string{"authorization": "Bearer " + s.accessToken})
	ctx := metadata.NewOutgoingContext(s.ctx, md)

	out, err := s.authSVCClient.Refresh(ctx, &pb.RefreshIn{
		RefreshToken: s.refreshToken,
	})
	s.Nil(err)

	s.NotNil(out)
	s.NotNil(out.Token)
	s.NotNil(out.RefreshToken)
	s.NotNil(out.ExpiresIn)

	s.refreshToken = out.RefreshToken
	s.accessToken = out.Token
}

func (s *AuthAPISuite) TestDLogout() {
	md := metadata.New(map[string]string{"authorization": "Bearer " + s.accessToken})
	ctx := metadata.NewOutgoingContext(s.ctx, md)

	_, err := s.authSVCClient.Logout(ctx, &pb.LogoutIn{
		Token:        s.accessToken,
		RefreshToken: s.refreshToken,
	})
	s.Nil(err)

	out, err := s.authSVCClient.Refresh(s.ctx, &pb.RefreshIn{
		RefreshToken: s.refreshToken,
	})
	s.NotNil(err)

	s.Nil(out)
}
