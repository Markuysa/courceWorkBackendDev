package http

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	GenerateOTP(ctx *fiber.Ctx) error
	ValidateOTP(ctx *fiber.Ctx) error
	ClientSignUP(ctx *fiber.Ctx) error
	PrepareClientSignIn(ctx *fiber.Ctx) error
	FinalizeClientSignIn(ctx *fiber.Ctx) error

	AdminSignIn(c *fiber.Ctx) error
	AdminSignUp(c *fiber.Ctx) error
}
