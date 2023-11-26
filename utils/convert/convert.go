package convert

import (
	"database/sql"
	"time"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"gopkg.in/guregu/null.v3"
)

func TaskToDBModel(request models.TaskModel) models.Task {
	return models.Task{
		Category: null.Int{
			NullInt64: sql.NullInt64{
				Int64: request.Category,
				Valid: request.Category != 0,
			},
		},
		Deadline: null.Time{
			Time:  time.Unix(request.Deadline, 0),
			Valid: request.Deadline != 0,
		},
		Status: null.Int{
			NullInt64: sql.NullInt64{
				Int64: request.Status,
				Valid: request.Status != 0,
			},
		},
		Priority: null.Int{
			NullInt64: sql.NullInt64{
				Int64: request.Priority,
				Valid: request.Priority != 0,
			},
		},
		CreatorID:   request.Creator,
		Description: request.Description,
		ParticipantID: null.Int{
			NullInt64: sql.NullInt64{
				Int64: request.ParticipantID,
				Valid: request.ParticipantID != 0,
			},
		},
	}
}
