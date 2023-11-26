package models

type (
	AssignTaskRequest struct {
	}
	AssignTaskResponse struct {
	}
	TaskModel struct {
		ID            int    `json:"id"`
		Category      int64  `json:"category"`
		Status        int64  `json:"status"`
		Creator       int64  `json:"creator"`
		Description   string `json:"description"`
		Priority      int64  `json:"priority"`
		ParticipantID int64  `json:"participant_id"`
		Deadline      int64  `json:"deadline"`
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
	PriorityItem struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	}
	StatusItem struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	}
)
