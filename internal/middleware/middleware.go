package middleware

import "github.com/gofiber/fiber/v2"

type Mw struct {
}

func New() *Mw {
	return &Mw{}
}

func (mw *Mw) AdminAuth(c *fiber.Ctx) error {

	return c.Next()
}

func (mw *Mw) ClientAuth(c *fiber.Ctx) error {

	return c.Next()
}

func (mw *Mw) DefaultAuth(c *fiber.Ctx) error {

	return c.Next()
}
