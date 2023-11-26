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

func (t TaskRepository) SaveOTPSecret(ctx context.Context, saveOTPParams models.SaveOTPRequest) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "SaveOTPSecret")
	defer span.End()

	_, err = t.db.ExecContext(
		ctx,
		querySaveOTPSecret,
		saveOTPParams.Secret,
		saveOTPParams.Username,
	)
	if err != nil {
		return err
	}

	return err
}

func (t TaskRepository) GetOTPSecret(ctx context.Context, getOTPParams models.GetOTPRequest) (secret string, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetOTPSecret")
	defer span.End()

	err = t.db.GetContext(
		ctx,
		&secret,
		queryGetOTP,
		getOTPParams.Username,
	)
	if err != nil {
		return secret, err
	}

	return secret, err
}

func (t TaskRepository) GetUserByUsername(
	ctx context.Context,
	getUserParams models.GetUserRequest,
) (user models.User, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetUserByUsername")
	defer span.End()

	err = t.db.GetContext(
		ctx,
		&user,
		queryGetUser,
		getUserParams.Username,
	)
	if err != nil {
		return user, err
	}

	return user, err
}

func (t TaskRepository) SaveUser(ctx context.Context, user models.User) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "SaveUser")
	defer span.End()

	_, err = t.db.ExecContext(
		ctx,
		querySaveUser,
		user.Username,
		user.Password,
		user.OtpSecret,
	)
	if err != nil {
		return err
	}

	return err
}

func (t TaskRepository) SaveAdmin(ctx context.Context, admin models.User) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "SaveUser")
	defer span.End()

	_, err = t.db.ExecContext(
		ctx,
		querySaveAdmin,
		admin.Username,
		admin.Password,
		admin.OtpSecret,
	)
	if err != nil {
		return err
	}

	return err
}

func (t TaskRepository) GetAdminOTPSecret(ctx context.Context, request models.GetOTPRequest) (secret string, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetOTPSecret")
	defer span.End()

	err = t.db.GetContext(
		ctx,
		&secret,
		queryGetAdminOTP,
		request.Username,
	)
	if err != nil {
		return secret, err
	}

	return secret, err
}

func (t TaskRepository) GetAdminByUsername(ctx context.Context, request models.GetUserRequest) (user models.User, err error) {
	ctx, span := oteltrace.NewSpan(ctx, "GetAdminByUsername")
	defer span.End()

	err = t.db.GetContext(
		ctx,
		&user,
		queryGetAdmin,
		request.Username,
	)
	if err != nil {
		return user, err
	}

	return user, err
}
