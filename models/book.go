package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Image     string    `json:"image"`
	Quantity  uint      `json:"quantity"`
	RentalFee float64   `json:"rental_fee" gorm:"default:0.0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
