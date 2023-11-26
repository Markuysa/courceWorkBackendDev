package http

import (
	"time"

	"github.com/Markuysa/courceWorkBackendDev/internal/auth/usecase"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/pkg/constants"
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

func (a AuthHandlers) AdminSignIn(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "AdminSignIn")
	defer span.End()

	in := models.AdminSignInRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.AdminSignIn(ctx, in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	c.Cookie(&fiber.Cookie{
		Name:     constants.SessionKey,
		Value:    response.SessionKey,
		Expires:  time.Now().Add(time.Minute * 30),
		HTTPOnly: true,
	})

	return c.JSON(models.FinalizeClientSignIn{
		Success: err == nil,
	})
}

func (a AuthHandlers) PrepareClientSignIn(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "PrepareClientSignIn")
	defer span.End()

	in := models.PrepareSignInRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.PrepareClientSignIn(ctx, in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	c.Cookie(&fiber.Cookie{
		Name:     constants.AccessKey,
		Value:    response.AccessToken,
		Expires:  time.Now().Add(time.Minute * 5),
		HTTPOnly: true,
	})

	return c.JSON(models.FinalizeClientSignIn{
		Success: err == nil,
	})
}

func (a AuthHandlers) FinalizeClientSignIn(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "FinalizeClientSignIn")
	defer span.End()

	in := models.FinalizeSignInRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	accessKey := c.Cookies(constants.AccessKey)

	in.AccessKey = accessKey

	response, err := a.uc.FinalizeClientSignIn(ctx, in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	c.Cookie(&fiber.Cookie{
		Name:     constants.SessionKey,
		Value:    response.SessionKey,
		Expires:  time.Now().Add(time.Minute * 30),
		HTTPOnly: true,
	})

	return c.JSON(models.FinalizeClientSignIn{
		Success: err == nil,
	})
}

func (a AuthHandlers) ClientSignUP(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "ClientSignUP")
	defer span.End()

	in := models.ClientSignUpRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.ClientSignUP(ctx, in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.JSON(response)
}

func (a AuthHandlers) AdminSignUp(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "ClientSignUP")
	defer span.End()

	in := models.AdminSignUpRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.AdminSignUP(ctx, in)
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
