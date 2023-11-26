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

	tasks, err := uc.clientRepo.GetTasksList(ctx, request.UserID)
	if err != nil {
		return response, err
	}

	return models.ShowTasksListResponse{
		Tasks: tasks,
	}, err
}

func (uc *UC) GetStatusList(ctx context.Context) (response []models.StatusListItem, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetStatusList")
	defer span.End()

	response, err = uc.clientRepo.GetStatusList(ctx)
	if err != nil {
		return response, err
	}

	return response, err
}

func (uc *UC) GetCategoryList(ctx context.Context) (response []models.CategoryListItem, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetCategoryList")
	defer span.End()

	response, err = uc.clientRepo.GetCategoryList(ctx)
	if err != nil {
		return response, err
	}

	return response, err
}

func (uc *UC) GetPriorityList(ctx context.Context) (response []models.PriorityListItem, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetPriorityList")
	defer span.End()

	response, err = uc.clientRepo.GetPriorityList(ctx)
	if err != nil {
		return response, err
	}

	return response, err
}

func (uc *UC) UpdateTask(ctx context.Context, request models.UpdateTask) (response models.UpdateTaskResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "MoveTask")
	defer span.End()

	err = uc.clientRepo.UpdateTask(ctx, models.UpdateTask{
		ID:       request.ID,
		Deadline: request.Deadline,
	})
	if err != nil {
		return models.UpdateTaskResponse{
			Success:   err != nil,
			FailCause: err.Error(),
		}, err
	}

	return models.UpdateTaskResponse{
		Success: err != nil,
	}, err
}

func (uc *UC) LinkTG(ctx context.Context, request models.LinkTgRequest) (response models.LinkTgResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "LinkTG")
	defer span.End()

	err = uc.clientRepo.LinkTelegram(
		ctx,
		request.UserID,
		request.TgChat,
	)
	if err != nil {
		return models.LinkTgResponse{
			Success:   err != nil,
			FailCause: err.Error(),
		}, err
	}

	return models.LinkTgResponse{
		Success: err != nil,
	}, err
}
