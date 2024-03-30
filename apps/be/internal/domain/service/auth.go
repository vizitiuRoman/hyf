package service

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type AuthService interface {
	Login(context.Context, *pb.LoginIn) (accessToken, refreshToken *model.AuthToken, err error)
	Register(context.Context, *pb.RegisterIn) (accessToken, refreshToken *model.AuthToken, err error)
	Refresh(context.Context, *pb.RefreshIn) (accessToken, refreshToken *model.AuthToken, err error)
	Logout(context.Context, *pb.LogoutIn) (err error)
}
