package usecase

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Usecase interface {
	LinkTG(ctx context.Context, request models.LinkTgRequest) (response models.LinkTgResponse, err error)
	ShowTasksList(ctx context.Context, request models.ShowTasksListRequest) (response models.ShowTasksListResponse, err error)
	UpdateTask(ctx context.Context, request models.UpdateTask) (response models.UpdateTaskResponse, err error)

	GetPriorityList(ctx context.Context) (priority []models.PriorityListItem, err error)
	GetStatusList(ctx context.Context) (status []models.StatusListItem, err error)
	GetCategoryList(ctx context.Context) (category []models.CategoryListItem, err error)
}
