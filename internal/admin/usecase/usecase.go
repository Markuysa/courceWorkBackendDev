package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/admin/repository"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"gopkg.in/guregu/null.v3"
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

	err = uc.adminRepo.AddTask(ctx, models.Task{
		Category: null.Int{
			NullInt64: sql.NullInt64{
				Int64: request.Category,
				Valid: request.Category != 0,
			},
		},
		Deadline: null.Time{
			Time:  time.Unix(request.Deadline, 0),
			Valid: request.Deadline != 0,
		},
		Status: null.Int{
			NullInt64: sql.NullInt64{
				Int64: request.Status,
				Valid: request.Status != 0,
			},
		},
		Priority: null.Int{
			NullInt64: sql.NullInt64{
				Int64: request.Priority,
				Valid: request.Priority != 0,
			},
		},
		CreatorID:   request.Creator,
		Description: request.Description,
		ParticipantID: null.Int{
			NullInt64: sql.NullInt64{
				Int64: request.ParticipantID,
				Valid: request.ParticipantID != 0,
			},
		},
	})
	if err != nil {
		return models.CreateTaskResponse{
			FailCause: err.Error(),
		}, err
	}

	return models.CreateTaskResponse{
		Success: true,
	}, err
}

func (uc *UC) DeleteTask(ctx context.Context, request models.DeleteTaskRequest) (response models.DeleteTaskResponse, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AssignTask")
	defer span.End()

	err = uc.adminRepo.DeleteTask(ctx, request.TaskID)
	if err != nil {
		return models.DeleteTaskResponse{
			FailCause: err.Error(),
		}, err
	}

	return models.DeleteTaskResponse{
		Success: true,
	}, err
}

func (uc *UC) GetUsersTaskList(ctx context.Context, filters models.TasksFilters) (tasks []models.TaskItem, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetUsersTaskList")
	defer span.End()

	tasks, err = uc.adminRepo.GetTasks(ctx, filters)
	if err != nil {
		return tasks, err
	}

	return tasks, err
}
