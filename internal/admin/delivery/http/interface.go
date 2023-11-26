package http

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	AssignTask(c *fiber.Ctx) error
	CreateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error

	GetUsersTaskList(c *fiber.Ctx) error
}
