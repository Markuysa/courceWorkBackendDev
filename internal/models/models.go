package models

import (
	"gopkg.in/guregu/null.v3"
)

type (
	Task struct {
		ID            int64     `db:"id"` // require
		Category      null.Int  `db:"category"`
		Deadline      null.Time `db:"deadline"`
		Status        null.Int  `db:"status"`
		Comments      []byte    `db:"comments"`
		Priority      null.Int  `db:"priority"`
		CreatorID     int64     `db:"creator_id"`  // require
		Description   string    `db:"description"` // require
		ParticipantID null.Int  `db:"participant_id"`
	}
	TaskItem struct {
		ID            int64       `db:"id"` // require
		Category      null.String `db:"category"`
		Deadline      null.Int    `db:"deadline"`
		Status        null.String `db:"status"`
		Comments      []byte      `db:"comments"`
		Priority      null.String `db:"priority"`
		CreatorID     int64       `db:"creator_id"`  // require
		Description   null.String `db:"description"` // require
		ParticipantID null.Int    `db:"participant_id"`
	}
	User struct {
		ID        int         `db:"id"`
		Username  string      `db:"username"`
		Password  string      `db:"password"`
		OtpSecret string      `db:"otp_secret"`
		TgChat    null.String `db:"tg_chat"`
	}
	StatusListItem struct {
		ID          int    `db:"id"`
		Description string `db:"description"`
	}
	PriorityListItem struct {
		ID          int    `db:"id"`
		Description string `db:"description"`
	}
	CategoryListItem struct {
		ID          int    `db:"id"`
		Description string `db:"description"`
	}

	Comment struct {
		Message string `db:"message" json:"message"`
		UserID  int    `db:"user_id" json:"user_id"`
	}
)
