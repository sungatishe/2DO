package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_Done"`
	Deadline    time.Time `json:"deadline"`
}
