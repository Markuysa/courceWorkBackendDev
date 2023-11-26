package http

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	UpdateTask(ctx *fiber.Ctx) error
	GetTasksList(ctx *fiber.Ctx) error
}
