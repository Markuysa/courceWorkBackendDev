package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	AdminAuth(c *fiber.Ctx) error
	ClientAuth(c *fiber.Ctx) error
	DefaultAuth(c *fiber.Ctx) error
}
