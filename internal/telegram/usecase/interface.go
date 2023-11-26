package usecase

import (
	"context"
)

type Usecase interface {
	ListenMessages(ctx context.Context) (err error)
	SendTaskNotifications(ctx context.Context) (err error)
	StartWorker(ctx context.Context)
}
