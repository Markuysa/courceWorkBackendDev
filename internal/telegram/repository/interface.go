package repository

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

type Repository interface {
	LinkTelegram(
		ctx context.Context,
		userID int,
		tgChat int,
	) (err error)
	GetNotificationMessages(ctx context.Context) (messages map[int][]models.TaskTgItem, err error)
}
