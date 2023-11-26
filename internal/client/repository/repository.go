package repository

import (
	"context"
	"time"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/Markuysa/courceWorkBackendDev/utils/pgconn"
)

type TaskRepository struct {
	db *pgconn.Connector
}

func New(
	db *pgconn.Connector,
) Repository {

	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) GetTasksList(ctx context.Context, userID int) (tasks []models.Task, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetTasksList")
	defer span.End()

	err = t.db.SelectContext(
		ctx,
		&tasks,
		queryGetTasksList,
		userID,
	)
	if err != nil {
		return tasks, err
	}

	return tasks, err
}

func (t *TaskRepository) GetStatusList(ctx context.Context) (status []models.StatusListItem, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetTasksList")
	defer span.End()

	err = t.db.SelectContext(
		ctx,
		&status,
		queryGetStatusList,
	)
	if err != nil {
		return status, err
	}

	return status, err
}

func (t *TaskRepository) GetCategoryList(ctx context.Context) (categories []models.CategoryListItem, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetTasksList")
	defer span.End()

	err = t.db.SelectContext(
		ctx,
		&categories,
		queryGetCategoriesList,
	)
	if err != nil {
		return categories, err
	}

	return categories, err
}

func (t *TaskRepository) GetPriorityList(ctx context.Context) (priority []models.PriorityListItem, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetTasksList")
	defer span.End()

	err = t.db.SelectContext(
		ctx,
		&priority,
		queryGetPriorityList,
	)
	if err != nil {
		return priority, err
	}

	return priority, err
}

func (t *TaskRepository) UpdateTask(ctx context.Context, task models.UpdateTask) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "UpdateTask")
	defer span.End()

	_, err = t.db.ExecContext(
		ctx,
		queryUpdateTask,
		time.Unix(task.Deadline, 0),
		task.ID,
	)
	if err != nil {
		return err
	}

	return err
}

func (t *TaskRepository) LinkTelegram(
	ctx context.Context,
	userID int,
	tgChat string,
) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "LinkTelegram")
	defer span.End()

	_, err = t.db.ExecContext(
		ctx,
		queryLinkTelegram,
		tgChat,
		userID,
	)
	if err != nil {
		return err
	}

	return err
}
