package service

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type AuthService interface {
	Login(context.Context, *pb.LoginIn) (*model.AuthToken, error)
	Register(context.Context, *pb.RegisterIn) (*model.AuthToken, error)
}
