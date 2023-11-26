package repository

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Repository interface {
	AssignTask(ctx context.Context, taskID string, assigneeID int) (err error)
	AddTask(ctx context.Context, task models.Task) (err error)
	DeleteTask(ctx context.Context, taskID string) (err error)
	GetTasks(ctx context.Context, filters models.TasksFilters) (tasks []models.TaskItem, err error)
}
