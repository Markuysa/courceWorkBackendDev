package repository

import (
	"context"

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

func (r *TaskRepository) AssignTask(ctx context.Context, taskID string, assigneeID int) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AssignTask")
	defer span.End()

	_, err = r.db.ExecContext(
		ctx,
		queryAssignTask,
		assigneeID,
		taskID,
	)
	if err != nil {
		return err
	}

	return err
}

func (r *TaskRepository) AddTask(ctx context.Context, task models.Task) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "AddTask")
	defer span.End()

	_, err = r.db.ExecContext(
		ctx,
		queryAddTask,
		task.Category,
		task.Deadline,
		task.Status,
		task.ParticipantID,
		task.CreatorID,
		task.Description,
		task.Priority,
	)
	if err != nil {
		return err
	}

	return err
}

func (r *TaskRepository) DeleteTask(ctx context.Context, taskID string) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "DeleteTask")
	defer span.End()

	return err
}

func (r *TaskRepository) GetTasks(ctx context.Context, filters models.TasksFilters) (tasks []models.Task, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetTasks")
	defer span.End()

	err = r.db.SelectContext(
		ctx,
		&tasks,
		queryGetUsersTasks,
		filters.UserID,
		filters.Limit,
		filters.Limit*(filters.Offset-1),
	)
	if err != nil {
		return tasks, err
	}

	return tasks, err
}
