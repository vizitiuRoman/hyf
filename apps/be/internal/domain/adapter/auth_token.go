package adapter

import (
	"context"

	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

type AuthTokenAdapterFactory interface {
	Create(context.Context) AuthTokenAdapter
}

type AuthTokenAdapter interface {
	MarshalJSON(context.Context, *model.AuthToken) (*string, error)
	UnmarshalJSON(context.Context, *string) (*model.AuthToken, error)
	ToProto(*model.AuthToken) *pb.AuthOut
	FromUserID(context.Context, int64) (*model.AuthToken, error)

	FromRegisterProtoToUser(*pb.RegisterIn) *model.User
}
