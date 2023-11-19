package repository

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Repository interface {
	GetTasksList(ctx context.Context, userID int) (tasks []models.Task, err error)
	UpdateTask(ctx context.Context, task models.UpdateTask) (err error)
}
