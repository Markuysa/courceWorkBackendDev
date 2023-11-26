package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapAdminRoutes(route fiber.Router, mw middleware.Middleware, handlers Handlers) {
	tasks := route.Group("/tasks")
	{
		tasks.Post("/create", mw.AdminAuth, handlers.CreateTask)
		tasks.Post("/list", mw.AdminAuth, handlers.GetUsersTaskList)
		tasks.Post("/delete", mw.AdminAuth, handlers.DeleteTask)
	}
}
