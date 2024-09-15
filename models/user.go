package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Cell          string    `json:"cell"`
	PassCode      string    `json:"passcode"`
	Image         string    `json:"image"`
	Token         string    `json:"token"`
	Refresh_token string    `json:"refresh_token"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
