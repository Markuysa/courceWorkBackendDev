package usecase

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Usecase interface {
	GenerateOTP(ctx context.Context, request models.GenerateOTPRequest) (response models.GenerateOTPResponse, err error)
	ValidateOTP(ctx context.Context, request models.ValidateOTPRequest) (response models.ValidateOTPResponse, err error)
	PrepareSignIn(
		ctx context.Context,
		request models.PrepareSignInRequest,
	) (response models.PrepareSignInResponse, err error)
	FinalizeSignIn(
		ctx context.Context,
		request models.FinalizeSignInRequest,
	) (response models.FinalizeSignInResponse, err error)
	SignUp(ctx context.Context, request models.SignUpRequest) (response models.SignUpResponse, err error)
}
