package usecase

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Usecase interface {
	GenerateOTP(ctx context.Context, request models.GenerateOTPRequest) (response models.GenerateOTPResponse, err error)
	ValidateOTP(ctx context.Context, request models.ValidateOTPRequest) (response models.ValidateOTPResponse, err error)

	PrepareClientSignIn(
		ctx context.Context,
		request models.PrepareSignInRequest,
	) (response models.PrepareSignInResponse, err error)
	FinalizeClientSignIn(
		ctx context.Context,
		request models.FinalizeSignInRequest,
	) (response models.FinalizeSignInResponse, err error)

	AdminSignIn(
		ctx context.Context,
		request models.AdminSignInRequest,
	) (response models.AdminSignInResponse, err error)

	ClientSignUP(ctx context.Context, request models.ClientSignUpRequest) (response models.ClientSignUpResponse, err error)
	AdminSignUP(ctx context.Context, request models.AdminSignUpRequest) (response models.AdminSignUpResponse, err error)
	ValidateAdminOTP(ctx context.Context, request models.ValidateOTPRequest) (response models.ValidateOTPResponse, err error)
}
