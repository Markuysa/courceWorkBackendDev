package usecase

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/client/repository"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
)

type UC struct {
	cfg       config.Config
	adminRepo repository.Repository
}

func New(
	cfg config.Config,
	adminRepo repository.Repository,
) Usecase {
	return &UC{
		cfg:       cfg,
		adminRepo: adminRepo,
	}
}

func (uc *UC) AssignTask(ctx context.Context, request models.AssignTaskRequest) (response models.AssignTaskResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AssignTask")
	defer span.End()

	return
}

func (uc *UC) CreateTask(ctx context.Context, request models.CreateTaskRequest) (response models.CreateTaskResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AssignTask")
	defer span.End()

	return
}

func (uc *UC) DeleteTask(ctx context.Context, request models.DeleteTaskRequest) (response models.DeleteTaskResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AssignTask")
	defer span.End()

	return

}
