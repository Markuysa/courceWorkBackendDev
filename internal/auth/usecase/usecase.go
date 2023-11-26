package usecase

import (
	"bytes"
	"context"
	"image/png"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/auth/cache"
	"github.com/Markuysa/courceWorkBackendDev/internal/auth/repository"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/pkg/lists"
	"github.com/Markuysa/courceWorkBackendDev/utils/coder"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
)

type UC struct {
	cfg          config.Config
	authRepo     repository.Repository
	sessionCache cache.Cache
}

func New(
	cfg config.Config,
	sessionCache cache.Cache,
	authRepo repository.Repository,
) Usecase {
	return &UC{
		cfg:          cfg,
		sessionCache: sessionCache,
		authRepo:     authRepo,
	}
}

func (uc *UC) GenerateOTP(
	ctx context.Context,
	request models.GenerateOTPRequest,
) (response models.GenerateOTPResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GenerateOTP")
	defer span.End()

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      request.ServiceName,
		AccountName: request.Username,
	})
	if err != nil {
		return models.GenerateOTPResponse{
			FailCause: err.Error(),
		}, err
	}

	buf := bytes.Buffer{}

	img, err := key.Image(200, 200)
	if err != nil {
		return models.GenerateOTPResponse{
			FailCause: err.Error(),
		}, err
	}

	err = png.Encode(&buf, img)
	if err != nil {
		return models.GenerateOTPResponse{
			FailCause: err.Error(),
		}, err
	}

	return models.GenerateOTPResponse{
		QR:     buf.Bytes(),
		Secret: key.Secret(),
	}, err
}

func (uc *UC) ValidateOTP(
	ctx context.Context,
	request models.ValidateOTPRequest,
) (response models.ValidateOTPResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "ValidateOTP")
	defer span.End()

	secret, err := uc.authRepo.GetOTPSecret(ctx, models.GetOTPRequest{
		Username: request.Username,
	})
	if err != nil {
		return response, err
	}

	isValid := totp.Validate(request.PassCode, secret)

	return models.ValidateOTPResponse{
		Passed: isValid,
	}, err
}

func (uc *UC) PrepareClientSignIn(
	ctx context.Context,
	request models.PrepareSignInRequest,
) (response models.PrepareSignInResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "PrepareClientSignIn")
	defer span.End()

	user, err := uc.authRepo.GetUserByUsername(ctx, models.GetUserRequest{
		Username: request.Username,
	})
	if err != nil {
		return models.PrepareSignInResponse{
			FailCause: err.Error(),
		}, err
	}

	valid := coder.CheckPassword(user.Password, request.Password)

	if !valid {
		return models.PrepareSignInResponse{
			FailCause: "invalid password",
		}, err
	}

	key := uuid.New().String()

	err = uc.sessionCache.SaveAccessKey(ctx, models.SaveAccessKeyRequest{
		Key:      key,
		Username: user.Username,
	})
	if err != nil {
		return response, err
	}

	return models.PrepareSignInResponse{
		AccessToken: key,
	}, nil
}

func (uc *UC) FinalizeClientSignIn(
	ctx context.Context,
	request models.FinalizeSignInRequest,
) (response models.FinalizeSignInResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "FinalizeClientSignIn")
	defer span.End()

	valid, err := uc.sessionCache.ValidateAccessKey(ctx, models.ValidateAccessKeyRequest{
		Key:      request.AccessKey,
		Username: request.Username,
	})
	if err != nil {
		return models.FinalizeSignInResponse{
			FailCause: err.Error(),
		}, err
	}

	if !valid {
		return models.FinalizeSignInResponse{
			FailCause: "invalid access token",
		}, err
	}

	user, err := uc.authRepo.GetUserByUsername(ctx, models.GetUserRequest{
		Username: request.Username,
	})
	if err != nil {
		return models.FinalizeSignInResponse{
			FailCause: err.Error(),
		}, err
	}

	validateOTP, err := uc.ValidateOTP(ctx, models.ValidateOTPRequest{
		Username: request.Username,
		PassCode: request.PassCode,
	})
	if err != nil {
		return models.FinalizeSignInResponse{
			FailCause: err.Error(),
		}, err
	}

	if !validateOTP.Passed {
		return models.FinalizeSignInResponse{
			FailCause: "invalid otp",
		}, err
	}

	sessionKey := uuid.New().String()

	err = uc.sessionCache.SaveSession(ctx, models.SaveSessionRequest{
		UserID:     user.ID,
		Username:   user.Username,
		SessionKey: sessionKey,
		Role:       lists.RoleUser,
	})
	if err != nil {
		return models.FinalizeSignInResponse{
			FailCause: err.Error(),
		}, err
	}

	return models.FinalizeSignInResponse{
		SessionKey: sessionKey,
	}, err
}

func (uc *UC) AdminSignIn(
	ctx context.Context,
	request models.AdminSignInRequest,
) (response models.AdminSignInResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AdminSignIn")
	defer span.End()

	user, err := uc.authRepo.GetAdminByUsername(ctx, models.GetUserRequest{
		Username: request.Username,
	})
	if err != nil {
		return models.AdminSignInResponse{
			FailCause: err.Error(),
		}, err
	}

	valid := coder.CheckPassword(user.Password, request.Password)

	if !valid {
		return models.AdminSignInResponse{
			FailCause: "invalid password",
		}, err
	}

	otpResp, err := uc.ValidateAdminOTP(ctx, models.ValidateOTPRequest{
		Username: request.Username,
		PassCode: request.PassCode,
	})
	if err != nil {
		return models.AdminSignInResponse{
			FailCause: err.Error(),
		}, err
	}

	if !otpResp.Passed {
		return models.AdminSignInResponse{
			FailCause: "invalid otp",
		}, err
	}

	key := uuid.New().String()

	err = uc.sessionCache.SaveSession(ctx, models.SaveSessionRequest{
		UserID:     user.ID,
		Username:   user.Username,
		SessionKey: key,
		Role:       lists.RoleAdmin,
	})
	if err != nil {
		return response, err
	}

	return models.AdminSignInResponse{
		SessionKey: key,
	}, nil
}

func (uc *UC) ClientSignUP(ctx context.Context, request models.ClientSignUpRequest) (response models.ClientSignUpResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "ClientSignUP")
	defer span.End()

	otpResp, err := uc.GenerateOTP(ctx, models.GenerateOTPRequest{
		Username:    request.Username,
		ServiceName: uc.cfg.ServiceName,
	})
	if err != nil {
		return models.ClientSignUpResponse{
			FailCause: otpResp.FailCause,
		}, err
	}

	password, err := coder.HashPassword(request.Password)
	if err != nil {
		return models.ClientSignUpResponse{
			FailCause: err.Error(),
		}, err
	}

	err = uc.authRepo.SaveUser(ctx, models.User{
		Username:  request.Username,
		Password:  password,
		OtpSecret: otpResp.Secret,
	})
	if err != nil {
		return models.ClientSignUpResponse{
			FailCause: err.Error(),
		}, err
	}

	return models.ClientSignUpResponse{
		Success: true,
		QR:      otpResp.QR,
	}, err
}

func (uc *UC) AdminSignUP(ctx context.Context, request models.AdminSignUpRequest) (response models.AdminSignUpResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AdminSignUP")
	defer span.End()

	otpResp, err := uc.GenerateOTP(ctx, models.GenerateOTPRequest{
		Username:    request.Username,
		ServiceName: uc.cfg.AdminServiceName,
	})
	if err != nil {
		return models.AdminSignUpResponse{
			FailCause: otpResp.FailCause,
		}, err
	}

	password, err := coder.HashPassword(request.Password)
	if err != nil {
		return models.AdminSignUpResponse{
			FailCause: err.Error(),
		}, err
	}

	err = uc.authRepo.SaveAdmin(ctx, models.User{
		Username:  request.Username,
		Password:  password,
		OtpSecret: otpResp.Secret,
	})
	if err != nil {
		return models.AdminSignUpResponse{
			FailCause: err.Error(),
		}, err
	}

	return models.AdminSignUpResponse{
		Success: true,
		QR:      otpResp.QR,
	}, err
}

func (uc *UC) ValidateAdminOTP(ctx context.Context, request models.ValidateOTPRequest) (response models.ValidateOTPResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "ValidateAdminOTP")
	defer span.End()

	secret, err := uc.authRepo.GetAdminOTPSecret(ctx, models.GetOTPRequest{
		Username: request.Username,
	})
	if err != nil {
		return response, err
	}

	isValid := totp.Validate(request.PassCode, secret)

	return models.ValidateOTPResponse{
		Passed: isValid,
	}, err
}
