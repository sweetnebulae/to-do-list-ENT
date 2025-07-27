package response

import (
	"github.com/google/uuid"
	"time"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TaskResponse struct {
	ID         uuid.UUID  `json:"id"`
	Title      string     `json:"title"`
	Note       string     `json:"note"`
	IsComplete bool       `json:"isComplete"`
	DueDate    *time.Time `json:"dueDate,omitempty"`
}

type TaskListResponse struct {
	Success bool           `json:"success"`
	Data    []TaskResponse `json:"data"`
}

type TaskDetailResponse struct {
	Success bool         `json:"success"`
	Data    TaskResponse `json:"data"`
}

type TaskUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TaskDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
