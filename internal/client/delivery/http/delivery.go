package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapClientRoutes(route fiber.Router, mw middleware.Middleware, handlers Handlers) {
	tasks := route.Group("tasks")
	{
		tasks.Post("/list", mw.ClientAuth, handlers.GetTasksList)
		tasks.Post("/update", mw.ClientAuth, handlers.UpdateTask)
		tasks.Post("/comment", mw.ClientAuth, handlers.AddComment)
	}
	lists := route.Group("/lists")
	{
		lists.Get("/status", mw.DefaultAuth, handlers.GetStatusList)
		lists.Get("/priority", mw.DefaultAuth, handlers.GetPriorityList)
		lists.Get("/category", mw.DefaultAuth, handlers.GetCategoryList)
	}
	route.Post("link_tg", mw.ClientAuth, handlers.LinkTelegram)
}
