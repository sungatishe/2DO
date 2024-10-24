package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}
