package usecase

import (
	"bytes"
	"context"
	"image/png"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/pquerna/otp/totp"
)

func (uc *UC) GenerateOTP(ctx context.Context, request models.GenerateOTPRequest) (response models.GenerateOTPResponse, err error) {
	ctx, span := uc.tracer.Start(ctx, "GenerateOTP")
	defer span.End()

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      uc.cfg.ServiceName,
		AccountName: request.Username,
	})

	var buf bytes.Buffer
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

	response = models.GenerateOTPResponse{
		QR: buf.Bytes(),
	}

	// save here in cache key object

	return response, err
}

func (uc *UC) ValidateOTP(
	ctx context.Context,
	request models.ValidateOTPRequest,
) (response models.ValidateOTPResponse, err error) {
	ctx, span := uc.tracer.Start(ctx, "ValidateOTP")
	defer span.End()

	return response, err
}
