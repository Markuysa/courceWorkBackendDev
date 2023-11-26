package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapAdminRoutes(route fiber.Router, mw middleware.Middleware, handlers Handlers) {
	route.Group("/tasks")
	{
		route.Post("/create", mw.AdminAuth, handlers.CreateTask)
		route.Post("/list", mw.AdminAuth, handlers.GetUsersTaskList)
		route.Post("/delete", mw.AdminAuth, handlers.DeleteTask)
	}
}
