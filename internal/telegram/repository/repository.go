package repository

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/Markuysa/courceWorkBackendDev/utils/pgconn"
)

type TgRepository struct {
	db *pgconn.Connector
}

func New(db *pgconn.Connector) Repository {
	return &TgRepository{db: db}
}

func (t *TgRepository) LinkTelegram(
	ctx context.Context,
	userID int,
	tgChat int,
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

func (t *TgRepository) GetNotificationMessages(ctx context.Context) (messages map[int][]models.TaskTgItem, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetNotificationMessages")
	defer span.End()

	rows, err := t.db.QueryContext(
		ctx,
		queryGetUsersTask,
	)
	if err != nil {
		return messages, err
	}

	messages = make(map[int][]models.TaskTgItem)

	for rows.Next() {
		var task models.TaskTgItem
		var chatID int

		err = rows.Scan(
			&chatID,
			&task.Description,
			&task.Deadline,
		)
		if err != nil {
			return messages, err
		}

		messages[chatID] = append(messages[chatID], task)
	}

	return messages, err
}
