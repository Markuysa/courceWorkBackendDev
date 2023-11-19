package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/auth/usecase"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	GenerateOTP(ctx *fiber.Ctx) error
	ValidateOTP(ctx *fiber.Ctx) error
}

type AuthHandlers struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handlers {
	return &AuthHandlers{
		uc: uc,
	}
}

func (a AuthHandlers) GenerateOTP(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "GenerateOTP")
	defer span.End()

	in := models.GenerateOTPRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.GenerateOTP(ctx, in)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}

func (a AuthHandlers) ValidateOTP(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "ValidateOTP")
	defer span.End()

	in := models.ValidateOTPRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.ValidateOTP(ctx, in)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}
