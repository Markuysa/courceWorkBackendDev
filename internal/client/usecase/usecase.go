package usecase

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/client/repository"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
)

type UC struct {
	cfg config.Config

	clientRepo repository.Repository
}

func New(
	cfg config.Config,

	clientRepo repository.Repository,
) Usecase {
	return &UC{
		cfg: cfg,

		clientRepo: clientRepo,
	}
}

func (uc *UC) ShowTasksList(ctx context.Context, request models.ShowTasksListRequest) (response models.ShowTasksListResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "ShowTasksList")
	defer span.End()

	return
}

func (uc *UC) MoveTask(ctx context.Context, request models.MoveTaskRequest) (response models.MoveTaskResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "MoveTask")
	defer span.End()

	return
}

func (uc *UC) LinkTG(ctx context.Context, request models.LinkTgRequest) (response models.LinkTgResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "LinkTG")
	defer span.End()

	return
}
