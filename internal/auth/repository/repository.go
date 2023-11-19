package repository

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/Markuysa/courceWorkBackendDev/utils/pgconnector"
)

type TaskRepository struct {
	db *pgconnector.Connector
}

func (t TaskRepository) SaveOTPSecret(ctx context.Context, saveOTPParams models.SaveOTPRequest) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "SaveOTPSecret")
	defer span.End()

	return err
}

func (t TaskRepository) GetOTPSecret(ctx context.Context, getOTPParams models.GetOTPRequest) (secret string, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetOTPSecret")
	defer span.End()

	return secret, err
}

func (t TaskRepository) GetUserByUsername(ctx context.Context, getUserParams models.GetUserRequest) (user models.User, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetUserByUsername")
	defer span.End()

	return user, err
}

func (t TaskRepository) SaveUser(ctx context.Context, user models.User) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "SaveUser")
	defer span.End()

	return err
}

func New(
	db *pgconnector.Connector,
) Repository {

	return &TaskRepository{
		db: db,
	}
}
