package http

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	UpdateTask(ctx *fiber.Ctx) error
	AddComment(ctx *fiber.Ctx) error
	GetTasksList(ctx *fiber.Ctx) error
	LinkTelegram(ctx *fiber.Ctx) error

	GetStatusList(ctx *fiber.Ctx) error
	GetPriorityList(ctx *fiber.Ctx) error
	GetCategoryList(ctx *fiber.Ctx) error
}
