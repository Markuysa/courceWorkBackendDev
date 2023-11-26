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
		Link      string `json:"link"`
		FailCause string `json:"fail_cause"`
	}
	LinkTgRequest struct {
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
	AddComment struct {
		TaskID  int     `json:"task_id"`
		Comment Comment `json:"comment"`
	}
)
