package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(route fiber.Router, mw middleware.Middleware, handlers Handlers) {
	route.Post("/generate_otp", mw.DefaultAuth, handlers.GenerateOTP)
	route.Post("/validate_otp", mw.DefaultAuth, handlers.ValidateOTP)
}
