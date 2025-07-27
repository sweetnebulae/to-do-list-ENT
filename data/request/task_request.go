package request

import (
	"time"
)

type CreateTask struct {
	Title string    `json:"title"`
	Note  string    `json:"note"`
	Due   time.Time `json:"due_date"`
}

type UpdateTask struct {
	Title    *string    `json:"title"`
	Note     *string    `json:"note"`
	Complete *bool      `json:"complete"`
	Due      *time.Time `json:"due_date"`
}
