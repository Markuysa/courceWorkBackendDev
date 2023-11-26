package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapClientRoutes(route fiber.Router, mw middleware.Middleware, handlers Handlers) {
	route.Group("tasks")
	{
		route.Post("/list", mw.ClientAuth, handlers.GetTasksList)
		route.Post("/update", mw.ClientAuth, handlers.UpdateTask)
	}
}
