package models

import (
	"gopkg.in/guregu/null.v3"
)

type (
	Task struct {
		ID            int         `db:"id"` // require
		Category      null.String `db:"category"`
		Deadline      null.Time   `db:"deadline"`
		Status        string      `db:"status"` // require
		Comments      []byte      `db:"comments"`
		Priority      null.String `db:"priority"`
		CreatorID     int         `db:"creator_id"`  // require
		Description   string      `db:"description"` // require
		ParticipantID null.Int    `db:"participant_id"`
	}
	User struct {
		ID        int         `db:"id"`
		Username  string      `db:"username"`
		Password  string      `db:"password"`
		OtpSecret string      `db:"otp_secret"`
		TgChat    null.String `db:"tg_chat"`
	}
)
