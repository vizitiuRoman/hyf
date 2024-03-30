package adapter

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/config"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain"
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

		accessSecretKey:  f.cfg.Auth.AccessTokenSecretKey,
		refreshSecretKey: f.cfg.Auth.RefreshTokenSecretKey,
	}
}

type authTokenAdapter struct {
	logger log.Logger

	accessSecretKey  string
	refreshSecretKey string
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

func (a *authTokenAdapter) ToProto(accessToken, refreshToken *model.AuthToken) *pb.AuthOut {
	return &pb.AuthOut{
		ExpiresIn:    accessToken.ExpiresIn,
		Token:        accessToken.Token,
		RefreshToken: refreshToken.Token,
	}
}

func (a *authTokenAdapter) FromUserID(ctx context.Context, userID int64) (*model.AuthToken, *model.AuthToken, error) {
	logger := a.logger.WithMethod(ctx, "AccessTokenFromUserID")

	// ACCESS TOKEN
	u, err := uuid.NewV4()
	if err != nil {
		logger.Error("failed to create new uuidv4", zap.Error(err))
		return nil, nil, err
	}

	expiresIn := time.Now().Add(time.Hour * 3).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uuid":      u.String(),
			"user_id":   userID,
			"expiresIn": expiresIn,
		})

	t, err := token.SignedString([]byte(a.accessSecretKey))
	if err != nil {
		logger.Error("failed to sign jwt", zap.Error(err))
		return nil, nil, err
	}
	accessToken := model.AuthToken{
		UUID:      u,
		Token:     t,
		UserID:    userID,
		ExpiresIn: expiresIn,
	}

	// REFRESH TOKEN
	u, err = uuid.NewV4()
	if err != nil {
		logger.Error("failed to create new uuidv4", zap.Error(err))
		return nil, nil, err
	}

	expiresIn = time.Now().Add(time.Hour * 24).Unix()

	token = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uuid":      u.String(),
			"user_id":   userID,
			"expiresIn": expiresIn,
		})

	t, err = token.SignedString([]byte(a.refreshSecretKey))
	if err != nil {
		logger.Error("failed to sign jwt", zap.Error(err))
		return nil, nil, err
	}
	refreshToken := model.AuthToken{
		UUID:      u,
		Token:     t,
		UserID:    userID,
		ExpiresIn: expiresIn,
	}

	return &accessToken, &refreshToken, nil
}

func (a *authTokenAdapter) ClaimsFromJWT(ctx context.Context, token string) (string, int64, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.accessSecretKey), nil
	})
	if err != nil {
		return "", 0, domain.ErrUnauthorized
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", 0, domain.ErrUnauthorized
	}

	return claims["uuid"].(string), int64(claims["user_id"].(float64)), nil
}

func (a *authTokenAdapter) ClaimsFromRefreshJWT(ctx context.Context, token string) (string, int64, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.refreshSecretKey), nil
	})
	if err != nil {
		return "", 0, domain.ErrUnauthorized
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", 0, domain.ErrUnauthorized
	}

	return claims["uuid"].(string), int64(claims["user_id"].(float64)), nil
}

func (a *authTokenAdapter) FromRegisterProtoToUser(in *pb.RegisterIn) *model.User {
	return &model.User{
		Email:    in.Email,
		Password: in.Password,
		Name:     in.Name,
		LastName: in.LastName,
	}
}
