package models

import (
	"database/sql"
)

type (
	TgTaskMessage struct {
		ChatID int64
		Tasks  []TaskTgItem
	}
	TaskTgItem struct {
		Description string       `db:"description"`
		Deadline    sql.NullTime `db:"deadline"`
	}
)
