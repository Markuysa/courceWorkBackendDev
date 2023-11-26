package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/pkg/constants"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/Markuysa/courceWorkBackendDev/utils/redisconnector"
	"github.com/redis/go-redis/v9"
)

type SessionCache struct {
	redis *redisconnector.Connector
	cfg   config.Config
}

func New(
	redis *redisconnector.Connector,
	cfg config.Config,
) Cache {
	return &SessionCache{
		cfg:   cfg,
		redis: redis,
	}
}

func (s SessionCache) SaveAccessKey(ctx context.Context, saveAccessParams models.SaveAccessKeyRequest) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "SaveAccessKey")
	defer span.End()

	return s.redis.Set(
		ctx,
		fmt.Sprintf(constants.AccessKeyf, saveAccessParams.Key),
		saveAccessParams.Username,
		s.cfg.Auth.AccessTTL.Duration,
	).Err()
}

func (s SessionCache) ValidateAccessKey(ctx context.Context, request models.ValidateAccessKeyRequest) (isValid bool, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "ValidateAccessKey")
	defer span.End()

	username, err := s.redis.Get(
		ctx,
		fmt.Sprintf(constants.AccessKeyf, request.Key),
	).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}

		return false, err
	}

	return username == request.Username, err
}

func (s SessionCache) SaveSession(ctx context.Context, request models.SaveSessionRequest) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "SaveSession")
	defer span.End()

	session, err := json.Marshal(models.Session{
		Role:     request.Role,
		UserID:   request.UserID,
		Username: request.Username,
	})
	if err != nil {
		return err
	}

	return s.redis.Set(
		ctx,
		fmt.Sprintf(constants.SessionKeyf, request.SessionKey),
		session,
		s.cfg.Auth.SessionTTL.Duration,
	).Err()
}

func (s SessionCache) GetSession(ctx context.Context, request models.GetSessionRequest) (session models.Session, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetSession")
	defer span.End()

	sessionByte, err := s.redis.Get(
		ctx,
		fmt.Sprintf(constants.SessionKeyf, request.SessionKey),
	).Bytes()
	if err != nil {
		return session, err
	}

	err = json.Unmarshal(sessionByte, &session)

	return session, err
}
