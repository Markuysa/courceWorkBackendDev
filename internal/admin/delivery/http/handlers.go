package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/admin/usecase"
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	AssignTask(c *fiber.Ctx) error
	CreateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error
}

type AdminHandlers struct {
	uc usecase.Usecase
}

func (a AdminHandlers) AssignTask(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a AdminHandlers) CreateTask(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a AdminHandlers) DeleteTask(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func New(uc usecase.Usecase) Handlers {
	return &AdminHandlers{
		uc: uc,
	}
}
