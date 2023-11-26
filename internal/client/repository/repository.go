package repository

import (
	"context"
	"errors"

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

func (t *TaskRepository) UpdateTask(ctx context.Context, task models.Task) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "UpdateTask")
	defer span.End()

	rows, err := t.db.ExecContext(
		ctx,
		queryUpdateTask,
		task.Deadline,
		task.ParticipantID,
		task.Description,
		task.Status,
		task.Category,
		task.Priority,
		task.ID,
	)
	if err != nil {
		return err
	}

	count, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("invalid rows count")
	}

	return nil
}

func (t *TaskRepository) AddComment(
	ctx context.Context,
	comment []byte,
	taskID int,
) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AddComment")
	defer span.End()

	_, err = t.db.ExecContext(
		ctx,
		queryAddComment,
		comment,
		taskID,
	)
	if err != nil {
		return err
	}

	return err
}
