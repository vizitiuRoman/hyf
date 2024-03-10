package adapter

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/config"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/adapter"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"go.uber.org/zap"
)

type authTokenAdapterFactory struct {
	logger log.Logger

	cfg *config.Config
}

func NewAuthTokenAdapterFactory(logger log.Logger, cfg *config.Config) adapter.AuthTokenAdapterFactory {
	return &authTokenAdapterFactory{
		logger: logger,
		cfg:    cfg,
	}
}

func (f *authTokenAdapterFactory) Create(ctx context.Context) adapter.AuthTokenAdapter {
	return &authTokenAdapter{
		logger: f.logger.WithComponent(ctx, "AuthTokenAdapter"),

		jwtSecretKey: f.cfg.Auth.JWTSecretKey,
	}
}

type authTokenAdapter struct {
	logger log.Logger

	jwtSecretKey string
}

func (a *authTokenAdapter) MarshalJSON(ctx context.Context, token *model.AuthToken) (*string, error) {
	if token == nil {
		return nil, nil
	}

	bytes, err := json.Marshal(token)
	if err != nil {
		a.logger.WithMethod(ctx, "MarshalJSON").Error("failed to marshal token", zap.Error(err))
		return nil, err
	}

	res := string(bytes)
	return &res, nil
}

func (a *authTokenAdapter) UnmarshalJSON(ctx context.Context, token *string) (*model.AuthToken, error) {
	if token == nil || *token == "" {
		return nil, nil
	}

	var res model.AuthToken
	err := json.Unmarshal([]byte(*token), &res)
	if err != nil {
		a.logger.WithMethod(ctx, "UnmarshalJSON").Error("failed to unmarshal token", zap.Error(err))
		return nil, err
	}

	return &res, nil
}

func (a *authTokenAdapter) ToProto(token *model.AuthToken) *pb.AuthOut {
	return &pb.AuthOut{
		ExpiresIn: token.ExpiresIn,
		Token:     token.Token,
	}
}

func (a *authTokenAdapter) FromUserID(ctx context.Context, userID int64) (*model.AuthToken, error) {
	logger := a.logger.WithMethod(ctx, "FromUserID")

	u, err := uuid.NewV4()
	if err != nil {
		logger.Error("failed to create new uuidv4", zap.Error(err))
		return nil, err
	}

	expiresIn := time.Now().Add(time.Hour * 3).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uuid":      u.String(),
			"user_id":   userID,
			"expiresIn": expiresIn,
		})

	t, err := token.SignedString([]byte("@TODO"))
	if err != nil {
		logger.Error("failed to sign jwt", zap.Error(err))
		return nil, err
	}

	return &model.AuthToken{
		UUID:      u,
		Token:     t,
		UserID:    userID,
		ExpiresIn: expiresIn,
	}, nil
}

func (a *authTokenAdapter) FromRegisterProtoToUser(in *pb.RegisterIn) *model.User {
	return &model.User{
		Email:    in.Email,
		Password: in.Password,
		Name:     in.Name,
		LastName: in.LastName,
	}
}
