package models

import (
	"time"
)

type Todo struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_Done"`
	Deadline    time.Time `json:"deadline"`
}
