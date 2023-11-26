package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/auth/usecase"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/gofiber/fiber/v2"
)

type AuthHandlers struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handlers {
	return &AuthHandlers{
		uc: uc,
	}
}

func (a AuthHandlers) PrepareSignIn(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "ClientSignUP")
	defer span.End()

	in := models.PrepareSignInRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.PrepareSignIn(ctx, in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.JSON(response)
}

func (a AuthHandlers) FinalizeSignIn(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "ClientSignUP")
	defer span.End()

	in := models.FinalizeSignInRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.FinalizeSignIn(ctx, in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.JSON(response)
}

func (a AuthHandlers) ClientSignUP(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "ClientSignUP")
	defer span.End()

	in := models.SignUpRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.SignUp(ctx, in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.JSON(response)
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
		return c.Status(fiber.StatusInternalServerError).JSON(response)
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
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.JSON(response)
}
