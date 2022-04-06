package models

type Rental struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	BookID     int    `json:"book_id"`
	UserID     int    `json:"user_id"`
	ReturnDate string `json:"should_return_at"`
	Book       Book   `gorm:"foreinKey:BookID"`
	User       User   `gorm:"foreignKey:UserID"`
	CreatedAt  string  `gorm:"autoCreateTime"`
	UpdatedAt  string  `gorm:"autoUpdateTime"`
}
