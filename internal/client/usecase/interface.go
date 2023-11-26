package usecase

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Usecase interface {
	LinkTG(ctx context.Context, request models.LinkTgRequest) (response models.LinkTgResponse, err error)
	ShowTasksList(ctx context.Context, request models.ShowTasksListRequest) (response models.ShowTasksListResponse, err error)
	UpdateTask(ctx context.Context, request models.UpdateTask) (response models.UpdateTaskResponse, err error)
}
