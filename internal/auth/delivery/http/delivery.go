package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(route fiber.Router, mw middleware.Middleware, handlers Handlers) {
	route.Post("/generate_otp", mw.DefaultAuth, handlers.GenerateOTP)
	route.Post("/validate_otp", mw.DefaultAuth, handlers.ValidateOTP)
	route.Post("/sign_up", mw.DefaultAuth, handlers.ClientSignUP)

	route.Post("/prepare_sign_in", mw.DefaultAuth, handlers.PrepareClientSignIn)
	route.Post("/finalize_sign_in", mw.DefaultAuth, handlers.FinalizeClientSignIn)

	adminRoute := route.Group("/admin")
	{
		adminRoute.Post("/sign_in", mw.DefaultAuth, handlers.AdminSignIn)
		adminRoute.Post("/sign_up", mw.DefaultAuth, handlers.AdminSignUp)
	}
}
