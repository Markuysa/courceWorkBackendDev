package cache

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Cache interface {
	SaveAccessKey(ctx context.Context, saveAccessParams models.SaveAccessKeyRequest) (err error)
	ValidateAccessKey(ctx context.Context, request models.ValidateAccessKeyRequest) (isValid bool, err error)
	SaveSession(ctx context.Context, request models.SaveSessionRequest) (err error)
	GetSession(ctx context.Context, request models.GetSessionRequest) (err error)
}
