package http

import (
	"github.com/Markuysa/courceWorkBackendDev/internal/client/usecase"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/pkg/constants"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/gofiber/fiber/v2"
)

type ClientHandlers struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handlers {
	return &ClientHandlers{
		uc: uc,
	}
}

func (h *ClientHandlers) UpdateTask(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "UpdateTask")
	defer span.End()

	in := models.UpdateTask{}

	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := h.uc.UpdateTask(ctx, in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.JSON(response)
}

func (h *ClientHandlers) GetTasksList(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "GetTasksList")
	defer span.End()

	userID, ok := c.Locals(constants.UserIDKey).(int)
	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tasks, err := h.uc.ShowTasksList(ctx, models.ShowTasksListRequest{
		UserID: userID,
	})
	if err != nil {
		return err
	}

	return c.JSON(tasks)
}

func (h *ClientHandlers) GetStatusList(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "GetStatusList")
	defer span.End()

	response, err := h.uc.GetStatusList(ctx)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}

func (h *ClientHandlers) GetPriorityList(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "GetPriorityList")
	defer span.End()

	response, err := h.uc.GetPriorityList(ctx)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}

func (h *ClientHandlers) GetCategoryList(c *fiber.Ctx) error {
	ctx, span := oteltrace.NewFiberSpan(c, "GetCategoryList")
	defer span.End()

	response, err := h.uc.GetCategoryList(ctx)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}
