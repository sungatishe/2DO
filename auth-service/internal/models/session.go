package models

import "time"

type Session struct {
	ID        uint      `json:"ID"`
	UserID    uint      `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
