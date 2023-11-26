package models

type (
	ShowTasksListRequest struct {
		UserID int
	}
	ShowTasksListResponse struct {
		Tasks []Task
	}
	MoveTaskRequest struct {
		Updates UpdateTask
	}
	MoveTaskResponse struct {
		Success bool
	}
	LinkTgResponse struct {
		Success   bool
		FailCause string
	}
	LinkTgRequest struct {
		TgChat string
		UserID int
	}
	UpdateTask struct {
		ID       int
		Deadline int64
	}
	UpdateTaskResponse struct {
		Success   bool   `json:"success"`
		FailCause string `json:"fail_cause"`
	}
)
