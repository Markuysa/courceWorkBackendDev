package http

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	GenerateOTP(ctx *fiber.Ctx) error
	ValidateOTP(ctx *fiber.Ctx) error
	ClientSignUP(ctx *fiber.Ctx) error
	PrepareSignIn(ctx *fiber.Ctx) error
	FinalizeSignIn(ctx *fiber.Ctx) error
}
