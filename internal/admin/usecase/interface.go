package usecase

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Usecase interface {
	AssignTask(ctx context.Context, request models.AssignTaskRequest) (response models.AssignTaskResponse, err error)
	CreateTask(ctx context.Context, request models.CreateTaskRequest) (response models.CreateTaskResponse, err error)
	DeleteTask(ctx context.Context, request models.DeleteTaskRequest) (response models.DeleteTaskResponse, err error)
	GetUsersTaskList(ctx context.Context, filters models.TasksFilters) (tasks []models.Task, err error)
}
