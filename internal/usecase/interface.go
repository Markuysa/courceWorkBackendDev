package usecase

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Usecase interface {
	Auth
}

type Auth interface {
	GenerateOTP(ctx context.Context, request models.GenerateOTPRequest) (response models.GenerateOTPResponse, err error)
	ValidateOTP(ctx context.Context, request models.ValidateOTPRequest) (response models.ValidateOTPResponse, err error)
}
