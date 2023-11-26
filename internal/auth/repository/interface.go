package repository

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Repository interface {
	SaveOTPSecret(ctx context.Context, saveOTPParams models.SaveOTPRequest) (err error)
	GetOTPSecret(ctx context.Context, getOTPParams models.GetOTPRequest) (secret string, err error)
	GetUserByUsername(ctx context.Context, getUserParams models.GetUserRequest) (user models.User, err error)
	SaveUser(ctx context.Context, user models.User) (err error)
	SaveAdmin(ctx context.Context, admin models.User) (err error)
	GetAdminOTPSecret(ctx context.Context, request models.GetOTPRequest) (secret string, err error)
	GetAdminByUsername(ctx context.Context, request models.GetUserRequest) (user models.User, err error)
}
