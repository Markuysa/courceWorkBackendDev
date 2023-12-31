package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/admin/usecase"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/pkg/constants"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type AdminHandlers struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handlers {
	return &AdminHandlers{
		uc: uc,
	}
}

func (a AdminHandlers) AssignTask(c *fiber.Ctx) error {
	//ctx, span := oteltrace.NewFiberSpan(c, "AssignTask")
	//defer span.End()

	return nil
}

func (a AdminHandlers) GetUsersTaskList(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "GetUsersTaskList")
	defer span.End()

	in := models.TasksFilters{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	tasks, err := a.uc.GetUsersTaskList(ctx, in)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(tasks)
}

func (a AdminHandlers) CreateTask(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "CreateTask")
	defer span.End()

	in := models.TaskModel{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	if in.Description == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.CreateTaskResponse{
			FailCause: "empty description",
		})
	}

	adminID, ok := c.Locals(constants.UserIDKey).(int)
	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	in.Creator = int64(adminID)

	response, err := a.uc.CreateTask(ctx, in)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.JSON(response)
}

func (a AdminHandlers) DeleteTask(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "DeleteTask")
	defer span.End()

	in := models.DeleteTaskRequest{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := a.uc.DeleteTask(ctx, in)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.JSON(response)
}
