package repository

import (
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
