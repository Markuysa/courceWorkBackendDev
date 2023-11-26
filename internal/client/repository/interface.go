package repository

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Repository interface {
	GetTasksList(ctx context.Context, userID int) (tasks []models.Task, err error)
	UpdateTask(ctx context.Context, task models.UpdateTask) (err error)
	LinkTelegram(ctx context.Context, userID int, tgChat string) (err error)

	GetStatusList(ctx context.Context) (tasks []models.StatusListItem, err error)
	GetPriorityList(ctx context.Context) (tasks []models.PriorityListItem, err error)
	GetCategoryList(ctx context.Context) (categories []models.CategoryListItem, err error)
}
