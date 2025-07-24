package request

import (
	"github.com/google/uuid"
	"time"
)

type CreateTask struct {
	Title string    `json:"title"`
	Note  string    `json:"note"`
	Due   time.Time `json:"due_date"`
}

type UpdateTask struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Note  string    `json:"note"`
	Due   time.Time `json:"due_date"`
}
