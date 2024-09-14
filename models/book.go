package models

type Book struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Image     string  `json:"image"`
	Quantity  uint    `json:"quantity"`
	RentalFee float64 `json:"rental_fee" gorm:"default:'0'"`
	CreatedAt string  `gorm:"autoCreateTime"`
	UpdatedAt string  `gorm:"autoUpdateTime"`
}
