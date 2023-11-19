package models

type (
	ShowTasksListRequest struct {
		UserID int
	}
	ShowTasksListResponse struct {
		Tasks []Task
	}
	MoveTaskRequest struct {
		TaskID  int
		Updates UpdateTask
	}
	MoveTaskResponse struct {
		Success bool
	}
	LinkTgResponse struct {
		Success bool
	}
	LinkTgRequest struct {
		TgChat string
		UserID int
	}
	UpdateTask struct {
		ID       int
		Deadline int64
	}
)
