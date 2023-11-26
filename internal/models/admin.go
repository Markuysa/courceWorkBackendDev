package models

import "time"

type (
	AssignTaskRequest struct {
	}
	AssignTaskResponse struct {
	}
	CreateTaskRequest struct {
		Category      string
		Status        string
		Creator       int
		Description   string
		Priority      string
		ParticipantID int64
		Deadline      time.Time
	}
	CreateTaskResponse struct {
		Success   bool
		FailCause string
	}
	DeleteTaskRequest struct {
		TaskID string `json:"task_id"`
	}
	DeleteTaskResponse struct {
		FailCause string `json:"fail_cause"`
		Success   bool   `json:"success"`
	}

	TasksFilters struct {
		Limit  int
		Offset int
		UserID *int
	}
)
