package repository

import (
	"context"
	"time"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/Markuysa/courceWorkBackendDev/utils/pgconnector"
)

type TaskRepository struct {
	db *pgconnector.Connector
}

func New(
	db *pgconnector.Connector,
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

	return tasks, err
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

	return err
}
